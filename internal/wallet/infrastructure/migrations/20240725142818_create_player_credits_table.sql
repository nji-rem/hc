-- +goose Up
-- +goose StatementBegin
CREATE TABLE walletsvc_users (
    id int(11) NOT NULL AUTO_INCREMENT,
    account_id int(11) NOT NULL,
    credits int(11) NOT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE walletsvc_users;
-- +goose StatementEnd
