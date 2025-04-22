package network

import (
	"fmt"
	"go-network-ui/models"
	"net"
	"time"
)

func TestConnection(server, port string) models.Result {
	start := time.Now()
	address := net.JoinHostPort(server, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	duration := time.Since(start)

	status := "Success"
	if err != nil {
		status = fmt.Sprintf("Failed: %v", err)
	}
	if conn != nil {
		conn.Close()
	}

	return models.Result{
		Server:       server,
		Port:         port,
		Status:       status,
		ResponseTime: duration,
		Timestamp:    time.Now(),
	}

}
