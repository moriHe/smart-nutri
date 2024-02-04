SET timezone = 'UTC';

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'INSERT' THEN
        NEW.created_at = current_timestamp;
    END IF;
    
    NEW.updated_at = current_timestamp;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE IF NOT EXISTS familys(
    id SERIAL NOT NULL PRIMARY KEY,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp,
    name TEXT NOT NULL,
    code TEXT NOT NULL
);

CREATE TRIGGER update_users_timestamp_trigger
BEFORE INSERT OR UPDATE
ON familys
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TABLE IF NOT EXISTS users(
    id SERIAL NOT NULL PRIMARY KEY,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp,
    supabase_uid TEXT NOT NULL,
    active_family_id INTEGER REFERENCES familys (id)
);

CREATE TRIGGER update_users_timestamp_trigger
BEFORE INSERT OR UPDATE
ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TABLE IF NOT EXISTS invitations(
    id SERIAL NOT NULL PRIMARY KEY,
    created_at timestamptz default current_timestamp,
    token UUID UNIQUE NOT NULL,
    family_id INTEGER NOT NULL REFERENCES familys (id)
);

CREATE TABLE IF NOT EXISTS users_familys(
    id SERIAL NOT NULL PRIMARY KEY,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp,
    family_id INTEGER NOT NULL REFERENCES familys (id),
    user_id INTEGER NOT NULL REFERENCES users (id),
    user_role TEXT NOT NULL 
);

CREATE TRIGGER update_users_timestamp_trigger
BEFORE INSERT OR UPDATE
ON users_familys
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TABLE IF NOT EXISTS meals(
    id SERIAL NOT NULL PRIMARY KEY,
    meal TEXT
);

CREATE TABLE IF NOT EXISTS recipes(
	id SERIAL NOT NULL PRIMARY KEY,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp,
    family_id INTEGER NOT NULL REFERENCES familys (id),
	name VARCHAR(40) NOT NULL,
    default_portions FLOAT NOT NULL,
    default_meal INTEGER REFERENCES meals (id)

);

CREATE TRIGGER update_users_timestamp_trigger
BEFORE INSERT OR UPDATE
ON recipes
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TABLE IF NOT EXISTS mealplans(
    id SERIAL NOT NULL PRIMARY KEY,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp,
    family_id INTEGER NOT NULL REFERENCES familys (id),
    recipe_id INTEGER NOT NULL REFERENCES recipes (id),
    date timestamptz NOT NULL,
    meal INTEGER REFERENCES meals (id),
    portions FLOAT NOT NULL,
    is_shopping_list_item BOOLEAN NOT NULL
);

CREATE TRIGGER update_users_timestamp_trigger
BEFORE INSERT OR UPDATE
ON mealplans
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TABLE IF NOT EXISTS ingredients(
    id SERIAL NOT NULL PRIMARY KEY,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp,
    name TEXT NOT NULL,
    brands TEXT,
    url TEXT NOT NULL,
    code TEXT NOT NULL
);

CREATE TRIGGER update_users_timestamp_trigger
BEFORE INSERT OR UPDATE
ON ingredients
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

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
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp,
    recipe_id INTEGER NOT NULL REFERENCES recipes (id),
    ingredient_id INTEGER NOT NULL REFERENCES ingredients (id),
    amount_per_portion FLOAT NOT NULL,
    unit INTEGER NOT NULL REFERENCES units (id),
    market INTEGER NOT NULL REFERENCES markets (id),
    is_bio BOOLEAN NOT NULL
);

CREATE TRIGGER update_users_timestamp_trigger
BEFORE INSERT OR UPDATE
ON recipes_ingredients
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE TABLE IF NOT EXISTS shopping_list(
    id SERIAL NOT NULL PRIMARY KEY,
    created_at timestamptz default current_timestamp,
    updated_at timestamptz default current_timestamp,
    family_id INTEGER NOT NULL REFERENCES familys (id),
    mealplan_id INTEGER NOT NULL REFERENCES mealplans (id),
    recipes_ingredients_id INTEGER NOT NULL REFERENCES recipes_ingredients (id),
    market INTEGER NOT NULL REFERENCES markets (id),
    is_bio BOOLEAN NOT NULL
);

CREATE TRIGGER update_users_timestamp_trigger
BEFORE INSERT OR UPDATE
ON shopping_list
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

INSERT INTO markets ("id", "name") VALUES
(1, 'REWE'),
(2, 'EDEKA'),
(3, 'BIO_COMPANY'),
(4, 'WEEKLY_MARKET'),
(5, 'ALDI'),
(6, 'LIDL'),
(7, 'NONE');
SELECT setval(pg_get_serial_sequence('markets', 'id'), max(id)) FROM markets;

INSERT INTO units ("id", "name") VALUES
(1, 'GRAM'),
(2, 'MILLILITER'),
(3, 'TABLESPOON'),
(4, 'TEASPOON');
SELECT setval(pg_get_serial_sequence('units', 'id'), max(id)) FROM units;

INSERT INTO meals (id, meal) VALUES (1, 'BREAKFAST'), (2, 'LUNCH'), (3, 'DINNER'), (4, 'NONE');
SELECT setval(pg_get_serial_sequence('meals', 'id'), max(id)) FROM meals;
