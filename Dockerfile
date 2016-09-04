FROM alpine:3.3

COPY .dist/dockermon /usr/bin/dockermon

ENTRYPOINT ["/usr/bin/dockermon"]