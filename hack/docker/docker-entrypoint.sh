#!/bin/bash
set -eo pipefail

PROVIDER=${PROVIDER:-}

# S3 or S3 compatible storage
AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID:-}
AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY:-}
AWS_ENDPOINT=${AWS_ENDPOINT:-}
AWS_AUTH_TYPE=${AWS_AUTH_TYPE:-}
AWS_REGION=${AWS_REGION:-}
AWS_DISABLE_SSL=${AWS_DISABLE_SSL:-}
CA_CERT_FILE=${CA_CERT_FILE:-}

# Google Cloud Storage
GOOGLE_PROJECT_ID=${GOOGLE_PROJECT_ID:-}
GOOGLE_SERVICE_ACCOUNT_JSON_KEY_PATH=${GOOGLE_SERVICE_ACCOUNT_JSON_KEY_PATH:-}
GOOGLE_SCOPES=${GOOGLE_SCOPES:-}

# Azure Blob storge
AZURE_ACCOUNT_NAME=${AZURE_ACCOUNT_NAME:-}
AZURE_ACCOUNT_KEY=${AZURE_ACCOUNT_KEY:-}

# Swift bucket
SWIFT_AUTH_TOKEN=${SWIFT_AUTH_TOKEN:-}
SWIFT_DOMAIN=${SWIFT_DOMAIN:-}
SWIFT_KEY=${SWIFT_KEY:-}

SWIFT_REGION=${SWIFT_REGION:-}
SWIFT_STORAGE_URL=${SWIFT_STORAGE_URL:-}
SWIFT_TENANT_AUTH_URL=${SWIFT_TENANT_AUTH_URL:-}
SWIFT_TENANT_DOMAIN=${SWIFT_TENANT_DOMAIN:-}
SWIFT_TENANT_ID=${SWIFT_TENANT_ID:-}
SWIFT_TENANT_NAME=${SWIFT_TENANT_NAME:-}
SWIFT_TRUST_ID=${SWIFT_TRUST_ID:-}
SWIFT_USERNAME=${SWIFT_USERNAME:-}

# local storage
LOCAL_PATH=${LOCAL_PATH:-}

# if command does not start with "osm", prepend it.
if [ "${1:0:3}" != 'osm' ]; then
  set -- osm "$@"
fi

# if configured via environment variables then create a configuration
if [ "$PROVIDER" != "" ]; then
  cmd="osm config set-context ${PROVIDER}"
  args=("--provider=$PROVIDER")
  case "$PROVIDER" in
    "s3")
      if [ "$AWS_ACCESS_KEY_ID" != "" ]; then
        args=(${args[@]} "--s3.access_key_id=$AWS_ACCESS_KEY_ID")
      fi
      if [ "$AWS_SECRET_ACCESS_KEY" != "" ]; then
        args=(${args[@]} "--s3.secret_key=$AWS_SECRET_ACCESS_KEY")
      fi
      if [ "$AWS_ENDPOINT" != "" ]; then
        args=(${args[@]} "--s3.endpoint=$AWS_ENDPOINT")
      fi
      if [ "$AWS_REGION" != "" ]; then
        args=(${args[@]} "--s3.region=$AWS_REGION")
      fi
      if [ "$AWS_AUTH_TYPE" != "" ]; then
        args=(${args[@]} "--s3.auth_type=$AWS_AUTH_TYPE")
      fi
      if [ "$CA_CERT_FILE" != "" ]; then
        args=(${args[@]} "--s3.cacert_file=$CA_CERT_FILE")
      fi
      if [ "$AWS_DISABLE_SSL" != "" ]; then
        args=(${args[@]} "--s3.disable_ssl=$AWS_DISABLE_SSL")
      fi
      ;;
    "azure")
      if [ "$AZURE_ACCOUNT_NAME" != "" ]; then
        args=(${args[@]} "--azure.account=$AZURE_ACCOUNT_NAME")
      fi
      if [ "$AZURE_ACCOUNT_KEY" != "" ]; then
        args=(${args[@]} "--azure.key=$AZURE_ACCOUNT_KEY")
      fi
      ;;
    "google")
      if [ "$GOOGLE_PROJECT_ID" != "" ]; then
        args=(${args[@]} "--google.project_id=$GOOGLE_PROJECT_ID")
      fi
      if [ "$GOOGLE_SERVICE_ACCOUNT_JSON_KEY_PATH" != "" ]; then
        args=(${args[@]} "--google.json_key_path=$GOOGLE_SERVICE_ACCOUNT_JSON_KEY_PATH")
      fi
      if [ "$GOOGLE_SCOPES" != "" ]; then
        args=(${args[@]} "--google.scopes=$GOOGLE_SCOPES")
      fi
      ;;
    "swift")
      if [ "$SWIFT_AUTH_TOKEN" != "" ]; then
        args=(${args[@]} "--swift.auth_token=$SWIFT_AUTH_TOKEN")
      fi
      if [ "$SWIFT_DOMAIN" != "" ]; then
        args=(${args[@]} "--swift.domain=$SWIFT_DOMAIN")
      fi
      if [ "$SWIFT_KEY" != "" ]; then
        args=(${args[@]} "--swift.key=$SWIFT_KEY")
      fi
      if [ "$SWIFT_REGION" != "" ]; then
        args=(${args[@]} "--swift.region=$SWIFT_REGION")
      fi
      if [ "$SWIFT_STORAGE_URL" != "" ]; then
        args=(${args[@]} "--swift.storage_url=$SWIFT_STORAGE_URL")
      fi
      if [ "$SWIFT_TENANT_AUTH_URL" != "" ]; then
        args=(${args[@]} "--swift.tenant_auth_url=$SWIFT_TENANT_AUTH_URL")
      fi
      if [ "$SWIFT_TENANT_DOMAIN" != "" ]; then
        args=(${args[@]} "--swift.tenant_domain=$SWIFT_TENANT_DOMAIN")
      fi
      if [ "$SWIFT_TENANT_ID" != "" ]; then
        args=(${args[@]} "--swift.tenant_id=$SWIFT_TENANT_ID")
      fi
      if [ "$SWIFT_TENANT_NAME" != "" ]; then
        args=(${args[@]} "--swift.tenant_name=$SWIFT_TENANT_NAME")
      fi
      if [ "$SWIFT_TRUST_ID" != "" ]; then
        args=(${args[@]} "--swift.trust_id=$SWIFT_TRUST_ID")
      fi
      if [ "$SWIFT_USERNAME" != "" ]; then
        args=(${args[@]} "--swift.username=$SWIFT_USERNAME")
      fi
      ;;
    "local")
      if [ "$LOCAL_PATH" != "" ]; then
        args=(${args[@]} "--local.path=$LOCAL_PATH")
      fi
      ;;
  esac
  echo "Configuring osm context for $PROVIDER storage"
  cmd=$(echo "$cmd ${args[@]}")
  echo "$cmd"
  $cmd
  echo "Successfully configured"
fi

# run final command
echo
echo "Running main command....."
echo "$@"
exec $@
