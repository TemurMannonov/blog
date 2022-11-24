CREATE TABLE IF NOT EXISTS "users"(
    "id" SERIAL PRIMARY KEY,
    "first_name" VARCHAR(30) NOT NULL,
    "last_name" VARCHAR(30) NOT NULL,
    "phone_number" VARCHAR(20) UNIQUE,
    "email" VARCHAR(50) NOT NULL UNIQUE,
    "gender" VARCHAR(10) CHECK ("gender" IN('male', 'female')),
    "password" VARCHAR NOT NULL,
    "username" VARCHAR(30) UNIQUE,
    "profile_image_url" VARCHAR,
    "type" VARCHAR(255) CHECK ("type" IN('superadmin', 'user')) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "categories"(
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR(100) NOT NULL UNIQUE,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "posts"(
    "id" SERIAL PRIMARY KEY,
    "title" VARCHAR NOT NULL,
    "description" TEXT NOT NULL,
    "image_url" VARCHAR,
    "user_id" INTEGER NOT NULL REFERENCES users(id),
    "category_id" INTEGER NOT NULL REFERENCES categories(id),
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE,
    "views_count" INTEGER NOT NULL DEFAULT 0
);
CREATE INDEX IF NOT EXISTS posts_title_idx ON posts(title);

