CREATE TABLE
    users (
        id TEXT PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE,
        password_hash TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT (DATETIME ('now', 'localtime'))
    );

CREATE TABLE
    posts (
        id TEXT PRIMARY KEY,
        user_id TEXT NOT NULL,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT (DATETIME ('now', 'localtime')),
        FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
    );

CREATE TABLE
    comments (
        id TEXT PRIMARY KEY,
        user_id TEXT NOT NULL,
        post_id TEXT NOT NULL,
        content TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT (DATETIME ('now', 'localtime')),
        FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
        FOREIGN KEY (post_id) REFERENCES posts (id) ON UPDATE CASCADE ON DELETE CASCADE
        
    );

CREATE TABLE
    categories (id TEXT PRIMARY KEY, name TEXT NOT NULL UNIQUE);

CREATE TABLE
    post_categories (
        post_id TEXT NOT NULL,
        category_id TEXT NOT NULL,
        PRIMARY KEY (post_id, category_id),
        FOREIGN KEY (post_id) REFERENCES posts (id) ON UPDATE CASCADE ON DELETE CASCADE,
        FOREIGN KEY (category_id) REFERENCES categories (id) ON UPDATE CASCADE ON DELETE CASCADE
    );

CREATE TABLE
    likes (
        id TEXT PRIMARY KEY,
        user_id TEXT NOT NULL,
        post_id TEXT,
        comment_id TEXT,
        react_type TEXT,
        created_at TIMESTAMP DEFAULT (DATETIME ('now', 'localtime')),
        FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
        FOREIGN KEY (post_id) REFERENCES posts (id) ON UPDATE CASCADE ON DELETE CASCADE,
        FOREIGN KEY (comment_id) REFERENCES comments (id) ON UPDATE CASCADE ON DELETE CASCADE,
        CONSTRAINT unique_user_post_comment UNIQUE (user_id, post_id, comment_id),
        CHECK (
            (
                post_id IS NOT NULL
                AND comment_id IS NULL
            )
            OR (
                post_id IS NULL
                AND comment_id IS NOT NULL
            )
        )
    );

CREATE TABLE
    sessions (
        session_id TEXT PRIMARY KEY,
        user_id TEXT NOT NULL,
        expires_at TIMESTAMP NOT NULL,
        created_at TIMESTAMP DEFAULT (DATETIME ('now', 'localtime')),
        FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE
    );

CREATE INDEX idx_user_id_session ON sessions (user_id);

CREATE INDEX idx_post_categories_post_id ON post_categories (post_id);

CREATE INDEX idx_post_categories_category_id ON post_categories (category_id);

CREATE INDEX idx_likes_post_id ON likes (post_id);

CREATE INDEX idx_likes_comment_id ON likes (comment_id);