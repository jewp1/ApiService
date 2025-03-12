package service

import (
	"ApiService/internal/dto"
	"ApiService/internal/repo"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
)

type Service interface {
	CreateTask(ctx *fiber.Ctx) error
	GetTasks(ctx *fiber.Ctx) error
	GetTaskById(ctx *fiber.Ctx) error
	UpdateTask(ctx *fiber.Ctx) error
	DeleteTask(ctx *fiber.Ctx) error
}

type service struct {
	repo repo.Repository
	log  *zap.SugaredLogger
}

func NewService(repo repo.Repository, logger *zap.SugaredLogger) Service {
	return &service{
		repo: repo,
		log:  logger,
	}
}

func (s *service) CreateTask(ctx *fiber.Ctx) error {
	var req TaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		s.log.Error("Request Body Error", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}
	task := repo.Task{
		Title:       req.Title,
		Description: req.Description,
	}
	taskID, err := s.repo.CreateTask(task)
	if err != nil {
		s.log.Error("Create Task Error", zap.Error(err))
		return dto.InternalServerError(ctx)
	}

	resp := dto.Response{
		Status: "success",
		Data:   map[string]int{"task_id": taskID},
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
func (s *service) GetTasks(ctx *fiber.Ctx) error {
	tasks, err := s.repo.GetTasks()
	if tasks == nil {
		s.log.Error("Tasks are empty", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Tasks are empty")
	}
	resp := dto.Response{
		Status: "success",
		Data:   tasks,
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
func (s *service) GetTaskById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.log.Error("Get Task Id Error", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid id")
	}

	task, err := s.repo.GetTaskById(id)
	if err != nil {
		s.log.Error("Get Task Id Error", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Task not found")
	}
	resp := dto.Response{
		Status: "success",
		Data:   task,
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
func (s *service) UpdateTask(ctx *fiber.Ctx) error {
	var req TaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		s.log.Error("Request Body Error", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request")
	}
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.log.Error("Task Id Error", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid id")
	}

	task := repo.Task{
		Title:       req.Title,
		Description: req.Description,
	}

	taskId, err := s.repo.UpdateTask(id, task)
	if err != nil {
		s.log.Error("Update Task Error", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Task not found")
	}
	resp := dto.Response{
		Status: "success",
		Data:   taskId,
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
func (s *service) DeleteTask(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.log.Error("Delete Task Id Error", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid id")
	}
	taskId, err := s.repo.DeleteTask(id)
	if err != nil {
		s.log.Error("Delete Task Id Error", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Task not found")
	}
	resp := dto.Response{
		Status: "success",
		Data:   taskId,
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
