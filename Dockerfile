FROM alpine:3.6 as builder

RUN apk add --update --no-cache ca-certificates

RUN mkdir -p /out && cp /etc/ssl/certs/ca-certificates.crt /out/

FROM scratch

COPY --from=builder /out/ca-certificates.crt /etc/ssl/certs/
ADD build/fastly-exporter /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/fastly-exporter"]