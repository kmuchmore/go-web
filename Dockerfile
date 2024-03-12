#syntax=docker/dockerfile:experimental

FROM golang:latest
ENV PIP_CACHE_DIR=/var/cache/buildkit/pip
RUN mkdir -p $PIP_CACHE_DIR
RUN rm -f /etc/apt/apt.conf.d/docker-clean
RUN --mount=type=cache,target=/var/cache/apt \
    apt-get update; \
    apt-get install -yqq --no-install-recommends \
    wget xz-utils iputils-ping; \
    rm -rf /var/lib/apt/lists/*

ENV PATH /usr/local/lib/nodejs/bin:$PATH
RUN mkdir -p /usr/local/lib/nodejs; \
    mkdir -p /tmp/build; \
    wget -P /tmp/build https://nodejs.org/dist/v21.7.0/node-v21.7.0-linux-x64.tar.xz; \
    tar -xJf /tmp/build/node-v21.7.0-linux-x64.tar.xz --strip-components 1 -C /usr/local/lib/nodejs; \
    rm -rf /tmp/build; \
    npm install -g npm@latest; \
    npm install -g esbuild

RUN go install github.com/a-h/templ/cmd/templ@latest; \
    go install github.com/cosmtrek/air@latest; \
    go install github.com/bokwoon95/wgo@latest; \
    go install golang.org/x/tools/gopls@latest; \
    go clean -cache; \
    go clean -modcache
