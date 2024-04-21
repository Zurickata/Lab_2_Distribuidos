package main

import (
    pb "github.com/Zurickata/Lab_2_Distribuidos/proto"
    "context"
    "fmt"
    "google.golang.org/grpc"
    "math/rand"
    "time"
)

func main() {
    // Conexión al servidor gRPC
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        fmt.Println("No se pudo conectar con el servidor: " + err.Error())
    }
    defer conn.Close()

    // Crear un cliente gRPC
    serviceClient := pb.NewMunicionServiceClient(conn)

	for {
		// Crear una solicitud con cantidades aleatorias
        idTeam := int32(randomInRange(1, 4))
		atCount := int32(randomInRange(10, 20))
		mpCount := int32(randomInRange(5, 10))

        message := fmt.Sprintf("Solicitando %d AT y %d MP", atCount, mpCount)

		// Enviar la solicitud al servidor
		res, err := serviceClient.RequestMunicion(context.Background(), &pb.MunicionRequest{
			TeamId:  idTeam,
			AtCount: atCount,
			MpCount: mpCount,
		})
		if err != nil {
			fmt.Println("Error al enviar solicitud: " + err.Error())
			continue
		}

		// Manejar la respuesta del servidor
		if res.Approved {
			fmt.Println(message + "; Resolución: -- APROBADA -- ; Conquista Exitosa!, cerrando comunicacion")
			break // Salir del bucle si la solicitud fue aprobada
		} else {
			fmt.Println(message + "; Resolución: -- DENEGADA -- ; Reintentando en 3 segs...")
			time.Sleep(3 * time.Second) // Esperar 3 segundos antes de reintentar
		}
	}
}

// Función para generar un número aleatorio dentro de un rango
func randomInRange(min, max int) int {
	return min + rand.Intn(max-min+1)
}