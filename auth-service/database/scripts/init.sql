--
-- Структура таблицы `News`
--

CREATE TABLE auth (
  uuid UUID PRIMARY KEY,
  refresh_token TEXT NOT NULL,
  ip INET NOT NULL,
  expiration TIMESTAMP
);


