CREATE OR REPLACE FUNCTION insert_task(
    task_title VARCHAR,
    task_description TEXT
)
RETURNS VOID AS $$
BEGIN
    INSERT INTO tasks (title, description) VALUES (task_title, task_description);
END;
$$ LANGUAGE plpgsql;
