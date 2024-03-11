FROM alpine:3.14
RUN apk add --no-cache ca-certificates tzdata



ENTRYPOINT ["./app"]