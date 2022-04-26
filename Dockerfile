FROM golang:1.18-alpine as build

WORKDIR /workspace

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY cmd cmd
COPY pkg pkg
COPY internal internal
RUN CGO_ENABLED=0 go build -v -ldflags '-extldflags "-static"' -o spirits ./cmd/spirits

FROM gcr.io/distroless/static:latest

COPY --from=build /workspace/spirits /spirits

EXPOSE 12345

ENTRYPOINT ["/spirits"]
