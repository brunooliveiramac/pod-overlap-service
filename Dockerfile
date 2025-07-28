FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pod-overlap-service ./cmd/pod-overlap-service

# Final Stage
FROM scratch

COPY --from=builder /app/pod-overlap-service /pod-overlap-service

EXPOSE 8080

ENTRYPOINT ["/pod-overlap-service"]