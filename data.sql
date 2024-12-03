CREATE TABLE IF NOT  users (
 id INTEGER PRIMARY KEY AUTOINCREMENT,
 username text NOT NULL,
 email text NOT NULL,
 passwd text NOT NULL,
 expired_at TIME DEFAULT TIMESTAMP
);

CREATE TABLE sessions (
 id INTEGER Primary Key AUTOINCREMENT,
 id_user NOT NULL,
 expired_at TIME DEFAULT TIMESTAMP,
 created_at TIME DEFAULT TIMESTAMP,
 FOREIGN KEY (id_user) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);
