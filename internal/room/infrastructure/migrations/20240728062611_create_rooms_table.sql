-- +goose Up
-- +goose StatementBegin
CREATE TABLE roomsvc_rooms (
    id INT(11) NOT NULL AUTO_INCREMENT,
    account_id INT(11) NOT NULL,
    name varchar(255) not null,
    model varchar(20) not null,
    description varchar(255),
    room_access_type TINYINT(1) NOT NULL,
    room_owner_visible TINYINT(1) NOT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE roomsvc_rooms;
-- +goose StatementEnd
