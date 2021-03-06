CREATE TABLE IF NOT EXISTS "ingredients" (
  "id" TEXT NOT NULL UNIQUE,
  "name" TEXT NOT NULL UNIQUE,
  PRIMARY KEY ("id")
);
CREATE TABLE IF NOT EXISTS "recipes" (
  "id" TEXT NOT NULL UNIQUE,
  PRIMARY KEY ("id")
);
CREATE TABLE IF NOT EXISTS "recipe_details" (
  "id" TEXT NOT NULL UNIQUE,
  "recipe" TEXT references recipes(id) NOT NULL,
  "name" TEXT NOT NULL,
  "equipment" TEXT,
  "source" JSONB,
  "servings" INTEGER,
  "quantity" INTEGER,
  "unit" TEXT,
  "version" INTEGER NOT NULL,
  "is_latest_version" BOOLEAN DEFAULT FALSE,
  PRIMARY KEY ("id"),
  unique("recipe", "version")
);
-- https://stackoverflow.com/a/11014977
create unique index one_latest_revision_of_recipe on recipe_details (recipe)
where is_latest_version;
CREATE TABLE IF NOT EXISTS "recipe_sections" (
  "id" TEXT NOT NULL UNIQUE,
  "recipe_detail" TEXT references recipe_details(id) NOT NULL,
  "sort" INTEGER,
  "duration_timerange" JSONB,
  PRIMARY KEY ("id")
);
CREATE TABLE IF NOT EXISTS "recipe_section_instructions" (
  "id" TEXT NOT NULL UNIQUE,
  "section" TEXT references recipe_sections(id) NOT NULL,
  "sort" INTEGER,
  "instruction" TEXT,
  PRIMARY KEY ("id")
);
CREATE TABLE IF NOT EXISTS "recipe_section_ingredients" (
  "id" TEXT NOT NULL UNIQUE,
  "section" TEXT references recipe_sections(id) NOT NULL,
  "sort" INTEGER,
  --   ingredient can be an `ingredient` or a `recipe`
  "ingredient" TEXT references ingredients(id),
  "recipe" TEXT references recipes(id),
  "grams" numeric(10, 2),
  "amount" numeric(10, 2),
  "unit" TEXT,
  "adjective" TEXT,
  "original" TEXT,
  "optional" boolean default false,
  "substitutes_for" TEXT references recipe_section_ingredients(id),
  PRIMARY KEY ("id"),
  constraint check_ingredient check (
    ingredient is not null
    or recipe is not null
  )
);