
CREATE TABLE auth (
  uuid UUID PRIMARY KEY,
  refresh_token_signature TEXT NOT NULL UNIQUE
);


