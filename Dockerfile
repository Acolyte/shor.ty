FROM golang:1.16-alpine AS downloader
ARG VERSION

RUN apk add --no-cache git gcc musl-dev

WORKDIR /go/src/github.com/golang-migrate/migrate
RUN git clone https://github.com/golang-migrate/migrate.git .

ENV GO111MODULE=on
ENV DATABASES="postgres mysql redshift cassandra spanner cockroachdb clickhouse mongodb sqlserver firebird"
ENV SOURCES="file go_bindata github github_ee aws_s3 google_cloud_storage godoc_vfs gitlab"

RUN go build -a -o build/migrate.linux-386 -ldflags="-s -w -X main.Version=${VERSION}" -tags "$DATABASES $SOURCES" ./cmd/migrate

FROM golang:alpine as build

RUN apk add git --no-cache
WORKDIR /usr/local/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags "-s -w" -o shorty ./cmd/shorty/main.go

FROM alpine:3.8 as app
RUN apk add ca-certificates tzdata mysql-client --update && rm -rf /var/cache/apk/*

WORKDIR /usr/local/app
COPY --from=build /usr/local/app/shorty bin/shorty
COPY --from=downloader /go/src/github.com/golang-migrate/migrate/build/migrate.linux-386 bin/migrate
COPY ./migration /migration
COPY ./web /usr/local/app/web

RUN mkdir /usr/local/app/bin/swagger
COPY --from=build /usr/local/app/api/swagger.json swagger/swagger.json
COPY --from=build /usr/local/app/api/swagger.yaml swagger/swagger.yaml

CMD /usr/local/app/bin/shorty
