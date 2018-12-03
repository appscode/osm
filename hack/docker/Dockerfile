FROM alpine

RUN set -x \
  && apk add --update --no-cache ca-certificates \
  && mkdir .osm \
  && touch .osm/config \
  && chmod 777 .osm/config

COPY osm /usr/bin/osm
COPY docker-entrypoint.sh /usr/local/bin/

RUN chmod +x /usr/local/bin/docker-entrypoint.sh

ENV PROVIDER ""

# S3 or S3 compatible storage
ENV AWS_ACCESS_KEY_ID ""
ENV AWS_SECRET_ACCESS_KEY ""
ENV AWS_ENDPOINT ""
ENV AWS_AUTH_TYPE ""
ENV AWS_REGION ""
ENV AWS_DISABLE_SSL ""
ENV CA_CERT_FILE ""

# Google Cloud Storage
ENV GOOGLE_PROJECT_ID ""
ENV GOOGLE_SERVICE_ACCOUNT_JSON_KEY_PATH ""
ENV GOOGLE_SCOPES ""

# Azure Blob storge
ENV AZURE_ACCOUNT_NAME ""
ENV AZURE_ACCOUNT_KEY ""

# Swift bucket
ENV SWIFT_AUTH_TOKEN ""
ENV SWIFT_DOMAIN ""
ENV SWIFT_KEY ""
ENV SWIFT_REGION ""
ENV SWIFT_STORAGE_URL ""
ENV SWIFT_TENANT_AUTH_URL ""
ENV SWIFT_TENANT_DOMAIN ""
ENV SWIFT_TENANT_ID ""
ENV SWIFT_TENANT_NAME ""
ENV SWIFT_TRUST_ID ""
ENV SWIFT_USERNAME ""

# local storage
ENV LOCAL_PATH ""

USER nobody:nobody
ENTRYPOINT ["docker-entrypoint.sh"]
