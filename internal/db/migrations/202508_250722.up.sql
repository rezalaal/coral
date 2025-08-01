CREATE TABLE otps (
  mobile VARCHAR(20) PRIMARY KEY,
  code VARCHAR(10) NOT NULL,
  expires_at BIGINT NOT NULL
);

CREATE TABLE otp_requests (
  id SERIAL PRIMARY KEY,
  mobile VARCHAR(20),
  requested_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX idx_otp_mobile_time ON otp_requests (mobile, requested_at);
