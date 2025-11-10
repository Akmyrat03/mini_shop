-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS product (
    id SERIAL PRIMARY KEY,
    category_id INT NOT NULL REFERENCES category(id),
    title_tk VARCHAR(255) NOT NULL,
    title_en VARCHAR(255) NOT NULL,
    title_ru VARCHAR(255) NOT NULL,
    description_ru TEXT NOT NULL,
    description_en TEXT NOT NULL,
    description_tk TEXT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS product;
-- +goose StatementEnd
