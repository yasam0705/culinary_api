DROP INDEX IF EXISTS "recipe_cooking_time_b_tree";

ALTER TABLE IF EXISTS "cooking_steps" DROP COLUMN IF EXISTS "cooking_time";
ALTER TABLE IF EXISTS "recipe" DROP COLUMN IF EXISTS "cooking_time";

DROP INDEX IF EXISTS "users_username_b_tree";

DROP TABLE IF EXISTS "users";