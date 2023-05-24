
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    uid serial PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    user_password VARCHAR (50) NOT NULL,
    email VARCHAR (255) UNIQUE NOT NULL,
    created_on TIMESTAMP NOT NULL

);
DROP TABLE IF EXISTS movies;
CREATE TABLE movies (
    mid serial PRIMARY KEY,
    title VARCHAR (50)  NOT NULL,
    overview VARCHAR (500) NOT NULL,
    release_date DATE NOT NULL, 
    rating VARCHAR (50)
);

