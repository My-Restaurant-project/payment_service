package clients

import (
	reser "Payment_service/genproto/reservation_service"
	"context"
	"google.golang.org/grpc"
)

type ReservationServiceClient interface {
	GetReservation(ctx context.Context, req *reser.GetReservationRequest, opts ...grpc.CallOption) (*reser.GetReservationResponse, error)
}
