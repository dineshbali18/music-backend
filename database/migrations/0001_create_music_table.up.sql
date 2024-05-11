CREATE TABLE music (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'Reepresents unique id for the song',
    `song_name` varchar(30) NOT NULL COMMENT 'Represents name of the song',
    `song_file` MEDIUMBLOB NOT NULL COMMENT 'Represents file that contians binary data related to the song',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Represents timestamp at which the record is created ',
    PRIMARY KEY(id)
) Engine = InnoDB
DEFAULT CHARSET = utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT = 'Table to store the music related data';

ALTER TABLE music
    ADD UNIQUE KEY `song_name_idx`(song_name),
    ADD UNIQUE KEY `song_file_idx`(song_file);
