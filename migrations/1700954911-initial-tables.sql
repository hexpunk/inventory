CREATE TABLE containers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
        CHECK(trim(name) != ""),
    description TEXT NOT NULL DEFAULT ""
);

CREATE TABLE items (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
        CHECK(trim(name) != ""),
    quantity INTEGER NOT NULL
        CHECK(quantity > 0),
    description TEXT NOT NULL DEFAULT "",
    container_id INTEGER REFERENCES containers(id)
        ON UPDATE CASCADE
        ON DELETE SET NULL
);
