ARG GO_VERSION=1.21.3-alpine3.18
ARG DELVE_VERSION=1.21.1
ARG OAPICODEGEN_VERSION=1.15.0


## Base image
FROM golang:${GO_VERSION} AS base

WORKDIR /go/src/app

ENV CGO_ENABLED=0


## Development image
FROM base AS development

ARG DELVE_VERSION
ARG OAPICODEGEN_VERSION

RUN apk add \
        bash \
        curl \
        git \
        jq \
        python3 \
        zsh \
 && go install github.com/go-delve/delve/cmd/dlv@v${DELVE_VERSION} \
 && go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v${OAPICODEGEN_VERSION}

ARG USER_ID=1000
ENV USER_NAME=default

ENV PROMPT="%B%F{cyan}%n%f@%m:%F{yellow}%~%f %F{%(?.green.red[%?] )}>%f %b"

RUN adduser -D -u $USER_ID ${USER_NAME} \
 && chown -R ${USER_NAME}: /go/pkg

USER ${USER_NAME}
