CREATE TABLE IF NOT EXISTS "comments"(
    "id" SERIAL PRIMARY KEY,
    "post_id" INTEGER NOT NULL REFERENCES posts(id),
    "user_id" INTEGER NOT NULL REFERENCES users(id),
    "description" TEXT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS "likes"(
    "id" SERIAL PRIMARY KEY,
    "post_id" INTEGER NOT NULL REFERENCES posts(id),
    "user_id" INTEGER NOT NULL REFERENCES users(id),
    "status" BOOLEAN NOT NULL,
    UNIQUE(post_id, user_id)
);