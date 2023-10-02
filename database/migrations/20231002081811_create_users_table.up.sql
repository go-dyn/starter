CREATE TABLE users (
  id INTEGER PRIMARY KEY,
  identifier TEXT NOT NULL,
  password TEXT NOT NULL,
  name TEXT DEFAULT '' NOT NULL,
  user_name TEXT DEFAULT NULL,
  email TEXT DEFAULT NULL,
  active INTEGER default 0,
  status INTEGER default 0,
  attachmentable_id INTEGER NOT NULL,
  attachmentable_type TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL
);