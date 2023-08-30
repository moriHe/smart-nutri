CREATE TABLE IF NOT EXISTS familys(
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS meals(
    id SERIAL NOT NULL PRIMARY KEY,
    meal TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS recipes(
	id SERIAL NOT NULL PRIMARY KEY,
    family_id INTEGER NOT NULL REFERENCES familys (id),
	name TEXT NOT NULL,
    default_portions FLOAT NOT NULL,
    default_meal INTEGER REFERENCES meals (id)

);

CREATE TABLE IF NOT EXISTS mealplans(
    id SERIAL NOT NULL PRIMARY KEY,
    family_id INTEGER NOT NULL REFERENCES familys (id),
    recipe_id INTEGER NOT NULL REFERENCES recipes (id),
    date DATE NOT NULL,
    meal INTEGER REFERENCES meals (id),
    portions FLOAT NOT NULL
);

CREATE TABLE IF NOT EXISTS ingredients(
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    synonym TEXT,
    category TEXT
);

CREATE TABLE IF NOT EXISTS markets(
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS  units(
    id SERIAL NOT NULL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS recipes_ingredients(
	id SERIAL NOT NULL PRIMARY KEY,
    recipe_id INTEGER NOT NULL REFERENCES recipes (id),
    ingredient_id INTEGER NOT NULL REFERENCES ingredients (id),
    amount_per_portion FLOAT NOT NULL,
    unit INTEGER NOT NULL REFERENCES units (id),
    market INTEGER NOT NULL REFERENCES markets (id),
    is_bio BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS mealplans_shopping_list(
    id SERIAL NOT NULL PRIMARY KEY,
    family_id INTEGER NOT NULL REFERENCES familys (id),
    mealplan_id INTEGER NOT NULL REFERENCES mealplans (id),
    recipes_ingredients_id INTEGER NOT NULL REFERENCES recipes_ingredients (id)
);

--Static Data
INSERT INTO ingredients (id, name, category) VALUES (1, 'Tomaten', 'Obst'), (2, 'Knoblauch', 'Knolle');
SELECT setval(pg_get_serial_sequence('ingredients', 'id'), max(id)) FROM ingredients;

INSERT INTO markets ("id", "name") VALUES
(1, 'Rewe'),
(2, 'Edeka'),
(3, 'Bio Company'),
(4, 'Wochenmarkt'),
(5, 'Aldi'),
(6, 'Lidl');
SELECT setval(pg_get_serial_sequence('markets', 'id'), max(id)) FROM markets;

INSERT INTO units ("id", "name") VALUES
(1, 'GRAM'),
(2, 'MILLILITER'),
(3, 'TABLESPOON'),
(4, 'TEASPOON');
SELECT setval(pg_get_serial_sequence('units', 'id'), max(id)) FROM units;

INSERT INTO meals (id, meal) VALUES (1, 'BREAKFAST'), (2, 'LUNCH'), (3, 'DINNER');
SELECT setval(pg_get_serial_sequence('meals', 'id'), max(id)) FROM meals;

-- Generate sample data as in init_pg
INSERT INTO familys ("id", "name") VALUES
(1, 'Eberhart'),
(2, 'Krakauer');
SELECT setval(pg_get_serial_sequence('familys', 'id'), max(id)) FROM familys;

INSERT INTO recipes ("id", "family_id", "name", "default_portions", "default_meal") VALUES
(1, 1, 'Spaghetti', 1, 1),
(2, 2, 'Pizza', 2, 2);
SELECT setval(pg_get_serial_sequence('recipes', 'id'), max(id)) FROM recipes;

INSERT INTO recipes_ingredients (id, recipe_id, ingredient_id, amount_per_portion, unit, market, is_bio) VALUES 
(1, 1, 1, 100, 1, 1, true),
(2, 1, 2, 200, 1, 1, false);
SELECT setval(pg_get_serial_sequence('recipes_ingredients', 'id'), max(id)) FROM recipes_ingredients;

INSERT INTO mealplans (id, family_id, recipe_id, date, meal, portions) VALUES 
(1, 1, 1, TO_DATE('2023/08/22', 'YYYY/MM/DD'), 3, 2),
(2, 1, 2, TO_DATE('2023/08/22', 'YYYY/MM/DD'), 1, 1);
SELECT setval(pg_get_serial_sequence('mealplans', 'id'), max(id)) FROM mealplans;

INSERT INTO mealplans_shopping_list(id, family_id, mealplan_id, recipes_ingredients_id) VALUES
(1, 1, 1, 1),
(2, 1, 2, 2);
SELECT setval(pg_get_serial_sequence('mealplans_shopping_list', 'id'), max(id)) FROM mealplans_shopping_list;