-- +goose Up
-- +goose StatementBegin
CREATE TABLE accountsvc_accounts(
    id int(11) NOT NULL AUTO_INCREMENT,
    username varchar(255) not null,
    password varchar(255) not null,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE accountsvc_accounts;
-- +goose StatementEnd
