package services

import (
	"context"
	"log/slog"
	"time"

	"github.com/mutsaevz/team-5-ambitious/internal/repository"
)

type TripStatusWorker struct {
	repo   repository.TripRepository
	logger *slog.Logger
	tick   time.Duration
}

func NewTripStatusWorker(
	repo repository.TripRepository,
	logger *slog.Logger,
	tick time.Duration,
) *TripStatusWorker {
	return &TripStatusWorker{
		repo:   repo,
		logger: logger,
		tick:   tick,
	}
}

func (w *TripStatusWorker) Start(ctx context.Context) {
	ticker := time.NewTicker(w.tick)

	go func() {
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				w.logger.Info("trip status worker stopped")
				return

			case <-ticker.C:
				now := time.Now().UTC()
				if err := w.repo.UpdateTripStatuses(now); err != nil {
					w.logger.Error(
						"failed to update trip statuses",
						slog.Any("error", err),
					)
				}
			}
		}
	}()
}
