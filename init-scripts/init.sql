-- Пользователи
CREATE TABLE IF NOT EXISTS users
(
    id         SERIAL PRIMARY KEY,
    username   VARCHAR(255) NOT NULL UNIQUE,
    password   VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Опросы
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

-- Для быстрого поиска по UUID
CREATE INDEX idx_polls_uid ON polls (uid);
-- Для быстрого поиска по создателю
CREATE INDEX idx_polls_creator_id ON polls (creator_id);

-- Варианты ответов на опросы
CREATE TABLE IF NOT EXISTS poll_options
(
    id          SERIAL PRIMARY KEY,
    poll_id     INTEGER      NOT NULL REFERENCES polls (id) ON DELETE CASCADE,
    option_text VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Для быстрого поиска вариантов по опросу
CREATE INDEX idx_poll_options_poll_id ON poll_options (poll_id);

-- Голоса пользователей за опросы
CREATE TABLE IF NOT EXISTS votes
(
    id         SERIAL PRIMARY KEY,
    user_id    INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    poll_id    INTEGER NOT NULL REFERENCES polls (id) ON DELETE CASCADE,
    option_id  INTEGER NOT NULL REFERENCES poll_options (id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, poll_id, option_id) -- Один пользователь не может проголосовать за один вариант в опросе дважды
);

CREATE INDEX idx_votes_user_id ON votes (user_id); -- Для быстрого поиска голосов пользователя
CREATE INDEX idx_votes_poll_id ON votes (poll_id);
-- Для быстрого поиска голосов в опросе
-- Для быстрого поиска по вариантам
CREATE INDEX idx_votes_option_id ON votes (option_id);

-- Участники опроса
CREATE TABLE poll_participants
(
    poll_id   INTEGER NOT NULL REFERENCES polls (id) ON DELETE CASCADE,
    user_id   INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (poll_id, user_id)
);

CREATE INDEX idx_poll_participants_poll_id ON poll_participants (poll_id); -- Для быстрого поиска участников опроса
CREATE INDEX idx_poll_participants_user_id ON poll_participants (user_id);
-- Для быстрого поиска опросов пользователя

-- Тестовые данные
    
-- Добавляем пользователей
INSERT INTO users (username, password)
VALUES ('alice', 'hashed_password_1'),
       ('bob', 'hashed_password_2'),
       ('charlie', 'hashed_password_3');

-- Создаем опросы
INSERT INTO polls (title, creator_id, poll_type, max_options)
VALUES ('Favorite Color?', 1, 'single', 3),
       ('Hobbies?', 2, 'multiple', 5);

-- Добавляем варианты для опросов
INSERT INTO poll_options (poll_id, option_text)
VALUES (1, 'Red'),
       (1, 'Blue'),
       (1, 'Green'),
       (2, 'Reading'),
       (2, 'Gaming'),
       (2, 'Hiking'),
       (2, 'Cooking');

-- Добавляем участников
INSERT INTO poll_participants (poll_id, user_id)
VALUES (1, 2),
       (1, 3),
       (2, 1),
       (2, 3);

-- Голосование в single_choice опросе (Favorite Color?, id=1)
INSERT INTO votes (user_id, poll_id, option_id)
VALUES (2, 1, 1); -- Bob голосует за Red
INSERT INTO votes (user_id, poll_id, option_id)
VALUES (3, 1, 3);
-- Charlie голосует за Green

-- Голосование в multiple_choice опросе (Hobbies?, id=2)
INSERT INTO votes (user_id, poll_id, option_id)
VALUES (1, 2, 4), -- Alice голосует за Reading
       (1, 2, 5), -- Alice голосует за Gaming
       (3, 2, 6);
-- Charlie голосует за Hiking

-- Проверки

-- Проверка опросов
SELECT *
FROM polls;

-- Проверка вариантов
SELECT *
FROM poll_options;

-- Проверка участников
SELECT *
FROM poll_participants;

-- Проверка голосов для Favorite Color? (id=1)
SELECT po.option_text, COUNT(v.id) as votes_count
FROM poll_options po
         LEFT JOIN votes v ON po.id = v.option_id
WHERE po.poll_id = 1
GROUP BY po.id, po.option_text;

-- Проверка голосов для Hobbies? (id=2)
SELECT po.option_text, COUNT(v.id) as votes_count
FROM poll_options po
         LEFT JOIN votes v ON po.id = v.option_id
WHERE po.poll_id = 2
GROUP BY po.id, po.option_text;


