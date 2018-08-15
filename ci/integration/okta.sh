#echo "Fetching Okta bearer token..."
#
#export TEST_BEARER_TOKEN=$(curl -s --request POST --url "https://$OKTA_HOST/oauth2/default/v1/token" \
# --header 'accept: application/json' \
# --header "Authorization: Basic $OKTA_BASIC" \
# --header 'content-type: application/x-www-form-urlencoded' --data "$OKTA_BODY" | jq -r ".access_token")
#
#if [[ -z "$TEST_BEARER_TOKEN" ]]; then
#    echo "Unable to set BEARER_TOKEN"
#    exit 1
#else
#    echo "Successfully set BEARER_TOKEN"
#fi
#--------------------------------------------------------------------------------
echo "Fetching Okta bearer token..."

OKTA_AUTHENTICATION_RESPONSE=$(curl --request POST \
                                    --url "https://$OKTA_HOST/oauth2/default/v1/token" \
                                    --header 'accept: application/json' \
                                    --header "Authorization: Basic $OKTA_BASIC" \
                                    --header 'content-type: application/x-www-form-urlencoded' \
                                    --data "$OKTA_BODY")
TEST_BEARER_TOKEN=$(echo $OKTA_AUTHENTICATION_RESPONSE | jq -r ".access_token")
export TEST_BEARER_TOKEN

if [[ -z "$TEST_BEARER_TOKEN" ]]; then
    echo "Unable to set BEARER_TOKEN"
    exit 1
else
    echo "Successfully set BEARER_TOKEN"
fi
