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
	CreateUser(ctx *fiber.Ctx) error
	GetTasksByUsername(ctx *fiber.Ctx) error
	GetTaskById(ctx *fiber.Ctx) error
	UpdateTask(ctx *fiber.Ctx) error
	DeleteTask(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
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
		UserId:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
	}
	exists, err := s.repo.CheckUserExists(ctx.Context(), req.UserId)
	if err != nil {
		s.log.Error("Check User Error", zap.Error(err))
	}
	if !exists {
		s.log.Error("User Not Found")
		return dto.BadResponseError(ctx, dto.FieldIncorrect, "User Not Found")
	}
	taskID, err := s.repo.CreateTask(ctx.Context(), task)
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

func (s *service) CreateUser(ctx *fiber.Ctx) error {
	var req UserRequest
	if err := ctx.BodyParser(&req); err != nil {
		s.log.Error("Request Body Error", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}
	user := repo.User{
		Username: req.UserName,
		Password: req.Password,
	}
	UserID, err := s.repo.CreateUser(ctx.Context(), user)
	if err != nil {
		s.log.Error("User Exists", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldIncorrect, "User Exists")
	}
	resp := dto.Response{
		Status: "success",
		Data:   map[string]int{"user_id": UserID},
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (s *service) GetTasksByUsername(ctx *fiber.Ctx) error {
	username := ctx.Params("username")

	tasks, err := s.repo.GetTasksByUsername(ctx.Context(), username)
	if err != nil {
		s.log.Error("Get Tasks Error", zap.Error(err))
		return dto.NotFoundError(ctx, dto.NotFound, "User Not Found")
	}
	if tasks == nil {
		return dto.NotFoundError(ctx, dto.NotFound, "No tasks found")
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
		s.log.Error("Get Task Error", zap.Error(err))
		return dto.NotFoundError(ctx, dto.FieldBadFormat, "Bad id")
	}

	task, err := s.repo.GetTaskById(ctx.Context(), id)
	if err != nil {
		s.log.Error("Get Task Error", zap.Error(err))
		return dto.NotFoundError(ctx, dto.NotFound, "Task Not Found")
	}
	resp := dto.Response{
		Status: "success",
		Data:   task,
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
func (s *service) UpdateTask(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.log.Error("Update Task Error", zap.Error(err))
		return dto.NotFoundError(ctx, dto.FieldBadFormat, "Bad id")
	}
	var req TaskRequest
	if err := ctx.BodyParser(&req); err != nil {
		s.log.Error("Request Body Error", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}
	task := repo.Task{
		UserId:      req.UserId,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}
	exists, err := s.repo.CheckUserExists(ctx.Context(), req.UserId)
	if err != nil {
		s.log.Error("Check User Error", zap.Error(err))
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "Invalid request body")
	}
	if !exists {
		s.log.Error("User Not Found")
		return dto.BadResponseError(ctx, dto.NotFound, "User Not Found")
	}
	id, err = s.repo.UpdateTask(ctx.Context(), id, task)
	if err != nil {
		s.log.Error("Update Task Error", zap.Error(err))
		return dto.NotFoundError(ctx, dto.NotFound, "Task Not Found")
	}
	resp := dto.Response{
		Status: "success",
		Data:   id,
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
func (s *service) DeleteTask(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.log.Error("Delete Task Error", zap.Error(err))
		return dto.NotFoundError(ctx, dto.FieldBadFormat, "Bad id")
	}
	taskID, err := s.repo.DeleteTask(ctx.Context(), id)
	if err != nil {
		s.log.Error("Delete Task Error", zap.Error(err))
		return dto.NotFoundError(ctx, dto.NotFound, "Task Not Found")
	}
	resp := dto.Response{
		Status: "success",
		Data:   taskID,
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}

func (s *service) DeleteUser(ctx *fiber.Ctx) error {
	UserId, err := strconv.Atoi(ctx.Params("user_id"))
	if err != nil {
		s.log.Error("Delete User Error", zap.Error(err))
		return dto.NotFoundError(ctx, dto.FieldBadFormat, "Bad id")
	}
	id, err := s.repo.DeleteUser(ctx.Context(), UserId)
	if err != nil {
		s.log.Error("Delete User Error", zap.Error(err))
		return dto.NotFoundError(ctx, dto.NotFound, "User Not Found")
	}
	resp := dto.Response{
		Status: "success",
		Data:   id,
	}
	return ctx.Status(fiber.StatusOK).JSON(resp)
}
