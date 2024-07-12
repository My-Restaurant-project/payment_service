package main

import (
	"Payment_service/config"
	pb "Payment_service/genproto/payment_service"
	reser "Payment_service/genproto/reservation_service"
	"Payment_service/pkg"
	repositories "Payment_service/repositories"
	services "Payment_service/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	config := config.Load()

	db, err := pkg.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":"+config.URL_PORT)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", config.URL_PORT, err)
	}
	defer listener.Close()

	log.Println("Server started on port " + config.URL_PORT)

	s := grpc.NewServer()

	reserConn, err := grpc.NewClient(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to reservation service: %v", err)
	}
	defer reserConn.Close()

	reserClient := reser.NewReservationServiceClient(reserConn)

	repo := repositories.NewPaymentRepository(db)

	ps := services.NewPaymentService(repo, reserClient)

	pb.RegisterPaymentServiceServer(s, ps)

	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
