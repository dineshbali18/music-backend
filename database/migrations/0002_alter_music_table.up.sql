ALTER TABLE Musics ADD COLUMN link VARCHAR(500) NULL COMMENT 'Represents song link' AFTER file;
ALTER TABLE Musics ADD COLUMN image VARCHAR(500) NULL COMMENT 'Represents song image link' AFTER file;