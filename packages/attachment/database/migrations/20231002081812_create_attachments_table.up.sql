CREATE TABLE attachments (
  id serial PRIMARY KEY NOT NULL,
  authorable_id serial NOT NULL,
  authorable_type varchar(191) NOT NULL,
  type int not null default 0,
  name varchar(191) not null,
  path varchar(191) not null,
  size int not null default 0,
  ext varchar(10) not null,
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL
);
