CREATE TABLE shorturl (
    id  text not null UNIQUE,
    url text not null, 
    userId integer,
    visits integer
);

CREATE TABLE users (
    id integer not null UNIQUE,
    email text not null,
    password text not null,
    cookieId text,
    telegramId integer
);
