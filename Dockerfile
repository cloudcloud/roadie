FROM node:slim AS fe
WORKDIR /app

COPY . .
RUN yarn && NODE_OPTIONS=--openssl-legacy-provider yarn build

FROM golang:alpine AS be

WORKDIR "/rd"
COPY . .
COPY --from=fe ["/app/dist/", "dist/"]
RUN apk add --no-cache git && \
      GO111MODULE=off go get -u github.com/kevinburke/go-bindata/... && \
      go-bindata -o ./pkg/server/assets.go -pkg server -prefix dist/ dist/... && \
      go build ./cmd/roadie && \
      mv roadie /

FROM golang:alpine AS release
ENTRYPOINT ["/roadie"]
COPY --from=be ["/roadie", "/roadie"]

