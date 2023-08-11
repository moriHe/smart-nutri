CREATE TABLE IF NOT EXISTS recipes(
	id serial PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS ingredients(
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT,
    synonym TEXT,
    category TEXT
);

CREATE TABLE IF NOT EXISTS recipes_ingredients(
    id serial NOT NULL PRIMARY KEY,
    recipe_id INTEGER NOT NULL REFERENCES recipes (id),
    ingredient_id INTEGER NOT NULL REFERENCES ingredients (id)
);

INSERT INTO recipes (name) values ('Hello World');