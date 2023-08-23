CREATE TABLE users (
  id serial primary key,
  login varchar(100),
  password varchar(100),
  name varchar(50),
  role varchar(50) default 'user',
  created_at time without time zone default now(),
  updated_at time without time zone default now()
);
