CREATE TABLE IF NOT EXISTS "recipe"(
    "guid"           UUID PRIMARY KEY,
    "title"          VARCHAR NOT NULL DEFAULT '',
    "description"    VARCHAR NOT NULL DEFAULT '',
    "created_at"     TIMESTAMP DEFAULT now(),
    "updated_at"     TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS "ingredients"(
    "guid"           UUID PRIMARY KEY,
    "name"           VARCHAR NOT NULL UNIQUE,
    "dimension"      VARCHAR NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS "recipe_ingredients" (
    "recipe_id" UUID NOT NULL,
    "ingredient_id" UUID NOT NULL,
    "count" NUMERIC NOT NULL DEFAULT 0,
    CONSTRAINT "fk_recipe_id" FOREIGN KEY ("recipe_id") REFERENCES "recipe"("guid"),
    CONSTRAINT "fk_ingredient_id" FOREIGN KEY ("ingredient_id") REFERENCES "ingredients"("guid")
);


CREATE TABLE IF NOT EXISTS "cooking_steps"(
    "guid"           UUID PRIMARY KEY,
    "recipe_id"      UUID NOT NULL,
    "order_number"   INTEGER NOT NULL,
    "description"    VARCHAR NOT NULL DEFAULT '',
    CONSTRAINT "fk_recipe_id" FOREIGN KEY("recipe_id") REFERENCES "recipe"("guid")
);

CREATE INDEX IF NOT EXISTS "reciepe_ingredients_recipe_id_b_tree" ON "recipe_ingredients"("recipe_id");
CREATE INDEX IF NOT EXISTS "reciepe_ingredients_ingredient_id_b_tree" ON "recipe_ingredients"("ingredient_id");

CREATE INDEX IF NOT EXISTS "cooking_steps_recipe_id_b_tree" ON "cooking_steps"("recipe_id");
