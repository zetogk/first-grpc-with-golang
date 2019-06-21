package main

import (
	"context"
	"log"
	//"os"
	"time"

	// Llamamos el paquete de gRPC
	"google.golang.org/grpc"
	
	// Llamamos el compilado que nos generó protoc
	pb "ejemplo_servicio/definicionservicio/notificador"
)

const (
	address     = "localhost:50051" // Definimos por que host y puerto nos comunicamos
)

func main() {
	// configurar la conexión al servidor
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error en la conexión: %v", err)
	}
	defer conn.Close()
	c := pb.NewNotificadorClient(conn)

	// Enviamos la petición al servidor
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.EnviarCorreo(ctx, &pb.HelloRequest{Destino: "zetogk@gmail.com", Asunto: "Prueba", Mensaje: "Este es un mensaje de prueba"})
	if err != nil {
		log.Fatalf("Error en la petición: %v", err)
	}
	// Imprimimos la respuesta
	log.Println("¿El correo fue enviado?: ", r.Enviado)
}
