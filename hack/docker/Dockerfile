FROM alpine

RUN set -x \
  && apk add --update --no-cache ca-certificates bash \
  && mkdir /.osm \
  && touch /.osm/config \
  && chmod -R 777 .osm

COPY osm /usr/bin/osm
COPY docker-entrypoint.sh /usr/local/bin/

USER nobody:nobody
ENTRYPOINT ["docker-entrypoint.sh"]
