ARG GO_VERSION=1.22.2-bookworm
ARG GOIMPORTS_VERSION=0.20.0
ARG DELVE_VERSION=1.22.1


## Base image
FROM golang:${GO_VERSION} AS base

WORKDIR /go/src/app

ENV CGO_ENABLED=0


## Development image
FROM base AS development

ARG GOIMPORTS_VERSION
ARG DELVE_VERSION

RUN apt update \
 && apt install -y \
        curl \
        git \
        jq \
        python3 \
        python3-pip \
        zsh \
 && apt clean \
 && pip install --break-system-packages \
        pyyaml \
 && go install golang.org/x/tools/cmd/goimports@v${GOIMPORTS_VERSION} \
 && go install github.com/go-delve/delve/cmd/dlv@v${DELVE_VERSION}

ARG USER_ID=1000
ENV USER_NAME=default

RUN useradd --user-group --create-home --uid ${USER_ID} ${USER_NAME} \
 && chown -R ${USER_NAME}: /go

USER ${USER_NAME}

ENV PROMPT="%B%F{cyan}%n%f@%m:%F{yellow}%~%f %F{%(?.green.red[%?] )}>%f %b"

RUN touch /home/${USER_NAME}/.zshrc
