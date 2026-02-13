FROM golang:1.22-alpine AS build
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o oris-api ./cmd/api

FROM alpine:3.20
WORKDIR /app
COPY --from=build /app/oris-api /app/oris-api
EXPOSE 8080
ENTRYPOINT ["/app/oris-api"]
