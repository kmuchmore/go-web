#syntax=docker/dockerfile:experimental

FROM golang:latest as go
FROM node:latest

RUN --mount=type=cache,target=/var/cache/apt \
    apt-get update; \
    apt-get install -yqq --no-install-recommends \
    sudo wget xz-utils iputils-ping sudo && \
    rm -rf /var/lib/apt/lists/*

COPY --from=go /usr/local/go /usr/local/go
ENV PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
ENV GOPATH=/go
RUN mkdir /go && \
    chmod 777 /go

RUN echo "node ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers

RUN npm install -g npm@latest && \
    npm install -g esbuild

USER node

RUN go install github.com/a-h/templ/cmd/templ@latest && \
    go install github.com/air-verse/air@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install github.com/bokwoon95/wgo@latest && \
    go clean -cache && \
    go clean -modcache

