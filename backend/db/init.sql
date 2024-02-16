-- .headers on
-- .mode column
-- .timer on

-- init.sql
CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  username TEXT DEFAULT NULL,
  email TEXT UNIQUE NOT NULL,
  password_hash TEXT NOT NULL,
  creation_time DATETIME DEFAULT (datetime('now', 'localtime'))
);

-- Add admin user
INSERT INTO
  users (username, email, password_hash)
VALUES
  (
    'admin',
    'admin@gmail.com',
    '713BFDA78870BF9D1B261F565286F85E97EE614EFE5F0FAF7C34E7CA4F65BACA'
  ); -- adminpass