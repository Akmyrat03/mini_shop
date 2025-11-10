-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS category (
    id SERIAL PRIMARY KEY,
    name_tk VARCHAR(255) NOT NULL,
    name_en VARCHAR(255) NOT NULL,
    name_ru VARCHAR(255) NOT NULL,  
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS category;
-- +goose StatementEnd
