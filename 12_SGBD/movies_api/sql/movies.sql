CREATE TABLE IF NOT EXISTS movies (
id SERIAL PRIMARY KEY,
name VARCHAR(100) NOT NULL,
year INT NOT NULL,
genre VARCHAR(100) NOT NULL,
acquired DATE,
stock INT DEFAULT 10,
price DECIMAL(10,2) NOT NULL
);