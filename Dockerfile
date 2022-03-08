FROM node:17-alpine as build-web

WORKDIR /workspace
ENV NODE_ENV production

COPY web web
RUN npm --prefix web install
RUN npm --prefix web run build

FROM golang:1.17-alpine as build-api

WORKDIR /workspace

COPY api api
RUN cd api && CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o spirits .

FROM gcr.io/distroless/static:latest

COPY --from=build-web /workspace/web/build /web
COPY --from=build-api /workspace/api/spirits /spirits

EXPOSE 12345

ENTRYPOINT ["/spirits", "-web-assets-dir", "/web"]
