-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS PRODUCTS( 
    id INTEGER PRIMARY KEY, 
    name VARCHAR(255),
    price INTEGER NOT NULL CHECK (price >= 0), 
    quantity INTEGER NOT NULL DEFAULT 0, 
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS PRODUCTS;
-- +goose StatementEnd
