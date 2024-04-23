package main

import (
    pb "github.com/Zurickata/Lab_2_Distribuidos/proto"
    "context"
    "fmt"
    "google.golang.org/grpc"
    "net"
    "strconv"
    "time"
	"sync"
)

// MunicionServer representa el servidor gRPC de la Tierra
type MunicionServer struct {
    pb.UnimplementedMunicionServiceServer
    availableAt int32
	availableMp int32
	mutex sync.Mutex
}

// RequestMunicion implementa el método gRPC para solicitar munición
func (s *MunicionServer) RequestMunicion(ctx context.Context, req *pb.MunicionRequest) (*pb.MunicionResponse, error) {
    // Bloquear el acceso concurrente a los contadores de munición
    s.mutex.Lock()
	defer s.mutex.Unlock()

    // Formatear la cadena y convertir los valores enteros a cadenas
	message := fmt.Sprintf("Recepción de solicitud desde equipo %s, %s AT y %s MP",
    strconv.Itoa(int(req.GetTeamId())),
    strconv.Itoa(int(req.GetAtCount())),
    strconv.Itoa(int(req.GetMpCount())))

    // Lógica para procesar la solicitud de munición
    if req.AtCount > s.availableAt || req.MpCount > s.availableMp {
        // Imprimir el mensaje formateado
        fmt.Println(message + "-- DENEGADA -- AT EN SISTEMA:" + strconv.Itoa(int(s.availableAt)) + " ; MP EN SISTEMA:" + strconv.Itoa(int(s.availableMp)))
		return &pb.MunicionResponse{Approved: false}, nil
	}

	// Restar la munición solicitada de los contadores disponibles
	s.availableAt -= req.AtCount
	s.availableMp -= req.MpCount

    // Imprimir el mensaje formateado
    fmt.Println(message + "-- APROBADA -- AT EN SISTEMA:" + strconv.Itoa(int(s.availableAt)) + " ; MP EN SISTEMA:" + strconv.Itoa(int(s.availableMp)))

    // Se puede implementar la lógica para verificar el inventario y responder adecuadamente
    return &pb.MunicionResponse{Approved: true,}, nil
}

func main() {
    // Inicializar el servidor con los contadores de munición en cero
	server := &MunicionServer{
		availableAt: 0,
		availableMp: 0,
	}

    // Crear un Listener gRPC
    conn, err := net.Listen("tcp", ":50051")
    if err != nil {
        fmt.Println("No se pudo crear la conexion TCP: " + err.Error())
        return
    }

    // Iniciar el servidor gRPC
    fmt.Println("Servidor en ejecución en el puerto 50051...")
    serv := grpc.NewServer()
    pb.RegisterMunicionServiceServer(serv, server)
    
    go func() {
		for {
			// Incrementar los contadores de munición cada 5 segundos
			time.Sleep(5 * time.Second)
            server.mutex.Lock()
			server.availableAt += 10
			server.availableMp += 5

			// Asegurarse de que los contadores no excedan los valores máximos
			if server.availableAt > 50 {
				server.availableAt = 50
			}
			if server.availableMp > 20 {
				server.availableMp = 20
			}
            server.mutex.Unlock()
            fmt.Println("NUEVA DATA: AT EN SISTEMA:" + strconv.Itoa(int(server.availableAt)) + " ; MP EN SISTEMA:" + strconv.Itoa(int(server.availableMp)))
		}
	}()


    if err = serv.Serve(conn); err != nil{
        fmt.Println("No se pudo levantar el servidor: " + err.Error())
        return
    } 
}