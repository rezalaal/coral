# Cafe Menu Application

یک اپلیکیشن ساده و سریع منوی کافه/رستوران با زبان Go و دیتابیس PostgreSQL، همراه با HTMX برای رابط کاربری ساده و SEO-friendly.

---

## ویژگی‌ها

- ثبت‌نام و ورود با شماره موبایل و OTP (در آینده توسعه داده می‌شود)
- مدیریت رستوران توسط مدیر هر رستوران
- ایجاد و مدیریت دسته‌بندی‌ها و محصولات
- نمایش منو به مشتریان با سرعت بالا و بهینه‌شده برای موتورهای جستجو
- معماری ماژولار و تمیز با Clean Architecture
- استفاده از `golang-migrate` برای مدیریت دیتابیس

---

## شروع سریع

### پیش‌نیازها

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://golang.org/dl/) (برای توسعه محلی)

### راه‌اندازی

1. کلون کردن پروژه:
   ```bash
   git clone https://github.com/rezalaal/coral.git
   cd coral
اجرای سرویس‌ها با Docker Compose:

```
docker-compose up --build
```
### دسترسی به اپلیکیشن:

اپ روی http://localhost:8080

دیتابیس PostgreSQL روی پورت 5432

pgAdmin http://localhost:8081

### ساختار پروژه
```
/cmd/server — نقطه ورود برنامه

/internal — کدهای منطق دامنه (restaurant, user, product, etc.)

/migrations — فایل‌های SQL مهاجرت دیتابیس

/web/templates — قالب‌های HTML برای HTMX

/web/static — فایل‌های استاتیک (CSS, JS, تصاویر)
```
مدیریت دیتابیس
برای اجرای migration ها از ابزار golang-migrate استفاده می‌شود.

### دستورات کاربردی:

```
migrate -path migrations -database "postgres://cafeuser:cafepass123@localhost:5432/cafedb?sslmode=disable" up
migrate -path migrations -database "postgres://cafeuser:cafepass123@localhost:5432/cafedb?sslmode=disable" down
```
توسعه و همکاری
کدها بر اساس اصول Clean Architecture و DDD ساختاربندی شده‌اند.

لطفاً قبل از ارسال Pull Request، تست‌های مربوطه را اجرا کنید.

برای هماهنگی و بحث‌های بیشتر، Issue باز کنید یا مستقیماً تماس بگیرید.

تماس
ایمیل: rezalaal@gmail.com

لینکدین: linkedin.com/in/rezalaal

لایسنس
این پروژه تحت مجوز MIT منتشر شده است.
