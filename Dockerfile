FROM alpine:3.3

COPY .dist/dockermon /usr/bin/dockermon

CMD ["/usr/bin/dockermon"]