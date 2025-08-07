# Use a specific version of Go for reproducibility
FROM golang:1.24-alpine AS builder

WORKDIR /app

# 1. Copy go.mod and go.sum first to leverage Docker cache
COPY ../../go.mod ../../go.sum ./
RUN go mod download

# 2. Copy the rest of the source code
COPY ../../ ./

# 3. Set the final working directory and build
WORKDIR /app/cmd/menu-service
# CGO_ENABLED=0 is important for distroless, as it creates a static binary
# with no dependencies on C libraries that might be missing in distroless.
RUN CGO_ENABLED=0 go build -o menu-service .

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/cmd/menu-service/menu-service .

EXPOSE 8080

CMD ["./menu-service"]