USE socialdb;

CREATE TABLE IF NOT EXISTS followers(
    user_id INT NOT NULL,
    follower_id INT NOT NULL,
    PRIMARY KEY (user_id, follower_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT chk_not_self_follow CHECK (user_id <> follower_id)
) ENGINE=INNODB;
