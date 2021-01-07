FROM alpine:3.9 as builder
RUN apk update && apk add ca-certificates tzdata

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo/
COPY go-keeper-linux64 go-keeper-linux64
EXPOSE 8080
CMD ["./go-keeper-linux64"]