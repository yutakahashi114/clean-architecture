create table restaurants(
    id         BIGSERIAL,
    name       VARCHAR(20) NOT NULL,
    content    VARCHAR(20) NOT NULL,
    status     VARCHAR(20),
    client_uid BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

create table restaurant_tags(
    id            BIGSERIAL,
    restaurant_id BIGINT NOT NULL,
    name          VARCHAR(20) NOT NULL,
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (restaurant_id) REFERENCES users (restaurants),
);

-- create table users(
--     id         BIGSERIAL,
--     first_name VARCHAR(20) NOT NULL,
--     last_name  VARCHAR(20) NOT NULL,
--     first_name_kana VARCHAR(20),
--     last_name_kana  VARCHAR(20),
--     version    BIGINT NOT NULL,
--     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMP,
--     PRIMARY KEY (id)
-- );

-- create table friends(
--     id         BIGSERIAL,
--     user_id    BIGINT NOT NULL,
--     friend_id  BIGINT NOT NULL,
--     version    BIGINT NOT NULL,
--     created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMP,
--     PRIMARY KEY (id),
--     FOREIGN KEY (user_id) REFERENCES users (id),
--     FOREIGN KEY (friend_id) REFERENCES users (id)
-- );
