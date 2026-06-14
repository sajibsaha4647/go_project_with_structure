CREATE TABLE IF NOT EXISTS products (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255)   NOT NULL,
    description TEXT           NOT NULL,
    price       VARCHAR(50)    NOT NULL,
    image_url   TEXT           NOT NULL
);
