#!/bin/bash -e

# login to system
./bin/scloud config set --key env --value $TEST_ENVIRONMENT_1
./bin/scloud config set --key username --value $TEST_USERNAME
./bin/scloud config set --key tenant --value system
./bin/scloud config set --key tenant-scoped --value "true"
./bin/scloud config set --key region --value $REGION
./bin/scloud login --use-pkce --pwd $TEST_PASSWORD
# also login to test tenant
./bin/scloud config set --key tenant --value $TEST_TENANT_SCOPED
./bin/scloud login --use-pkce --pwd $TEST_PASSWORD

# Cross-platform sed -i: https://stackoverflow.com/a/38595160
sedi () {
    sed --version >/dev/null 2>&1 && sed -i -- "$@" || sed -i "" "$@"
}

touch .env

if grep -q '^SCLOUD_CACHE_PATH=' .env; then
  if sedi "s/SCLOUD_CACHE_PATH=.*/SCLOUD_CACHE_PATH=.scloud_context/" .env; then
     echo "SCLOUD_CACHE_PATH updated in .env"
  fi
else
  if echo "SCLOUD_CACHE_PATH=.scloud_context" | tee -a .env >/dev/null; then
     echo "SCLOUD_CACHE_PATH written to .env"
  fi
fi
