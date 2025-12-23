package repository

import (
	"errors"
	"log/slog"

	"github.com/mutsaevz/team-5-ambitious/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error

	List() ([]models.User, error)

	GetByID(id uint) (*models.User, error)

	Update(id uint, user *models.User) error

	Delete(id uint) error
}

type gormUserRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewUserRepository(db *gorm.DB, logger *slog.Logger) UserRepository {
	return &gormUserRepository{
		db:     db,
		logger: logger,
	}
}

func (r *gormUserRepository) Create(user *models.User) error {
	op := "repository.user.create"

	r.logger.Debug("db call",
		slog.String("op", op),
		slog.String("name", user.Name),
	)

	if err := r.db.Create(&user).Error; err != nil {
		r.logger.Error("db error",
			slog.String("op", op),
			slog.Any("error", err),
		)
		return err
	}

	return nil
}

func (r *gormUserRepository) List() ([]models.User, error) {
	op := "repository.user.list"

	r.logger.Debug("db call",
		slog.String("op", op),
	)

	var users []models.User

	if err := r.db.Find(&users).Error; err != nil {
		r.logger.Error("db error",
			slog.String("op", op),
			slog.Any("error", err),
		)
		return nil, err
	}

	return users, nil
}

func (r *gormUserRepository) GetByID(id uint) (*models.User, error) {
	op := "repository.user.get_by_id"

	r.logger.Debug("db call",
		slog.String("op", op),
	)

	var user *models.User

	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Warn("user not found", slog.Uint64("user_id", uint64(id)))
			return nil, ErrNotFound
		}

		r.logger.Error("db error",
			slog.String("op", op),
			slog.Any("error", err),
		)
		return nil, err
	}

	return user, nil
}

func (r gormUserRepository) Update(id uint, user *models.User) error {
	op := "repository.user.update"

	r.logger.Debug("db call",
		slog.String("op", op),
		slog.String("user_name", user.Name),
	)

	if err := r.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		r.logger.Error("db error",
			slog.String("op", op),
			slog.Any("error", err),
		)

		return err
	}

	return nil
}

func (r *gormUserRepository) Delete(id uint) error {
	op := "repository.user.delete"

	r.logger.Debug("db call",
		slog.String("op", op),
		slog.Uint64("id", uint64(id)),
	)

	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		r.logger.Error("db error",
			slog.Any("error", err),
		)
		return err
	}

	return nil
}
