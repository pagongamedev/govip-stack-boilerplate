# 2 State Build
FROM golang:latest AS builder

# Load Argument
ARG CI_PLATFORM
ARG CI_LOGIN
ARG CI_TOKEN

WORKDIR /workspace
# Gitlab Access Token : Permission read_registry for Pull Private Repo Go Package for Build
RUN echo "machine ${CI_PLATFORM} login ${CI_LOGIN} password ${CI_TOKEN}" > ~/.netrc
ENV GO111MODULE on
ENV GOARCH amd64
ENV CGO_ENABLED 0
ENV GOOS linux

COPY . .
RUN go build -o /workspace/app -a -ldflags '-w -s' /workspace/main.go
#  -mod vendor

FROM alpine

# Load Argument
ARG API_ENV
ARG API_VERSION
ARG API_HOST
ARG PROJECT_NAME

ARG SENTRY_DSN

ARG DB_NAME
ARG DB_HOST
ARG DB_USER
ARG DB_PASS
ARG DB_PORT
ARG DB_BASE64_SSLKEY
ARG DB_BASE64_SSLCERT

# Set env By CI/CD
ENV API_ENV=${API_ENV}
ENV API_VERSION=${API_VERSION}

ENV API_HOST=${API_HOST}
ENV PROJECT_NAME=${PROJECT_NAME}

ENV SENTRY_DSN=${SENTRY_DSN}

# Database
ENV DB_NAME=${DB_NAME}
ENV DB_HOST=${DB_HOST}
ENV DB_USER=${DB_USER}
ENV DB_PASS=${DB_PASS}
ENV DB_PORT=${DB_PORT}

ENV DB_SSLMODE=require

RUN mkdir ~/.ssh/
RUN echo "${DB_BASE64_SSLKEY}" | tr -d ' ' | base64 -d >  ~/.ssh/client.key
RUN chmod 600 ~/.ssh/client.key
ENV DB_SSLKEY=root/.ssh/client.key

RUN echo "${DB_BASE64_SSLCERT}" | tr -d ' ' | base64 -d >  ~/.ssh/client.cert
RUN chmod 600 ~/.ssh/client.cert
ENV DB_SSLCERT=root/.ssh/client.cert

# RUN ssh-keyscan bitbucket.org >> ~/.ssh/known_hosts

# Set env
ENV SWAGGER_SERVICE_NAME=govip-boilerplate
ENV SWAGGER_SERVICE_TITLE="GoVIP Boilerplate"

ENV SWAGGER_BASEPATH=

ENV SWAGGER_VERSION=0.0.1
ENV SWAGGER_DESCRIPTION="GoVIP Boilerplate"

COPY /domain/auth/i18n /domain/auth/i18n

COPY --from=builder /workspace/app /app
RUN apk add --no-cache --upgrade curl
EXPOSE 8081 8082 8083
ENTRYPOINT ["/app"]
