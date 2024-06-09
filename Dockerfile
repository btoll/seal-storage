FROM golang:1.21.6-bookworm AS build
WORKDIR /
COPY go.mod main.go /
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o /seal-storage

FROM gcr.io/distroless/base-debian11 AS final
COPY --from=build /seal-storage .
ENTRYPOINT ["/seal-storage"]

