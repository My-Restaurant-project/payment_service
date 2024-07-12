package main

import (
	"Payment_service/config"
	pb "Payment_service/genproto/payment_service"
	"Payment_service/pkg"
	repositories "Payment_service/repositores"
	services "Payment_service/service"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Konfiguratsiyani yuklash
	config := config.Load()

	// Ma'lumotlar bazasiga ulanish
	db, err := pkg.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Belgilangan portda TCP listener yaratish
	listener, err := net.Listen("tcp", ":"+config.URL_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("Server started on port " + config.URL_PORT)

	// gRPC server yaratish
	s := grpc.NewServer()

	// Repository va xizmatlarni initialize qilish
	repo := repositories.NewPaymentRepository(db)
	ps := services.NewPaymentService(repo)

	// PaymentServiceServerni gRPC server bilan ro'yxatdan o'tkazish
	pb.RegisterPaymentServiceServer(s, ps)

	// gRPC serverda reflection xizmatini ro'yxatdan o'tkazish
	reflection.Register(s)

	// gRPC serverni boshlash
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
