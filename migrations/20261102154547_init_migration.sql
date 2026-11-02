-- +goose Up
-- +goose StatementBegin
CREATE TABLE roles (
    SERIAL PRIMARY KEY id,
    VARCHAR(15) UNIQUE NOT NULL role_name DEFAULT "student"
);
INSERT INTO roles (role_name) VALUES
    ("student"),
    ("teacher"),
    ("admin");
CREATE TABLE users (
    SERIAL PRIMARY KEY id, 
    VARCHAR(50) first_name NOT NULL,
    VARCHAR(50) last_name NOT NULL,
    VARCHAR(100) email NOT NULL UNIQUE,
    TIMESTAMP created_at DEFAULT CURRENT_TIMESTAMP,
    TEXT password_hash NOT NULL,
    INT role_id REFERENCES roles(id) ON DELETE CASCADE
);
CREATE TABLE courses(
    SERIAL PRIMARY KEY id,
    VARCHAR(100) course_name NOT NULL,
    TIMESTAMP created_at DEFAULT CURRENT_TIMESTAMP,
    TEXT descr,
    teacher_id INT REFERENCES users(id)
);
CREATE TABLE student_courses(
    student_id INT REFERENCES students(id) ON DELETE CASCADE,
    course_id INT REFERENCES cources(id) ON DELETE CASCADE
);
CREATE TABLE lessons(
    SERIAL PRIMARY KEY id,
    VARCHAR(100) lesson_name NOT NULL,
    TIMESTAMP created_at DEFAULT CURRENT_TIMESTAMP,
    VARCHAR(100) name,
    course_id INT REFERENCES cources(id) ON DELETE CASCADE
);
CREATE TABLE tests(
    SERIAL PRIMARY KEY id,
    VARCHAR(100) test_name NOT NULL,
    TIMESTAMP created_at DEFAULT CURRENT_TIMESTAMP,
    course_id INT REFERENCES cources(id) ON DELETE CASCADE
);
CREATE TABLE question_types (
    SERIAL PRIMARY KEY id,
    VARCHAR(15) UNIQUE NOT NULL type_name DEFAULT "text"
);
INSERT INTO question_types (type_name) VALUES
    ("single"),
    ("multi"),
    ("text"),
    ("file");
CREATE TABLE questions(
    SERIAL PRIMARY KEY id,
    tyoe_id INT REFERENCES types(id) ON DELETE CASCADE
    TEXT question NOT NULL,
    test_id INT REFERENCES tests(id) ON DELETE CASCADE
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- +goose StatementEnd