
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `shots` (
    id                          BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    game_id                     VARCHAR(20) NOT NULL,
    player_id                   INT NOT NULL,
    shot_number                 SMALLINT NOT NULL,
    made                        TINYINT(1) NOT NULL,
    points                      TINYINT(4) NOT NULL,
    home                        TINYINT(1) NOT NULL,
    period                      TINYINT(4) NOT NULL,
    game_clock                  SMALLINT NOT NULL,
    shot_clock                  DOUBLE NOT NULL,
    dribbles                    SMALLINT NOT NULL,
    touch_time_seconds          DOUBLE NOT NULL,
    distance                    DOUBLE NOT NULL,
    points_type                 TINYINT(4) NOT NULL,
    closest_defender_player_id  INT NOT NULL,
    closest_defender_distance   DOUBLE NOT NULL,
    shot_type                   VARCHAR(50) NOT NULL,
    description                 VARCHAR(50) NOT NULL,
    zone                        VARCHAR(50) NOT NULL,
    location_x                  SMALLINT NOT NULL,
    location_y                  SMALLINT NOT NULL,
    UNIQUE(game_id, player_id, shot_number)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `shots`;
