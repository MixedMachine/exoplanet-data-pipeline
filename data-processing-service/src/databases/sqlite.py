import sqlite3


sqlite_client = sqlite3.connect("../data/sqlite/exoplanets.db")
sqlite_cursor = sqlite_client.cursor()
sqlite_cursor.execute("""
CREATE TABLE IF NOT EXISTS exoplanets (
    id TEXT PRIMARY KEY,
    name TEXT UNIQUE,
    mass REAL,
    radius REAL,
    period REAL,
    distance REAL,
    year_discovered INTEGER,
    method TEXT,
    confirmed BOOLEAN,
    planet_last_updated INTEGER,
    created_at TEXT,
    last_updated TEXT
);
""")

sqlite_client.commit()


def insert_into_table(table_name: str, **kwargs):
    field_names = ", ".join(kwargs.keys())
    placeholders = ", ".join(["?" for _ in kwargs])
    query = f"INSERT INTO {table_name} ({field_names}) VALUES ({placeholders});"
    values = tuple(kwargs.values())
    
    sqlite_cursor.execute(query, values)
    sqlite_client.commit()
