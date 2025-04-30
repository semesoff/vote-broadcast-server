-- Users
CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(255) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (username, password)
VALUES ('alice', 'hashed_password_1'),
       ('bob', 'hashed_password_2'),
       ('charlie', 'hashed_password_3');


-- Polls
CREATE TABLE IF NOT EXISTS polls
(
    id          SERIAL PRIMARY KEY,
    uid         UUID         NOT NULL UNIQUE DEFAULT gen_random_uuid(),
    title       VARCHAR(255) NOT NULL,
    creator_id  INTEGER      NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    poll_type   VARCHAR(20)  NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE     DEFAULT CURRENT_TIMESTAMP,
    max_options INTEGER      NOT NULL        DEFAULT 10 CHECK (max_options > 0 AND max_options < 10),
    CONSTRAINT valid_poll_type CHECK (poll_type IN ('single', 'multiple'))
);

CREATE INDEX idx_polls_uid ON polls (id);

INSERT INTO polls (title, creator_id, poll_type, max_options)
VALUES ('Favorite Color?', 1, 'single', 3),
       ('Hobbies?', 2, 'multiple', 5);


-- Poll options
CREATE TABLE IF NOT EXISTS poll_options
(
    id          SERIAL PRIMARY KEY,
    poll_id     INTEGER      NOT NULL REFERENCES polls (id) ON DELETE CASCADE,
    option_text VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_poll_options_poll_id ON poll_options (poll_id);

INSERT INTO poll_options (poll_id, option_text)
VALUES (1, 'Red'),
       (1, 'Blue'),
       (1, 'Green'),
       (2, 'Reading'),
       (2, 'Gaming'),
       (2, 'Hiking'),
       (2, 'Cooking');


-- Votes
CREATE TABLE IF NOT EXISTS votes
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    poll_id    INTEGER NOT NULL REFERENCES polls (id) ON DELETE CASCADE,
    option_id  INTEGER NOT NULL REFERENCES poll_options (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, poll_id, option_id)
);

CREATE INDEX idx_votes_user_id ON votes (user_id);
CREATE INDEX idx_votes_poll_id ON votes (poll_id);

INSERT INTO votes (user_id, poll_id, option_id)
VALUES (2, 1, 1);
INSERT INTO votes (user_id, poll_id, option_id)
VALUES (3, 1, 3);

INSERT INTO votes (user_id, poll_id, option_id)
VALUES (1, 2, 4),
       (1, 2, 5),
       (3, 2, 6);
