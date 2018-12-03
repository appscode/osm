#!/bin/bash
set -eo pipefail

# if command does not start with "osm", prepend it.
if [ "${1:0:3}" != 'osm' ]; then
  set -- osm "$@"
fi

# if configured via environment variables then create a configuration
if [ "$PROVIDER" != "" ]; then
  CMD="osm config set-context ${PROVIDER}"
  ARGS=("--provider=$PROVIDER")
  case "$PROVIDER" in
    "s3")
      if [ "$AWS_ACCESS_KEY_ID" != "" ]; then
        ARGS=(${ARGS[@]} "--s3.access_key_id=$AWS_ACCESS_KEY_ID")
      fi
      if [ "$AWS_SECRET_ACCESS_KEY" != "" ]; then
        ARGS=(${ARGS[@]} "--s3.secret_key=$AWS_SECRET_ACCESS_KEY")
      fi
      if [ "$AWS_ENDPOINT" != "" ]; then
        ARGS=(${ARGS[@]} "--s3.endpoint=$AWS_ENDPOINT")
      fi
      if [ "$AWS_REGION" != "" ]; then
        ARGS=(${ARGS[@]} "--s3.region=$AWS_REGION")
      fi
      if [ "$AWS_AUTH_TYPE" != "" ]; then
        ARGS=(${ARGS[@]} "--s3.auth_type=$AWS_AUTH_TYPE")
      fi
      if [ "$CA_CERT_FILE" != "" ]; then
        ARGS=(${ARGS[@]} "--s3.cacert_file=$CA_CERT_FILE")
      fi
      if [ "$AWS_DISABLE_SSL" != "" ]; then
        ARGS=(${ARGS[@]} "--s3.disable_ssl=$AWS_DISABLE_SSL")
      fi
      ;;
    "azure")
      if [ "$AZURE_ACCOUNT_NAME" != "" ]; then
        ARGS=(${ARGS[@]} "--azure.account=$AZURE_ACCOUNT_NAME")
      fi
      if [ "$AZURE_ACCOUNT_KEY" != "" ]; then
        ARGS=(${ARGS[@]} "--azure.key=$AZURE_ACCOUNT_KEY")
      fi
      ;;
    "google")
      if [ "$GOOGLE_PROJECT_ID" != "" ]; then
        ARGS=(${ARGS[@]} "--google.project_id=$GOOGLE_PROJECT_ID")
      fi
      if [ "$GOOGLE_SERVICE_ACCOUNT_JSON_KEY_PATH" != "" ]; then
        ARGS=(${ARGS[@]} "--google.json_key_path=$GOOGLE_SERVICE_ACCOUNT_JSON_KEY_PATH")
      fi
      if [ "$GOOGLE_SCOPES" != "" ]; then
        ARGS=(${ARGS[@]} "--google.scopes=$GOOGLE_SCOPES")
      fi
      ;;
    "swift")
      if [ "$SWIFT_AUTH_TOKEN" != "" ]; then
        ARGS=(${ARGS[@]} "--swift.auth_token=$SWIFT_AUTH_TOKEN")
      fi
      if [ "$SWIFT_DOMAIN" != "" ]; then
        ARGS=(${ARGS[@]} "--swift.domain=$SWIFT_DOMAIN")
      fi
      if [ "$SWIFT_KEY" != "" ]; then
        ARGS=(${ARGS[@]} "--swift.key=$SWIFT_KEY")
      fi
      if [ "$SWIFT_REGION" != "" ]; then
        ARGS=(${ARGS[@]} "--swift.region=$SWIFT_REGION")
      fi
      if [ "$SWIFT_STORAGE_URL" != "" ]; then
        ARGS=(${ARGS[@]} "--swift.storage_url=$SWIFT_STORAGE_URL")
      fi
      if [ "$SWIFT_TENANT_AUTH_URL" != "" ]; then
        ARGS=(${ARGS[@]} "--swift.tenant_auth_url=$SWIFT_TENANT_AUTH_URL")
      fi
      if [ "$SWIFT_TENANT_DOMAIN" != "" ]; then
        ARGS=(${ARGS[@]} "--swift.tenant_domain=$SWIFT_TENANT_DOMAIN")
      fi
      if [ "$SWIFT_TENANT_ID" != "" ]; then
        ARGS=(${ARGS[@]} "--swift.tenant_id=$SWIFT_TENANT_ID")
      fi
      if [ "$SWIFT_TENANT_NAME" != "" ]; then
        ARGS=(${ARGS[@]} "--swift.tenant_name=$SWIFT_TENANT_NAME")
      fi
      if [ "$SWIFT_TRUST_ID" != "" ]; then
        ARGS=(${ARGS[@]} "--swift.trust_id=$SWIFT_TRUST_ID")
      fi
      if [ "$SWIFT_USERNAME" != "" ]; then
        ARGS=(${ARGS[@]} "--swift.username=$SWIFT_USERNAME")
      fi
      ;;
    "local")
      if [ "$LOCAL_PATH" != "" ]; then
        ARGS=(${ARGS[@]} "--local.path=$LOCAL_PATH")
      fi
      ;;
  esac
  echo "Configuring osm context for $PROVIDER storage"
  echo "${CMD} ${ARGS[@]}"
  ${CMD} "${ARGS[@]}" >/dev/null
  echo "Successfully configured"
fi

# run final command
echo
echo "Running main command....."
echo "$@"
exec "$@"
