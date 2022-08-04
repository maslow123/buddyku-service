CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "username" varchar(45) NOT NULL,
  "password" varchar(255) NOT NULL,
  "level" int DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

-- Publisher
INSERT INTO users (name, username, password) 
VALUES
('user 1', 'user1@gmail.com', '$2a$05$wQ8lYAdEw7ZzF3OSzWeCKee8wc0KWxbBqfJpNu.lb.f1rvuSyy/I2');

-- Company
INSERT INTO users (name, username, password, level) 
VALUES
('company 1', 'company1@gmail.com', '$2a$05$wQ8lYAdEw7ZzF3OSzWeCKee8wc0KWxbBqfJpNu.lb.f1rvuSyy/I2', 1);


CREATE TABLE "articles" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar(150) NOT NULL,
  "content" text NOT NULL,
  "point" int DEFAULT 0,
  "view" int DEFAULT 0,
  "created_by" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE "articles" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON DELETE CASCADE;

INSERT INTO articles (title, content, created_by)
VALUES
('Title 1', 'Content 1', 1);