package services

import (
	clients "Payment_service/clients"
	pay "Payment_service/genproto/payment_service"
	reser "Payment_service/genproto/reservation_service"
	repositories "Payment_service/repositories"
	"context"
	"errors"
)

type PaymentService struct {
	pay.UnimplementedPaymentServiceServer
	reserClient clients.ReservationServiceClient
	repo        repositories.PaymentRepository
}

func NewPaymentService(repo repositories.PaymentRepository, reserClient clients.ReservationServiceClient) *PaymentService {
	return &PaymentService{repo: repo, reserClient: reserClient}
}

func (s *PaymentService) CreatePayment(ctx context.Context, req *pay.CreatePaymentRequest) (*pay.CreatePaymentResponse, error) {
	// Check reservation
	reserRes, err := s.reserClient.GetReservation(ctx, &reser.GetReservationRequest{Id: req.Payment.ReservationId})
	if err != nil {
		return nil, err
	}
	if reserRes.Reservation == nil {
		return nil, errors.New("reservation not found")
	}

	// Create payment
	payment, err := s.repo.CreatePayment(req)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *PaymentService) GetPayment(ctx context.Context, req *pay.GetPaymentRequest) (*pay.GetPaymentResponse, error) {
	payment, err := s.repo.GetPayment(req)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *PaymentService) ListPayments(ctx context.Context, req *pay.ListPaymentsRequest) (*pay.ListPaymentsResponse, error) {
	payments, err := s.repo.ListPayments(req)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (s *PaymentService) UpdatePayment(ctx context.Context, req *pay.UpdatePaymentRequest) (*pay.UpdatePaymentResponse, error) {
	payment, err := s.repo.UpdatePayment(req)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *PaymentService) DeletePayment(ctx context.Context, req *pay.DeletePaymentRequest) (*pay.DeletePaymentResponse, error) {
	res, err := s.repo.DeletePayment(req)
	if err != nil {
		return nil, err
	}
	return &pay.DeletePaymentResponse{Id: res.Id}, nil
}
