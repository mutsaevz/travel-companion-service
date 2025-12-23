package services

import (
	"log/slog"

	"github.com/mutsaevz/team-5-ambitious/internal/models"
	"github.com/mutsaevz/team-5-ambitious/internal/repository"
)

type UserService interface {
	Create(req *models.UserCreateRequest) (*models.User, error)

	List(filter models.UserFilter) ([]models.User, error)

	GetByID(id uint) (*models.User, error)

	Update(id uint, req models.UserUpdateRequest) (*models.User, error)

	Delete(id uint) error
}

type userService struct {
	repo   repository.UserRepository
	logger *slog.Logger
}

func NewUserService(userRepo repository.UserRepository, logger *slog.Logger) UserService {
	return &userService{
		repo:   userRepo,
		logger: logger,
	}
}

func (s *userService) Create(req *models.UserCreateRequest) (*models.User, error) {
	var user = models.User{
		Name:    req.Name,
		Phone:   req.Phone,
		Balance: req.Balance,
	}

	if err := s.repo.Create(&user); err != nil {
		s.logger.Error("error adding user",
			slog.Any("error", err),
		)
		return nil, err
	}

	return &user, nil
}

func (s *userService) List(filter models.UserFilter) ([]models.User, error) {
	users, err := s.repo.List(filter)
	if err != nil {
		s.logger.Error("user list error",
			slog.Any("error", err),
		)
		return nil, err
	}

	return users, nil
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Error("user output error",
			slog.Any("error", err),
		)
		return nil, err
	}

	return user, nil
}

func (s *userService) Update(id uint, req models.UserUpdateRequest) (*models.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		s.logger.Error("user not found",
			slog.Uint64("user_id", uint64(id)),
			slog.Any("error", err),
		)
		return nil, err
	}

	if req.Name != nil {
		user.Name = *req.Name
	}

	if req.Phone != nil {
		user.Phone = *req.Phone
	}

	if err := s.repo.Update(id, user); err != nil {
		s.logger.Error("error saving changes",
			slog.String("user_name", *req.Name),
			slog.Any("error", err),
		)
		return nil, err
	}

	return user, nil
}

func (s *userService) Delete(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		s.logger.Error("failed to delete user",
			slog.Uint64("user_id", uint64(id)),
			slog.Any("error", err),
		)
		return err
	}

	return nil
}
