// Package main Implementación del servidor para el servicio Notificador
package main

import (
	"context"
	"log"
	"net"

	// Llamamos el paquete de gRPC
	"google.golang.org/grpc"

	// Llamamos el compilado que nos generó protoc
	pb "github.com/zetogk/first-grpc-with-golang/notificador"
)

const (
	// Puerto por donde expondremos el servicio
	port = ":50051"
)

// server es usado para implementar Notificador.
type server struct{}


// EnviarCorreo implementar Notificador.EnviarCorreo
// El parámetro in, corresponde al mensaje CorreoRequest
func (s *server) EnviarCorreo(ctx context.Context, in *pb.CorreoRequest) (*pb.CorreoResponse, error) {
	log.Printf("Enviando correo: %v a %v", in.Asunto, in.Destino)
	// Generamos y retornamos la respuesta de tipo CorreoResponse
	return &pb.CorreoResponse{Enviado: true}, nil
}

func main() {
	// Exponemos nuestro servidor para que reciba llamados
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNotificadorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
