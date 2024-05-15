ALTER TABLE Musics ADD COLUMN link VARCHAR(30) NULL COMMENT 'Represents song link' AFTER file;
ALTER TABLE Musics ADD COLUMN image VARCHAR(30) NULL COMMENT 'Represents song image link' AFTER file;