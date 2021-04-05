package trip

import (
	"context"

	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/auth"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implements a trip service.
type Service struct {
	Logger *zap.Logger
}

// CreateTrip creates a trip.
func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	aid, err := auth.AccountIDFromContext(c)
	if err != nil {
		return nil, err
	}
	s.Logger.Info("create trip", zap.String("Start", req.Start.String()), zap.String("account_id", aid.String()))
	return nil, status.Error(codes.Unimplemented, "")
}

// GetTrip  get the trip data.
func (s *Service) GetTrip(c context.Context, t *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	return nil, nil
}

// GetTrips  get the trips data.
func (s *Service) GetTrips(c context.Context, in *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	return nil, nil
}

// UpdateTrip  udate the trip data.
func (s *Service) UpdateTrip(c context.Context, in *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	return nil, nil
}
