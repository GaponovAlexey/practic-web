CREATE TABLE
  users (
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
  )
CREATE TABLE
  todo_lists(
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL
  )
CREATE TABLE
  users_list(
    id SERIAL NOT NULL,
    user_id INT REFERENCES users (id) ON DELETE CASCADE,
    list_id INT REFERENCES todo_lists (id) ON DELETE CASCADE
  )
CREATE TABLE
  todo_items(
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    done BOOLEAN NOT NULL DEFAULT FALSE
  )
CREATE TABLE
  list_items(
    id SERIAL NOT NULL UNIQUE,
    list_id INT REFERENCES todo_lists (id) ON DELETE CASCADE,
    items_id INT REFERENCES todo_items (id) ON DELETE CASCADE
  )