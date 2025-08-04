CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS posts(
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL,
    content text NOT NULL,
    title text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);