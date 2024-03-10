#Build stage
FROM golang:1.22-alpine AS build-stage
WORKDIR /
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY pkg ./pkg
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o api /cmd/api/main.go

#Test stage
FROM build-stage AS tests-stage
RUN go test -v ./...

#Release stage
FROM gcr.io/distroless/base-debian11 AS release-stage
WORKDIR /
COPY --from=build-stage  /api /api
USER nonroot:nonroot
ENTRYPOINT [ "/api"]