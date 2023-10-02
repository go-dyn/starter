CREATE TABLE examples (
  id SERIAL PRIMARY KEY NOT NULL,
  title varchar(191) not null,
  content text not null,
  active int not null default 0,
  status int not null default 0,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  deleted_at timestamptz DEFAULT NULL
);
