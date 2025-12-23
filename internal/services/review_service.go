package services

import (
	"errors"
	"log/slog"

	"github.com/mutsaevz/team-5-ambitious/internal/models"
	"github.com/mutsaevz/team-5-ambitious/internal/repository"
)

var (
	ErrTripNotCompleted     = errors.New("trip not completed")
	ErrUserNotPassenger     = errors.New("user is not a passenger in this trip")
	ErrReviewAlreadyPresent = errors.New("review already exists for this user and trip")
)

type ReviewService interface {
	Create(userID uint, tripID uint, req *models.Review) (*models.Review, error)

	// List() (*models.Review, error)

	// GetByID(id uint) (*models.Review, error)

	// Update(id uint, req *models.Review) (*models.Review, error)

	// Delete(id uint) error
}

type reviewService struct {
	reviewRepo repository.ReviewRepository
	tripRepo   repository.TripRepository
	logger     *slog.Logger
}

func NewReviewService(
	reviewRepo repository.ReviewRepository,
	tripRepo repository.TripRepository,
	logger *slog.Logger,
) ReviewService {
	return &reviewService{
		reviewRepo: reviewRepo,
		tripRepo:   tripRepo,
		logger:     logger,
	}
}

func (s *reviewService) Create(userID uint, tripID uint, req *models.Review) (*models.Review, error) {

	op := "service.review.create"

	trip, err := s.tripRepo.GetByID(tripID)
	if err != nil {
		return nil, err
	}

	if trip.TripStatus != "completed" {
		s.logger.Error("cannot review a trip that is not completed",
			slog.String("op", op),
			slog.Uint64("tripID", uint64(tripID)),
		)
		return nil, ErrTripNotCompleted
	}

	if trip.AvailableSeats == trip.TotalSeats {
		s.logger.Error("user was not a passenger in this trip",
			slog.String("op", op),
			slog.Uint64("userID", uint64(userID)),
			slog.Uint64("tripID", uint64(tripID)),
		)
		return nil, ErrUserNotPassenger
	}

	exists, err := s.reviewRepo.ExistsByTripAndUser(tripID, userID)

	if err != nil {
		s.logger.Error("error checking review existence", slog.String("op", op), slog.Any("error", err))
		return nil, err
	}

	if exists {
		s.logger.Error("review already exists for this user and trip",
			slog.String("op", op), slog.Uint64("userID", uint64(userID)), slog.Uint64("tripID", uint64(tripID)))
		return nil, ErrReviewAlreadyPresent
	}

	review := &models.Review{
		AuthorID: userID,
		TripID:   tripID,
		Rating:   req.Rating,
		Text:     req.Text,
	}

	if err := s.reviewRepo.Create(review); err != nil {
		s.logger.Error("error creating review", slog.String("op", op), slog.Any("error", err))
		return nil, err
	}

	avgRating, err := s.reviewRepo.GetAvgRatingByTrip(tripID)

	if err != nil {
		s.logger.Error("error calculating average rating", slog.String("op", op), slog.Any("error", err))
		return nil, err
	}

	if err := s.tripRepo.UpdateAvgRating(tripID, avgRating); err != nil {
		s.logger.Error("error updating trip average rating", slog.String("op", op), slog.Any("error", err))
		return nil, err
	}
	return review, nil
}
