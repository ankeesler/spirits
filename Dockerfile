FROM node:17-alpine as build-web

WORKDIR /workspace
ENV NODE_ENV production

COPY web/package.json web/package-lock.json web/
RUN npm --prefix web install
COPY web/public web/public
COPY web/src web/src
RUN npm --prefix web run build

FROM golang:1.17-alpine as build-api

WORKDIR /workspace

COPY api/go.mod api/go.sum api/
RUN cd api && go mod download
COPY api/main.go api/main.go
COPY api/internal api/internal
RUN cd api && CGO_ENABLED=0 go build -v -ldflags '-extldflags "-static"' -o spirits .

FROM gcr.io/distroless/static:latest

COPY --from=build-web /workspace/web/build /web
COPY --from=build-api /workspace/api/spirits /spirits

EXPOSE 12345

ENTRYPOINT ["/spirits", "-web-assets-dir", "/web"]
