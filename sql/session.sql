CREATE TABLE sessions(
session_id VARCHAR(255) PRIMARY KEY,
TTL TINYTEXT COMMENT "过期时间",
login_name VARCHAR(64) 	COMMENT "登录名"
);