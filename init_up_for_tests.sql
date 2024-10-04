CREATE TABLE groups (
    id BIGSERIAL PRIMARY KEY,
    name varchar(100) NOT NULL UNIQUE
);

CREATE TABLE songs (
    id BIGSERIAL PRIMARY KEY,
    group_id INT REFERENCES groups(id) NOT NULL,
    name varchar(100),
    link text,
    release_date varchar(20)
);

CREATE TABLE verses(
    id SERIAL PRIMARY KEY,
    song_id INT REFERENCES songs(id) ON DELETE CASCADE,
    verse_number INT NOT NULL,
    text TEXT NOT NULL
);

CREATE INDEX idx_groups_name ON groups(name);

CREATE UNIQUE INDEX idx_songs_id ON songs(id);

CREATE UNIQUE INDEX idx_verses_number ON verses(song_id, verse_number);