# Sử dụng image Go để build
FROM golang:1.22-alpine AS builder

# Cài đặt các công cụ cần thiết
RUN apk add --no-cache git

# Đặt thư mục làm việc
WORKDIR /app

# Copy go.mod và go.sum trước (tối ưu cache khi cài deps)
COPY go.mod go.sum ./

# Tải dependencies
RUN go mod download

# Copy toàn bộ source code
COPY . .

# Build app (tạo file nhị phân tên app)
RUN go build -o app .

# Image chạy thật sự
FROM alpine:latest

# Thêm certificate để gọi HTTPS API nếu cần
RUN apk add --no-cache ca-certificates

# Tạo thư mục app
WORKDIR /root/

# Copy binary từ builder sang
COPY --from=builder /app/app .

# Set biến PORT (Railway cần để detect)
ENV PORT=8080

# Expose port cho Railway
EXPOSE 8080

# Lệnh chạy app
CMD ["./app"]
