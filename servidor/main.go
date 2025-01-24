package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("error creando conexion")
		os.Exit(1)
	}

	fmt.Println("Servidor corriendo en puerto 8080")

	for {
		fmt.Println("Esperando conexion ...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print("error aceptando conexion")
			continue
		}

		fmt.Println("Conexion aceptada")

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buff := make([]byte, 1024)
		large, err := conn.Read(buff)
		if err != nil {
			break
		}

		msg := string(buff)[0:large]

		if msg == ".exit" {
			break
		}

		fmt.Println(msg)
	}
	fmt.Println("Conexion cerrada")
}
