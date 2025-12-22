package repository

import (
	"log/slog"

	"github.com/mutsaevz/team-5-ambitious/internal/models"
	"gorm.io/gorm"
)

type BookingRepository interface {
	Create(booking *models.Booking) error

	// List() ([]models.Booking, error)

	// GetByID(id uint) (*models.Booking, error)

	// Update(booking *models.Booking) (*models.Booking, error)

	// Delete(id uint) error
}

type gormBookingRepository struct {
	DB     *gorm.DB
	logger *slog.Logger
}

func NewBookingRepository(db *gorm.DB, logger *slog.Logger) BookingRepository {
	return &gormBookingRepository{
		DB:     db,
		logger: logger,
	}
}

func (r *gormBookingRepository) Create(booking *models.Booking) error {
	op := "repository.booking.create"

	r.logger.Debug("db call",
		slog.String("op", op),
		slog.Uint64("trip_id", uint64(booking.TripID)),
		slog.Uint64("passenger_id", uint64(booking.PassengerID)),
	)

	if err := r.DB.Create(booking).Error; err != nil {
		r.logger.Error("db error",
			slog.String("op", op),
			slog.Any("error", err),
		)
		return err
	}

	return nil
}
