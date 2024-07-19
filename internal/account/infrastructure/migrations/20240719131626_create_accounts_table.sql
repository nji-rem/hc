-- +goose Up
-- +goose StatementBegin
--- +goose up
CREATE TABLE accounts(
    id int(11) NOT NULL AUTO_INCREMENT,
    username varchar(255) not null,
    password varchar(255) not null,
    look varchar(255) not null,
    motto varchar(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE accounts;
-- +goose StatementEnd
