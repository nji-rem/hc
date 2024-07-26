-- +goose Up
-- +goose StatementBegin
CREATE TABLE profilesvc_profiles (
    id int(11) NOT NULL AUTO_INCREMENT,
    account_id int(11) NOT NULL,
    look VARCHAR(255) NOT NULL,
    motto VARCHAR(255) NOT NULL DEFAULT "",
    sex ENUM("m", "f") NOT NULL DEFAULT "m",
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE profilesvc_profiles;
-- +goose StatementEnd
