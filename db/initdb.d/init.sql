CREATE DATABASE IF NOT EXISTS golang_todo_api;
USE golang_todo_api;

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    description TEXT,
    status ENUM('TODO', 'IN_PROGRESS', 'DONE') NOT NULL,
    created_at datetime NOT NULL,
    deleted_at datetime DEFAULT NULL
);