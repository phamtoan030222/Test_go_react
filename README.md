# 📝 Todo App (Go + React + Docker)

## 🌍 Demo
👉 [Truy cập ứng dụng tại đây](https://test-go-react.vercel.app/)

---

## 📌 Giới thiệu
Dự án này bao gồm:
- **Backend**: Viết bằng **Golang** (build binary nhẹ, tối ưu).
- **Frontend**: Xây dựng bằng **React + Vite + TypeScript + Chakra UI**.
- **Triển khai**: Sử dụng **Docker** và **docker-compose** để chạy đồng thời backend (API) và frontend (web UI).

---

## ⚙️ Cấu trúc thư mục

├── backend/              # Source code backend (Go)
│   ├── handlers/         # Xử lý request
│   ├── models/           # Định nghĩa models
│   ├── storage/          # Xử lý DB/storage
│   ├── main.go           # Entry point
│   └── Dockerfile        # Dockerfile backend
│
├── client/               # Source code frontend (React + Vite)
│   ├── src/components/   # Các component UI
│   ├── public/           # Static files
│   ├── nginx.conf        # Config Nginx
│   └── Dockerfile        # Dockerfile frontend
│
├── docker-compose.yml    # Orchestrate frontend + backend
└── README.md             # Tài liệu hướng dẫn
🚀 Cách chạy project
1. Yêu cầu hệ thống
Docker >= 20.x

Docker Compose >= v2.x

2. Chạy bằng Docker Compose
Tại thư mục gốc project, chạy lệnh:

docker-compose up --build
Sau khi build xong:

API (backend): http://localhost:8080

Web (frontend): http://localhost:8081

3. Cấu hình môi trường
Bạn có thể cấu hình thông qua biến môi trường trong docker-compose.yml:

Backend (api)

environment:
  - PORT=8080
  - CORS_ORIGINS=http://localhost:8081
Frontend (web)

args:
  - VITE_API_URL=http://localhost:8080/api
Khi gọi API từ web, frontend sẽ sử dụng biến VITE_API_URL.

4. Build và chạy thủ công (không Docker)
Backend

cd backend
go mod tidy
go run main.go
Truy cập tại: http://localhost:8080

Frontend

cd client
npm install
npm run dev
Truy cập tại: http://localhost:5173

🛠️ Triển khai Production
Build image:

docker-compose build
Chạy background:

docker-compose up -d
Kiểm tra logs:

docker-compose logs -f
📂 Các file quan trọng
backend/Dockerfile: Build Go API thành binary statically linked, chạy trên distroless.

client/Dockerfile: Build React app với Vite → serve qua Nginx.

nginx.conf: Reverse proxy cho frontend.

docker-compose.yml: Orchestrate toàn bộ hệ thống.

✅ Tính năng
CRUD Todo (thêm, xoá, đánh dấu hoàn thành).

Giao diện đẹp (Chakra UI).

Backend tối ưu (Go).

Dễ triển khai với Docker.

📜 License
ToanDev © 2025
