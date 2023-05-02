CREATE TABLE IF NOT EXISTS users (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	name VARCHAR(35) NOT NULL,
	email VARCHAR(254) NOT NULL UNIQUE,
    about TEXT DEFAULT '',
	avatar_id INT,
	password VARCHAR(128) NOT NULL
);

CREATE TABLE IF NOT EXISTS tag (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	name VARCHAR(35) NOT NULL,
    about TEXT DEFAULT '',
	color VARCHAR(10) NOT NULL
);

CREATE TABLE IF NOT EXISTS project (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	name VARCHAR(35) NOT NULL,
    about TEXT DEFAULT '',
	color VARCHAR(10) NOT NULL,
    is_private boolean NOT NULL
);

CREATE TABLE IF NOT EXISTS goal (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	name VARCHAR(35) NOT NULL,
    project_id INT NOT NULL REFERENCES project(id) ON DELETE CASCADE,
    description TEXT DEFAULT '',
	time_start TIMESTAMP NOT NULL,
    time_end TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS entry (
	id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    project_id INT REFERENCES project(id) ON DELETE CASCADE,
    description TEXT DEFAULT '',
	time_start TIMESTAMP NOT NULL,
    time_end TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS tag_entry (
	tag_id INT NOT NULL REFERENCES tag(id) ON DELETE CASCADE,
    entry_id INT NOT NULL REFERENCES entry(id) ON DELETE CASCADE,
	PRIMARY KEY (tag_id, entry_id)
);

INSERT INTO users (name, email, about, avatar_id, password) 
VALUES ('test', 'test', 'test', 0, '');