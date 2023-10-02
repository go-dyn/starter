CREATE TABLE comments (
  id serial PRIMARY KEY NOT NULL,
  authorable_id serial NOT NULL,
  authorable_type varchar(191) NOT NULL,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);
