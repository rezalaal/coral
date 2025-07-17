# استفاده از نسخه Go سازگار با پروژه
FROM golang:1.24.5-alpine

# اضافه کردن ابزارهای مفید (مثل git)
RUN apk add --no-cache git

# مسیر کاری اپلیکیشن
WORKDIR /app

# کپی فایل‌های ماژول
COPY go.mod go.sum ./

# دانلود وابستگی‌ها
RUN go mod download

# کپی کل سورس پروژه
COPY . .

# ساخت باینری در مسیر مشخص
RUN mkdir -p bin
RUN go build -o bin/server ./cmd/server

# اجرای باینری
CMD ["./bin/server"]
