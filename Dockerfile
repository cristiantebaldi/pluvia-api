FROM alpine:latest AS timezone_build
RUN apk --no-cache add tzdata ca-certificates  


FROM golang:1.25.0-alpine3.22 AS builder

RUN apk --no-cache add tzdata ca-certificates

ADD . /go/api

WORKDIR /go/api

RUN go install github.com/swaggo/swag/cmd/swag@v1.8.7

RUN /go/bin/swag init -d adapter/http --parseDependency --parseInternal --parseDepth 3 -o adapter/http/docs

RUN rm -rf deploy
RUN mkdir deploy
RUN go mod tidy

RUN CGO_ENABLED=0 go build -o go_app adapter/http/main.go 
RUN mv go_app ./deploy/go_app
RUN mv config.json ./deploy/config.json
RUN mv adapter/http/docs/ ./deploy/docs
RUN mv database ./deploy/database


FROM scratch AS production

COPY --from=timezone_build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=timezone_build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=builder /go/api/deploy /api/

WORKDIR /api

ENTRYPOINT  ["/api/go_app"]