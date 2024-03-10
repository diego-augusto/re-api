#Build stage
FROM golang:1.22-alpine AS builder
WORKDIR /
COPY . . 
# RUN go mod download
# COPY cmd ./cmd
# COPY pkg ./pkg
# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o api /cmd/api/main.go
RUN go build -o api /cmd/api/main.go

#Test stage
FROM build AS test
RUN go test -v ./...

#Release stage
FROM alpine AS release
WORKDIR /
COPY --from=builder  /api /api
# USER nonroot:nonroot
ENTRYPOINT [ "/api"]