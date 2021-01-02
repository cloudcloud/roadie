FROM node:slim AS fe

COPY yarn.lock yarn.lock
COPY package.json package.json
COPY babel.config.js babel.config.js
RUN yarn

COPY src src
RUN yarn build

FROM golang:alpine AS be

WORKDIR "/rd"
COPY . .
COPY --from=fe ["dist/", "dist/"]
RUN apk add --no-cache git && \
      GO111MODULE=off go get -u github.com/kevinburke/go-bindata/... && \
      go-bindata -o ./pkg/server/assets.go -prefix dist/ dist/... && \
      sed -i "s/package main/package server/g" ./pkg/server/assets.go && \
      go build ./cmd/roadie && \
      mv roadie /

FROM golang:alpine AS release
ENTRYPOINT ["/roadie"]
COPY --from=be ["/roadie", "/roadie"]

