DROP DATABASE IF EXISTS library_db;
CREATE DATABASE library_db;
USE library_db;

CREATE TABLE IF NOT EXISTS books (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY, 
    title varchar(64) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO books (id, title, created_at) VALUES (1, "Harry Potter", "2021-02-12 19:30:57");
INSERT INTO books (id, title, created_at) VALUES (2, "Lord of the Rings", "2021-02-12 19:31:57");
INSERT INTO books (id, title, created_at) VALUES (3, "The Great Gatsby", "2021-02-12 19:32:57");
INSERT INTO books (id, title, created_at) VALUES (4, "A book", "2021-02-12 19:33:57");
INSERT INTO books (id, title, created_at) VALUES (5, "Some book", "2021-02-12 19:34:57");
INSERT INTO books (id, title, created_at) VALUES (6, "Funny book", "2021-02-12 19:35:57");

INSERT INTO books (title) VALUES ("Interesting Book");

UPDATE books
SET title="Boring Book"
WHERE id=6; 

SELECT id, title as Title FROM books ; 

WHERE