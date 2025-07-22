CREATE TABLE otp_codes (
    id SERIAL PRIMARY KEY,
    mobile VARCHAR(15) NOT NULL,
    code VARCHAR(6) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP WITH TIME ZONE
);

-- ایجاد ایندکس برای mobile جهت جستجو سریع‌تر
CREATE INDEX idx_otp_mobile ON otp_codes(mobile);
