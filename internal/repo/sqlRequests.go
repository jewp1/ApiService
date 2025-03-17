package repo

const (
	CreateTask         = `INSERT INTO tasks (user_id, title, description) VALUES ($1, $2, $3) RETURNING id`
	GetTask            = `SELECT * FROM tasks WHERE id = $1`
	GetTasksByUsername = `
	SELECT tasks.*
	FROM tasks
	JOIN users ON tasks.user_id = users.id
	WHERE users.username = $1;`
	UpdateTask = `
	UPDATE tasks
	SET 
    title = COALESCE(NULLIF($2, ''), title), 
    description = COALESCE(NULLIF($3, ''), description), 
    status = COALESCE(NULLIF($4, ''), status),
    user_id = COALESCE(NULLIF($5, 0), user_id)
	WHERE id = $1
	RETURNING id;`
	DeleteTask = `DELETE FROM tasks WHERE id = $1 RETURNING id;`
	DeleteUser = `DELETE FROM users WHERE id = $1 returning id;`
	CreateUser = `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	CheckUser  = `SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)`
)
