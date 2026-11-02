-- +goose Up
-- +goose StatementBegin

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    role_name VARCHAR(15) UNIQUE NOT NULL DEFAULT 'student'
);

INSERT INTO roles (role_name) VALUES
    ('student'),
    ('teacher'),
    ('admin');

CREATE TABLE users (
    id SERIAL PRIMARY KEY, 
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    password_hash TEXT NOT NULL,
    role_id INT REFERENCES roles(id) ON DELETE CASCADE
);

CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    course_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    descr TEXT,
    teacher_id INT REFERENCES users(id)
);

CREATE TABLE student_courses (
    student_id INT REFERENCES users(id) ON DELETE CASCADE,
    course_id INT REFERENCES courses(id) ON DELETE CASCADE,
    PRIMARY KEY (student_id, course_id) 
);

CREATE TABLE lessons (
    id SERIAL PRIMARY KEY,
    lesson_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(100),
    course_id INT REFERENCES courses(id) ON DELETE CASCADE
);

CREATE TABLE tests (
    id SERIAL PRIMARY KEY,
    test_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    course_id INT REFERENCES courses(id) ON DELETE CASCADE 
);

CREATE TABLE question_types (
    id SERIAL PRIMARY KEY,
    type_name VARCHAR(15) UNIQUE NOT NULL DEFAULT 'text'
);

INSERT INTO question_types (type_name) VALUES
    ('single'),
    ('multi'),
    ('text'),
    ('file');

CREATE TABLE questions (
    id SERIAL PRIMARY KEY,
    type_id INT REFERENCES question_types(id) ON DELETE CASCADE, 
    question TEXT NOT NULL,
    test_id INT REFERENCES tests(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
