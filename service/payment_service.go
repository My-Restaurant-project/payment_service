package services

import (
	"context"
	"Payment_service/genproto/payment_service"
	"Payment_service/repositores"
)

type PaymentService struct {
	payment_service.UnimplementedPaymentServiceServer
	repo *repositories.PaymentRepository
}

func NewPaymentService(repo *repositories.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) CreatePayment(ctx context.Context, req *payment_service.CreatePaymentRequest) (*payment_service.CreatePaymentResponse, error) {
	payment, err := s.repo.CreatePayment(req)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *PaymentService) GetPayment(ctx context.Context, req *payment_service.GetPaymentRequest) (*payment_service.GetPaymentResponse, error) {
	payment, err := s.repo.GetPayment(req)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *PaymentService) ListPayments(ctx context.Context, req *payment_service.ListPaymentsRequest) (*payment_service.ListPaymentsResponse, error) {
	payments, err := s.repo.ListPayments(req)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (s *PaymentService) UpdatePayment(ctx context.Context, req *payment_service.UpdatePaymentRequest) (*payment_service.UpdatePaymentResponse, error) {
	payment, err := s.repo.UpdatePayment(req)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (s *PaymentService) DeletePayment(ctx context.Context, req *payment_service.DeletePaymentRequest) (*payment_service.DeletePaymentResponse, error) {
	_,err := s.repo.DeletePayment(req)
	if err != nil {
		return nil, err
	}
	return &payment_service.DeletePaymentResponse{Id: req.Id}, nil
}
