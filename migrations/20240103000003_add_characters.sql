-- +goose Up
-- +goose StatementBegin
CREATE TABLE characters
(
    id               BIGSERIAL PRIMARY KEY,
    user_id          INT         NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    name             TEXT        NOT NULL,
    age              INT         NOT NULL,
    profession       TEXT        NOT NULL,
    burnout_level    INT         NOT NULL,
    motivation_level INT         NOT NULL,
    created_at       TIMESTAMPTZ NOT NULL,
    updated_at       TIMESTAMPTZ NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS characters;
-- +goose StatementEnd
