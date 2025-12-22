package services

import (
	"log/slog"

	"github.com/mutsaevz/team-5-ambitious/internal/models"
	"github.com/mutsaevz/team-5-ambitious/internal/repository"
)

type CarService interface {
	Create(id uint, req models.CarCreateRequest) (*models.Car, error)

	List() ([]models.Car, error)

	GetByOwner(id uint) (*models.Car, error)

	GetByID(id uint) (*models.Car, error)

	Update(id uint, req models.CarUpdateRequest) (*models.Car, error)

	Delete(id uint) error
}

type carService struct {
	carRepo  repository.CarRepository
	userRepo repository.UserRepository
	logger   *slog.Logger
}

func NewCarService(carRepo repository.CarRepository, userRepo repository.UserRepository, logger *slog.Logger) CarService {
	return &carService{
		carRepo:  carRepo,
		userRepo: userRepo,
		logger:   logger,
	}
}

func (s *carService) Create(id uint, req models.CarCreateRequest) (*models.Car, error) {
	driver, err := s.userRepo.GetByID(id)
	if err != nil {
		s.logger.Error("Пользователь не найден", slog.Uint64("user_id", uint64(id)), slog.String("error", err.Error()))
		return nil, err
	}

	car := models.Car{
		OwnerID:  driver.ID,
		Brand:    req.Brand,
		CarModel: req.CarModel,
		Seats:    req.Seats,
	}

	if err := s.carRepo.Create(&car); err != nil {
		return nil, err
	}

	return &car, nil
}

func (s *carService) GetByOwner(id uint) (*models.Car, error) {
	car, err := s.carRepo.GetByOwner(id)
	if err != nil {
		s.logger.Error("Ошибка при получении автомобиля по владельцу", slog.Uint64("owner_id", uint64(id)), slog.String("error", err.Error()))
		return nil, err
	}

	return car, nil
}

func (s *carService) GetByID(id uint) (*models.Car, error) {
	car, err := s.carRepo.GetByID(id)
	if err != nil {
		s.logger.Error("Ошибка при получении автомобиля по ID", slog.Uint64("car_id", uint64(id)), slog.String("error", err.Error()))
		return nil, err
	}
	return car, nil
}

func (s *carService) List() ([]models.Car, error) {
	cars, err := s.carRepo.List()
	if err != nil {
		s.logger.Error("Ошибка при получении списка автомобилей", slog.String("error", err.Error()))
		return nil, err
	}
	return cars, nil
}

func (s *carService) Update(id uint, req models.CarUpdateRequest) (*models.Car, error) {
	car, err := s.carRepo.GetByID(id)
	if err != nil {
		s.logger.Error("Автомобиль не найден для обновления", slog.Uint64("car_id", uint64(id)), slog.String("error", err.Error()))
		return nil, err
	}

	if req.Brand != nil {
		car.Brand = *req.Brand
	}
	if req.CarModel != nil {
		car.CarModel = *req.CarModel
	}
	if req.Seats != nil {
		car.Seats = *req.Seats
	}

	updatedCar, err := s.carRepo.Update(car)
	if err != nil {
		s.logger.Error("Ошибка при обновлении автомобиля", slog.Uint64("car_id", uint64(id)), slog.String("error", err.Error()))
		return nil, err
	}

	return updatedCar, nil
}

func (s *carService) Delete(id uint) error {
	_, err := s.carRepo.GetByID(id)
	if err != nil {
		s.logger.Error("Автомобиль не найден для удаления", slog.Uint64("car_id", uint64(id)), slog.String("error", err.Error()))
		return err
	}

	if err := s.carRepo.Delete(id); err != nil {
		s.logger.Error("Ошибка при удалении автомобиля", slog.Uint64("car_id", uint64(id)), slog.String("error", err.Error()))
		return err
	}

	s.logger.Info("Автомобиль успешно удалён", slog.Uint64("car_id", uint64(id)))
	return nil
}
