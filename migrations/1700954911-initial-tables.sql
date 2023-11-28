CREATE TABLE containers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
        CHECK(trim(name) != ""),
    description TEXT NOT NULL DEFAULT "",
    created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
    updated_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
    deleted_at TEXT
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
        ON DELETE SET NULL,
    created_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
    updated_at TEXT NOT NULL DEFAULT (strftime('%Y-%m-%dT%H:%M:%fZ', 'now')),
    deleted_at TEXT
);

CREATE INDEX items_container_id_index ON items(container_id);

CREATE TRIGGER update_containers AFTER UPDATE ON containers
    BEGIN
        UPDATE containers
            SET updated_at = strftime('%Y-%m-%dT%H:%M:%fZ', 'now')
            WHERE id = new.id;
    END;

CREATE TRIGGER update_items AFTER UPDATE ON items
    BEGIN
        UPDATE items
            SET updated_at = strftime('%Y-%m-%dT%H:%M:%fZ', 'now')
            WHERE id = new.id;
    END;
