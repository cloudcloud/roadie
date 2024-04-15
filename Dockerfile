FROM node:slim AS fe
WORKDIR /app

COPY . .
RUN yarn && NODE_OPTIONS=--openssl-legacy-provider yarn build

FROM golang:alpine AS be

WORKDIR "/rd"
COPY . .
COPY --from=fe ["/app/dist/", "dist/"]
RUN apk add --no-cache git && \
    cp -r ./dist ./pkg/server/dist && \
    go build ./cmd/roadie && \
    mv roadie /

FROM golang:alpine AS release
ENTRYPOINT ["/roadie"]
COPY --from=be ["/roadie", "/roadie"]

