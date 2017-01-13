FROM alpine:3.3

COPY .dist/dockemon /usr/bin/dockemon

CMD ["/usr/bin/dockemon"]