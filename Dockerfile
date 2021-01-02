FROM node:slim AS fe
WORKDIR "/app"

COPY yarn.lock yarn.lock
COPY package.json package.json
COPY babel.config.js babel.config.js
RUN yarn --frozen-lockfile --silent --non-interactive --link-duplicates

COPY public public
COPY src src
RUN yarn build

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

