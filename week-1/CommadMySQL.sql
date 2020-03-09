USE crawler;



show databases;
SHOW TABLES;

DROP TABLE crawler;

-- CREATE TABLE crawler
-- (
-- stt INT NOT NULL AUTO_INCREMENT,
-- url VARCHAR(255) NOT NULL,
-- title VARCHAR(255) NOT NULL,
-- author VARCHAR(255) NOT NULL,
-- datepublish VARCHAR(255) NOT NULL,
-- PRIMARY KEY (stt)
-- );

ALTER TABLE crawltable AUTO_INCREMENT = 1;
DELETE FROM crawltable WHERE stt>0;

INSERT INTO crawltable VALUES(null,'Url', 'Title', 'Author', 'DatePublish');
SELECT * FROM articles LIMIT 10;
SELECT * FROM crawltable LIMIT 2 OFFSET 4;


