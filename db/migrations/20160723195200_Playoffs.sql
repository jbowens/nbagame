-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE `games` ADD COLUMN `playoffs` TINYINT(1) NOT NULL AFTER season;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE `games` DROP COLUMN `playoffs`;
