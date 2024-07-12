package repositories

import (
	"Payment_service/genproto/payment_service"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db}
}

func (r *PaymentRepository) CreatePayment(req *payment_service.CreatePaymentRequest) (*payment_service.CreatePaymentResponse, error) {
	query := `
		INSERT INTO Payments (id, reservation_id, amount, payment_method, payment_status)
		VALUES ($1, $2, $3, $4, $5)
	`
	id := uuid.NewString()
	_, err := r.db.Exec(query, id, &req.Payment.ReservationId, req.Payment.Amount, req.Payment.PaymentMethod, req.Payment.PaymentStatus)
	if err != nil {
		return nil, err
	}

	req.Payment.Id = id

	return &payment_service.CreatePaymentResponse{Payment: req.Payment}, nil
}

func (r *PaymentRepository) GetPayment(req *payment_service.GetPaymentRequest) (*payment_service.GetPaymentResponse, error) {
	query := `
		SELECT id, reservation_id, amount, payment_method, payment_status, created_at, updated_at 
		FROM Payments WHERE id = $1 AND deleted_at IS NULL
	`
	row := r.db.QueryRow(query, req.Id)
	payment := &payment_service.Payment{}
	err := row.Scan(&payment.Id, &payment.ReservationId, &payment.Amount, &payment.PaymentMethod, &payment.PaymentStatus, &payment.CreatedAt, &payment.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("payment not found")
		}
		return nil, err
	}

	return &payment_service.GetPaymentResponse{Payment: payment}, nil
}

func (r *PaymentRepository) ListPayments(req *payment_service.ListPaymentsRequest) (*payment_service.ListPaymentsResponse, error) {
	query := `
		SELECT id, reservation_id, amount, payment_method, payment_status, created_at, updated_at
		FROM Payments WHERE reservation_id = $1 AND payment_method = $2 AND payment_status = $3 AND deleted_at IS NULL
		LIMIT $4 OFFSET $5
	`
	rows, err := r.db.Query(query, req.ReservationId, req.PaymentMethod, req.PaymentStatus, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*payment_service.Payment
	for rows.Next() {
		payment := &payment_service.Payment{}
		err := rows.Scan(&payment.Id, &payment.ReservationId, &payment.Amount, &payment.PaymentMethod, &payment.PaymentStatus, &payment.CreatedAt, &payment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return &payment_service.ListPaymentsResponse{Payments: payments}, nil
}

func (r *PaymentRepository) UpdatePayment(req *payment_service.UpdatePaymentRequest) (*payment_service.UpdatePaymentResponse, error) {
	query := `
		UPDATE Payments SET reservation_id = $1, amount = $2, payment_method = $3, payment_status = $4, updated_at = now()
		WHERE id = $5 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, req.Payment.ReservationId, req.Payment.Amount, req.Payment.PaymentMethod, req.Payment.PaymentStatus, req.Payment.Id)
	if err != nil {
		return nil, err
	}

	return &payment_service.UpdatePaymentResponse{Payment: req.Payment}, nil
}

func (r *PaymentRepository) DeletePayment(req *payment_service.DeletePaymentRequest) (*payment_service.DeletePaymentResponse, error) {
	query := `UPDATE Payments SET deleted_at = now() WHERE id = $1`
	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return &payment_service.DeletePaymentResponse{Id: req.Id}, nil
}
