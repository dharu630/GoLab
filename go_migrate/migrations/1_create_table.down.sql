
-- Example down.sql
-- Dropping the tasks table
DROP TABLE IF EXISTS tasks;

-- Log that the migration is being rolled back (does not persist beyond session)
DO $$ BEGIN
   RAISE NOTICE 'Rolling back the "Create tasks table" migration.';
END $$;
