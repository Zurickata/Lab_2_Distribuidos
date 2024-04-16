package main

import (
    "context"
    "fmt"
    "log"
    "grpc_municion/proto"

    "google.golang.org/grpc"
)

func main() {
    // Conexión al servidor gRPC
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    // Crear un cliente gRPC
    client := proto.NewMunicionClient(conn)

    // Ejemplo de solicitud al servidor
    resp, err := client.SolicitarM(context.Background(), &proto.SolicitudMunicionRequest{
        IdEquipo:    "Equipo1",
        CantidadAt:  20,
        CantidadMp:  10,
    })
    if err != nil {
        log.Fatalf("Error al enviar solicitud: %v", err)
    }

    // Manejar la respuesta del servidor
    if resp.GetAprobado() {
        fmt.Println("Solicitud aprobada por el servidor")
    } else {
        fmt.Println("Solicitud rechazada por el servidor")
    }
}
