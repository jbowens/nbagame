
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `teams` (
    id                  INT NOT NULL PRIMARY KEY,
    abbreviation        VARCHAR(5) NULL,
    city                VARCHAR(255) NOT NULL,
    name                VARCHAR(255) NOT NULL,
    start_year          YEAR(4) NOT NULL,
    end_year            YEAR(4) NULL,
    games               INT NOT NULL,
    wins                INT NOT NULL,
    losses              INT NOT NULL,
    playoff_appearances INT NOT NULL,
    division_titles     INT NOT NULL,
    conference_titles   INT NOT NULL,
    league_titles       INT NOT NULL
);

CREATE TABLE `games` (
    id                    VARCHAR(20) NOT NULL PRIMARY KEY,
    season                CHAR(7) NOT NULL,
    home_team_id          INT NOT NULL,
    visitor_team_id       INT NOT NULL,
    status                TINYINT(4) NOT NULL,
    time                  TIMESTAMP NOT NULL,
    last_meeting_game_id  VARCHAR(20) NULL,
    length_minutes        INT NULL,
    attendance            INT NULL,
    lead_changes          INT NULL,
    times_tied            INT NULL,
    INDEX(home_team_id),
    INDEX(visitor_team_id)
);

CREATE TABLE `players` (
    id                INT NOT NULL PRIMARY KEY,
    team_id           INT NOT NULL,
    team_abbreviation VARCHAR(5) NOT NULL,
    first_name        VARCHAR(255) NULL,
    last_name         VARCHAR(255) NULL,
    roster_status     TINYINT(4) NOT NULL,
    career_start      YEAR(4) NOT NULL,
    career_end        YEAR(4) NULL,
    birthdate         DATE NULL,
    school            VARCHAR(255) NULL,
    country           VARCHAR(255) NULL,
    height            INT NULL,
    weight            INT NULL,
    season_experience INT NOT NULL,
    jersey            VARCHAR(255) NULL,
    position          VARCHAR(255) NULL,
    dleague           TINYINT(1) NOT NULL DEFAULT 0,
    INDEX(team_id),
    INDEX(team_abbreviation),
    INDEX(career_end)
);

CREATE TABLE `stats` (
    id                        INT NOT NULL AUTO_INCREMENT,
    seconds_played            INT NOT NULL,
    field_goals_made          INT NOT NULL,
    field_goals_attempted     INT NOT NULL,
    three_pointers_made       INT NOT NULL,
    three_pointers_attempted  INT NOT NULL,
    free_throws_made          INT NOT NULL,
    free_throws_attempted     INT NOT NULL,
    offensive_rebounds        INT NOT NULL,
    defensive_rebounds        INT NOT NULL,
    assists                   INT NOT NULL,
    steals                    INT NOT NULL,
    blocks                    INT NOT NULL,
    turnovers                 INT NOT NULL,
    personal_fouls            INT NOT NULL,
    points                    INT NOT NULL,
    plus_minus                INT NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE `player_game_stats` (
    player_id   INT NOT NULL,
    game_id     INT NOT NULL,
    stats_id    INT NOT NULL,
    PRIMARY KEY(player_id, game_id)
);

CREATE TABLE `team_game_stats` (
    team_id     INT NOT NULL,
    game_id     INT NOT NULL,
    stats_id    INT NOT NULL,
    PRIMARY KEY(team_id, game_id)
);

CREATE TABLE `officials` (
    id            INT NOT NULL PRIMARY KEY,
    first_name    VARCHAR(255) NULL,
    last_name     VARCHAR(255) NULL,
    jersey_number VARCHAR(3) NULL
);

CREATE TABLE `officiated` (
    game_id     INT NOT NULL,
    official_id INT NOT NULL,
    PRIMARY KEY(game_id, official_id)
);

CREATE TABLE `gameplay_events` (
    id                    INT NOT NULL,
    event_type            TINYINT NOT NULL,
    game_id               varchar(20) NOT NULL,
    period                TINYINT(4) NOT NULL,
    home_score            INT NULL,
    visitor_score         INT NULL,
    time                  TIMESTAMP NULL,
    period_time_remaining INT NULL,
    description           TEXT NULL,
    shot_attributes       INT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `teams`;
DROP TABLE `games`;
DROP TABLE `stats`;
DROP TABLE `player_game_stats`;
DROP TABLE `team_game_stats`;
DROP TABLE `players`;
DROP TABLE `officials`;
DROP TABLE `officiated`;
DROP TABLE `gameplay_events`;
