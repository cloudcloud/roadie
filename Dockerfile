FROM node:slim AS fe
WORKDIR /app

COPY . .
RUN yarn && NODE_OPTIONS=--openssl-legacy-provider yarn build

FROM golang:alpine AS be

WORKDIR "/rd"
COPY . .
COPY --from=fe ["/app/pkg/server/dist/", "pkg/server/dist/"]
RUN apk add --no-cache git && \
    go build ./cmd/roadie && \
    mv roadie /

FROM golang:alpine AS release
ENTRYPOINT ["/roadie"]
COPY --from=be ["/roadie", "/roadie"]

