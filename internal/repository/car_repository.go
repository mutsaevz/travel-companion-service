package repository

import (
	"errors"
	"log/slog"

	"github.com/mutsaevz/team-5-ambitious/internal/models"
	"gorm.io/gorm"
)

type CarRepository interface {
	Create(car *models.Car) error

	List() ([]models.Car, error)

	GetByOwner(id uint) (*models.Car, error)

	Update(car *models.Car) (*models.Car, error)

	Delete(id uint) error

	GetByID(id uint) (*models.Car, error)
}

type gormCarRepository struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewCarRepository(db *gorm.DB, logger *slog.Logger) CarRepository {
	return &gormCarRepository{
		db:     db,
		logger: logger,
	}
}

func (r *gormCarRepository) Create(car *models.Car) error {

	r.logger.Info(
		"Создание нового автомобиля",
		slog.Uint64("owner_id", uint64(car.OwnerID)),
		slog.String("brand", car.Brand),
	)

	err := r.db.Create(car).Error

	if err != nil {
		r.logger.Error(
			"Ошибка при создании автомобиля",
			slog.String("error", err.Error()),
		)
		return err
	}
	r.logger.Info(
		"Автомобиль успешно создан",
		slog.Uint64("car_id", uint64(car.ID)),
	)

	return nil
}

func (r *gormCarRepository) GetByOwner(id uint) (*models.Car, error) {
	var car models.Car

	if err := r.db.Where("owner_id = ?", id).First(&car).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &car, nil
}
func (r *gormCarRepository) GetByID(id uint) (*models.Car, error) {
	r.logger.Info(
		"Запрос автомобиля по ID",
		slog.Uint64("car_id", uint64(id)),
	)

	var car models.Car

	if err := r.db.First(&car, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r.logger.Warn(
				"Автомобиль не найден",
				slog.Uint64("car_id", uint64(id)),
			)
			return nil, ErrNotFound
		}

		r.logger.Error(
			"Ошибка при получении автомобиля",
			slog.Uint64("car_id", uint64(id)),
			slog.String("error", err.Error()),
		)
		return nil, err
	}

	r.logger.Info(
		"Автомобиль успешно получен",
		slog.Uint64("car_id", uint64(car.ID)),
		slog.String("brand", car.Brand),
		slog.String("model", car.CarModel),
	)

	return &car, nil
}

func (r *gormCarRepository) List() ([]models.Car, error) {
	r.logger.Info("Запрос списка автомобилей")

	var cars []models.Car

	if err := r.db.Find(&cars).Error; err != nil {
		r.logger.Error(
			"Ошибка при получении списка автомобилей",
			slog.String("error", err.Error()),
		)
		return nil, err
	}

	r.logger.Info(
		"Список автомобилей успешно получен",
		slog.Int("count", len(cars)),
	)

	return cars, nil
}

func (r *gormCarRepository) Update(car *models.Car) (*models.Car, error) {
	r.logger.Info(
		"Обновление автомобиля",
		slog.Uint64("car_id", uint64(car.ID)),
	)

	if err := r.db.Save(car).Error; err != nil {
		r.logger.Error(
			"Ошибка при обновлении автомобиля",
			slog.Uint64("car_id", uint64(car.ID)),
			slog.String("error", err.Error()),
		)
		return nil, err
	}

	r.logger.Info(
		"Автомобиль успешно обновлён",
		slog.Uint64("car_id", uint64(car.ID)),
	)

	return car, nil
}

func (r *gormCarRepository) Delete(id uint) error {
	r.logger.Info(
		"Удаление автомобиля",
		slog.Uint64("car_id", uint64(id)),
	)

	if err := r.db.Delete(&models.Car{}, id).Error; err != nil {
		r.logger.Error(
			"Ошибка при удалении автомобиля",
			slog.Uint64("car_id", uint64(id)),
			slog.String("error", err.Error()),
		)
		return err
	}

	r.logger.Info(
		"Автомобиль успешно удалён",
		slog.Uint64("car_id", uint64(id)),
	)

	return nil
}
