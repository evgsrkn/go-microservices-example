CREATE TABLE tasks (
  id serial primary key,
  name varchar(100),
  description text,
  project_id int not null,
  user_id int not null,
  status varchar(50)
);
