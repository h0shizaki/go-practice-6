DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR (50) NOT NULL,
  age INT NOT NULL,
  profession VARCHAR (50) NOT NULL,
  friendly BOOLEAN NOT NULL
);

INSERT INTO users(name,age,profession,friendly) VALUES
  ( 'kevin', 35, 'waiter', true),
  ( 'angela', 21, 'concierge', true),
  ( 'alex', 26, 'zoo keeper', false),
  ( 'becky', 67, 'retired', false),
  ( 'kevin', 15, 'in school', true),
  ( 'frankie', 45, 'teller', true);