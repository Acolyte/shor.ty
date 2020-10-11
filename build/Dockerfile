FROM golang:alpine as build

RUN apk add git --no-cache
WORKDIR /usr/local/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init -g ./internal/app/shorty/api/routers.go -o ./swagger

RUN go build -ldflags "-s -w" -o shorty cmd/shorty/main.go

FROM alpine:3.8
RUN apk add ca-certificates tzdata --update && rm -rf /var/cache/apk/*

WORKDIR /usr/local/app
COPY --from=build /usr/local/app/shorty bin/shorty

RUN mkdir /usr/local/app/bin/swagger
COPY --from=build /usr/local/app/swagger/swagger.json swagger/swagger.json
COPY --from=build /usr/local/app/swagger/swagger.yaml swagger/swagger.yaml

CMD /usr/local/app/bin/shorty
