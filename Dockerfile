FROM alpine:3.9 as builder
#RUN apk update && apk add ca-certificates tzdata
RUN apk update

FROM scratch
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo/
COPY gprc_example-linux64 gprc_example-linux64
EXPOSE 5300:5300
#CMD ["./gprc_example-linux64"]
ENTRYPOINT [ "./gprc_example-linux64" ]
