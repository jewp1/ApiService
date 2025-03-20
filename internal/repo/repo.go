package repo

import (
	"ApiService/internal/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"time"
)

type Task struct {
	Id          int       `json:"id, omitempty"`
	UserId      int       `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreateAt    time.Time `json:"createAt"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type repository struct {
	pool *pgxpool.Pool
}

type Repository interface {
	CreateTask(ctx context.Context, task Task) (int, error)
	GetTaskById(ctx context.Context, id int) (Task, error)
	UpdateTask(ctx context.Context, id int, task Task) (int, error)
	DeleteTask(ctx context.Context, id int) (int, error)
	CreateUser(ctx context.Context, user User) (int, error)
	GetTasksByUsername(ctx context.Context, username string) ([]Task, error)
	CheckUserExists(ctx context.Context, userId int) (bool, error)
	DeleteUser(ctx context.Context, userId int) (int, error)
	Close()
}

func NewRepo(ctx context.Context, cfg config.PostgreSQL) (Repository, error) {
	connString := fmt.Sprintf(
		`user=%s password=%s host=%s port=%d dbname=%s sslmode=%s 
        pool_max_conns=%d pool_max_conn_lifetime=%s pool_max_conn_idle_time=%s`,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSLMode,
		cfg.PoolSize,
		cfg.PoolConnLifeTime.String(),
		cfg.PoolMaxConnIdleTime.String(),
	)

	conf, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse db config")
	}

	conf.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe

	pool, err := pgxpool.NewWithConfig(ctx, conf)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create poll")
	}

	return &repository{pool}, nil
}

func (r *repository) CheckUserExists(ctx context.Context, userId int) (bool, error) {
	var exists bool
	err := r.pool.QueryRow(ctx, CheckUser, userId).Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "unable to check user exists")
	}
	return exists, nil
}

func (r *repository) CreateTask(ctx context.Context, task Task) (int, error) {
	var id int
	err := r.pool.QueryRow(ctx, CreateTask, task.UserId, task.Title, task.Description).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "unable to create task")
	}
	return id, nil
}

func (r *repository) CreateUser(ctx context.Context, user User) (int, error) {
	var id int
	err := r.pool.QueryRow(ctx, CreateUser, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "user exists")
	}
	return id, nil
}

func (r *repository) GetTasksByUsername(ctx context.Context, username string) ([]Task, error) {
	var tasks []Task
	rows, err := r.pool.Query(ctx, GetTasksByUsername, username)
	if err != nil {
		return tasks, errors.Wrap(err, "unable to get tasks")
	}
	defer rows.Close()
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.Status, &task.CreateAt); err != nil {
			return nil, errors.Wrap(err, "unable to scan task")
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *repository) GetTaskById(ctx context.Context, id int) (Task, error) {
	var task Task
	err := r.pool.QueryRow(ctx, GetTask, id).Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.Status, &task.CreateAt)
	if err != nil {
		return Task{}, errors.Wrap(err, "unable to get task")
	}
	if task.Id == 0 {
		return Task{}, errors.New("task not found")
	}
	return task, nil
}
func (r *repository) UpdateTask(ctx context.Context, id int, task Task) (int, error) {
	var taskId int
	err := r.pool.QueryRow(ctx, UpdateTask, id, task.Title, task.Description, task.Status, task.UserId).Scan(&taskId)
	if err != nil {
		return 0, errors.Wrap(err, "unable to update task")
	}
	return taskId, nil
}
func (r *repository) DeleteTask(ctx context.Context, id int) (int, error) {
	var taskId int
	err := r.pool.QueryRow(ctx, DeleteTask, id).Scan(&taskId)
	if err != nil {
		return 0, errors.Wrap(err, "unable to delete task")
	}
	return taskId, nil
}

func (r *repository) DeleteUser(ctx context.Context, userId int) (int, error) {
	var taskId int
	err := r.pool.QueryRow(ctx, DeleteUser, userId).Scan(&userId)
	if err != nil {
		return 0, errors.Wrap(err, "unable to delete user")
	}
	return taskId, nil
}

func (r *repository) Close() {
	if r.pool != nil {
		r.pool.Close()
	}
}
