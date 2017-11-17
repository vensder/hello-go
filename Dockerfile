FROM golang:onbuild
ADD . /root/
WORKDIR /root
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -x webserver.go && chmod +x webserver

FROM alpine:latest
WORKDIR /root
COPY --from=0 /root/webserver .
EXPOSE 8080
CMD ["./webserver"]

