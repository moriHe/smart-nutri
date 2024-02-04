-- pg_down.sql

-- Drop tables in reverse order to avoid foreign key constraints
DROP TABLE IF EXISTS shopping_list;
DROP TABLE IF EXISTS recipes_ingredients;
DROP TABLE IF EXISTS ingredients;
DROP TABLE IF EXISTS mealplans;
DROP TABLE IF EXISTS recipes;
DROP TABLE IF EXISTS meals;
DROP TABLE IF EXISTS users_familys;
DROP TABLE IF EXISTS invitations;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS familys;
DROP TABLE IF EXISTS markets;
DROP TABLE IF EXISTS units;

-- Drop triggers
DROP TRIGGER IF EXISTS update_users_timestamp_trigger ON recipes_ingredients;
DROP TRIGGER IF EXISTS update_users_timestamp_trigger ON shopping_list;
DROP TRIGGER IF EXISTS update_users_timestamp_trigger ON ingredients;
DROP TRIGGER IF EXISTS update_users_timestamp_trigger ON mealplans;
DROP TRIGGER IF EXISTS update_users_timestamp_trigger ON recipes;
DROP TRIGGER IF EXISTS update_users_timestamp_trigger ON users_familys;

-- Reset sequences if used
-- SELECT setval(pg_get_serial_sequence('table_name', 'id'), 1, false);

-- Drop functions
DROP FUNCTION IF EXISTS update_timestamp;

-- If you have other specific cleanup operations, include them here
