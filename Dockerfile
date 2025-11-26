# -------- Build stage --------
FROM golang:1.23 AS builder

# Set working directory
WORKDIR /app

# Copy all source code
COPY . .

# Build a fully static binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o market_client main.go

# -------- Final stage --------
FROM scratch

# Copy the binary from the builder
COPY --from=builder /app/market_client /market_client

# Run the binary
CMD ["/market_client"]
