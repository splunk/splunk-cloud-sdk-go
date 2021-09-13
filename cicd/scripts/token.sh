# Cross-platform sed -i: https://stackoverflow.com/a/38595160
sedi () {
    sed --version >/dev/null 2>&1 && sed -i -- "$@" || sed -i "" "$@"
}

# If any of the required env variables are unset, fail fast

if [[ -z "$BASE64_ENCODED_BASIC_AUTH_TENANT_SCOPED" ]]; then
    echo "BASE64_ENCODED_BASIC_AUTH_TENANT_SCOPED was not set"
    exit 1
fi

if [[ -z "$IDP_HOST_TENANT_SCOPED" ]]; then
    echo "IDP_HOST_TENANT_SCOPED was not set"
    exit 1
fi

# Optional env variables
IDP_TOKEN_BODY="grant_type=client_credentials"

echo "Fetching access token..."
CURL_RESPONSE=$(curl --request POST --url $IDP_HOST_TENANT_SCOPED/token \
 --header 'accept: application/json' \
 --header "Authorization: Basic $BASE64_ENCODED_BASIC_AUTH_TENANT_SCOPED" \
 --header 'content-type: application/x-www-form-urlencoded' --data "$IDP_TOKEN_BODY")

ACCESS_TOKEN=$(printf "$CURL_RESPONSE" | jq -r ".access_token")

if [ -z "$ACCESS_TOKEN" ] || [ "$ACCESS_TOKEN" = "null" ]; then
    echo "Unable to set ACCESS_TOKEN, response from idp:\n\n$CURL_RESPONSE"
    exit 1
fi

touch .env

if grep -q '^BEARER_TOKEN=' .env; then
  if sedi "s/BEARER_TOKEN=.*/BEARER_TOKEN=${ACCESS_TOKEN}/" .env; then
    echo "access_token updated in .env"
  fi
else
  if echo "BEARER_TOKEN=${ACCESS_TOKEN}" | tee -a .env >/dev/null; then
    echo "access_token written to .env"
  fi
fi

grep -q '^BEARER_TOKEN=' .env

if [ $? -ne 0 ]
then
  echo "failed writing access_token to .env"
  exit 1
fi
