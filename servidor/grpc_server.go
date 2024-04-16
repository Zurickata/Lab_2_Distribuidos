package server

import (
    "context"
    "log"
    "time"

    "grpc_municion/proto"

    "google.golang.org/grpc"
)

// MunicionServer representa el servidor gRPC de la Tierra
type MunicionServer struct {
    proto.UnimplementedMunicionServer
}

// SolicitarM implementa el método gRPC para solicitar munición
func (s *MunicionServer) SolicitarM(ctx context.Context, req *proto.SolicitudMunicionRequest) (*proto.SolicitudMunicionResponse, error) {
    // Lógica para procesar la solicitud de munición
    log.Printf("Recibida solicitud de munición de equipo %s: AT %d, MP %d", req.GetIdEquipo(), req.GetCantidadAt(), req.GetCantidadMp())

    // Simulamos un procesamiento
    time.Sleep(2 * time.Second)

    // Se puede implementar la lógica para verificar el inventario y responder adecuadamente

    return &proto.SolicitudMunicionResponse{Aprobado: true}, nil
}
