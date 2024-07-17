CREATE TABLE users (
                         id   BIGSERIAL PRIMARY KEY,
                         name text      NOT NULL,
                         display_name text,
                         email text,
                         picture text,
                         uuid   uuid     NOT NULL,
                         title text,
                         description text,
                         labels text[],
                         annotations text[],
                         tags text[],
                         links_id int[],
                         created_at timestamp with time zone DEFAULT now()
);



