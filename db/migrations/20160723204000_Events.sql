-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
DROP TABLE IF EXISTS `gameplay_events`;

CREATE TABLE `events` (
    game_id             VARCHAR(20) NOT NULL,
    seq                 INTEGER NOT NULL,
    event_type          VARCHAR(255) NOT NULL,
    period              TINYINT(1) NOT NULL,
    score_home          SMALLINT,
    score_visitor       SMALLINT,
    period_time         SMALLINT NOT NULL,
    wall_clock          VARCHAR(255) NOT NULL,
    player1_id          INTEGER,
    player2_id          INTEGER,
    player3_id          INTEGER,
    home_description    VARCHAR(255),
    neutral_description VARCHAR(255),
    visitor_description VARCHAR(255),
    reference_id        BIGINT,
    PRIMARY KEY(game_id, seq)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `events`;
