create table sample (
  id INTEGER PRIMARY KEY NOT NULL, 
  name TEXT NOT NULL,
);

create table users (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  password TEXT NOT NULL,
  first_name TEXT NOT NULL,
  middle_name TEXT NULL,
  last_name TEXT NULL,
  created_at TEXT DEFAULT CURRENT_TIMESTAMP,
  updated_at TEXT DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_user_timestamp
AFTER UPDATE ON users 
BEGIN
  UPDATE users 
  SET updated_at = CURRENT_TIMESTAMP
  WHERE id NEW.id;
END;

alter table users 
rename column name to first_name;

alter table users 
add column middle_name TEXT NULL; 

alter table users 
add column last_name TEXT NULL;

create table user_logs (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  action TEXT NOT NULL,
  created_at INTEGER DEFAULT (unixepoch('now')),
  updated_at INTEGER DEFAULT (unixepoch('now'))
);

create table todos (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  title TEXT NOT NULL,
  description TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  deleted_at DATETIME DEFAULT NULL
);

create table wallets (
  id TEXT PRIMARY KEY NOT NULL,
  user_id INTEGER NOT NULL,
  balance INTEGER NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY (user_id) REFERENCES users(id)
);

create table addresses (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  user_id INTEGER NOT NULL,
  address TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,

  FOREIGN KEY (user_id) REFERENCES users(id)
);

create table products (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  price INTEGER NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

create table user_like_product(
  user_id INTEGER  NOT NULL,
  product_id TEXT  NOT NULL,
  primary key(user_id, product_id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);
