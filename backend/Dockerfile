# --- Stage 1: Build binary ---
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Toi uu cache
COPY go.mod go.sum ./
RUN go mod download

# Copy code con lai
COPY . .

#Build binary statically, bo symbol de nhe
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server .

FROM gcr.io/distroless/static-debian12
WORKDIR /app

COPY --from=builder /app/server /app/server

ENV GIN_MODE=release
ENV PORT=8080
EXPOSE 8080

CMD ["/app/server"]