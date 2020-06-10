FROM golang:latest AS builder
WORKDIR /go/src/mail

COPY . .
ENV GOPROXY https://goproxy.io
RUN export GO111MODULE=on && go mod download

#RUN go test ./... -coverprofile .testCoverage.txt \
#    && go tool cover -func=.testCoverage.txt
RUN CGO_ENABLED=0 go build -o app_d ./cmd/main.go
#     && CGO_ENABLED=0 go build ./cmd/migrate

FROM alpine:3.10
RUN apk update && apk --no-cache add ca-certificates

LABEL \
    SERVICE_80_NAME=mail_http \
    SERVICE_NAME=mail \
    description="mail" \
    maintainer="pangxieke"

EXPOSE 3000
COPY --from=builder /go/src/mail/app_d /bin/app
#COPY --from=builder /go/src/mail/migrate /bin/migrate
ENTRYPOINT ["app"]
