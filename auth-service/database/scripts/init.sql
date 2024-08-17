--
-- Структура таблицы `News`
--

CREATE TABLE news (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(256) NOT NULL,
  content TEXT NOT NULL
);

CREATE TABLE news_categories (
  news_id BIGINT NOT NULL,
  category_id BIGINT NOT NULL,
  PRIMARY KEY(news_id, category_id),
  FOREIGN KEY(news_id) 
    REFERENCES news(id)
    ON DELETE CASCADE
);

