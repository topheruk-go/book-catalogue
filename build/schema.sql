CREATE TABLE book (
    isbn VARCHAR(13) PRIMARY KEY CHECK (LENGTH(isbn) = 13),
    title TEXT NOT NULL,
    author TEXT NOT NULL
);
CREATE TABLE author (
    id TEXT PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);
CREATE TABLE review (
    id TEXT PRIMARY KEY,
    full_name VARCHAR(50) NOT NULL
);