
CREATE TABLE IF NOT EXISTS users (
    id uuid UNIQUE NOT NULL,
    name varchar(255) NOT NULL,
    email varchar(255) UNIQUE NOT NULL,
    password varchar(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS authors (
    id uuid UNIQUE NOT NULL,
    name varchar(255) UNIQUE NOT NULL
);


CREATE TABLE IF NOT EXISTS categories (
    id uuid UNIQUE NOT NULL,
    name varchar(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS books (
    id uuid UNIQUE NOT NULL,
    title varchar(255),
    category varchar(255),
    subtitle varchar(255),
    description TEXT,
    release_date DATE,
    publisher varchar(255),
    language varchar(255),
    author varchar(255),
    page_number INTEGER,
    imagem varchar(255),
    rate float,
    owner varchar(255)
);
