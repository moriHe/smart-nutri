CREATE TABLE IF NOT EXISTS recipes(
	id SERIAL NOT NULL PRIMARY KEY,
	name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS ingredients(
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    synonym TEXT,
    category TEXT
);

CREATE TABLE IF NOT EXISTS recipes_ingredients(
	id SERIAL NOT NULL PRIMARY KEY,
    recipe_id INTEGER NOT NULL REFERENCES recipes (id),
    ingredient_id INTEGER NOT NULL REFERENCES ingredients (id)
);

INSERT INTO recipes (name) values ('Spaghetti');
INSERT INTO recipes (name) values ('Pizza');
INSERT INTO ingredients (name, category) values ('Tomaten', 'Obst');
INSERT INTO ingredients (name, category) values ('Knoblauch', 'Knolle');
INSERT INTO recipes_ingredients (recipe_id, ingredient_id) values (1, 1);
INSERT INTO recipes_ingredients (recipe_id, ingredient_id) values (1, 2);