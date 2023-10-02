CREATE TABLE posts (
  id serial PRIMARY KEY NOT NULL,
  content varchar(191) NOT NULL,
  authorable_id serial NOT NULL,
  authorable_type varchar(191) NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);
