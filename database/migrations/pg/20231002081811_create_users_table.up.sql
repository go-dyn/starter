CREATE TABLE users (
  id serial PRIMARY KEY NOT NULL,
  identifier varchar(191) NOT NULL,
  password varchar(191) NOT NULL,
  name varchar(191) DEFAULT '' NOT NULL,
  user_name varchar(191) DEFAULT NULL,
  email varchar(191) DEFAULT NULL,
  active int default 0,
  status int default 0,
  attachmentable_id serial NOT NULL,
  attachmentable_type varchar(191) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  deleted_at timestamptz DEFAULT NULL
);

CREATE INDEX idx_users_created_at ON users (created_at);
CREATE INDEX idx_users_updated_at ON users (updated_at);