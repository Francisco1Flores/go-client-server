package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.100.153:8080")
	if err != nil {
		fmt.Println("Error al conectarse al servidor")
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("Conexion al servidor exitosa")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese su nombre")
	scanner.Scan()
	user := scanner.Text()
	user = "[" + user + "]: "

	msg := ""

	go waitMessages(conn, user)

	for msg != ".exit" {
		fmt.Print("-> ")
		scanner.Scan()
		text := scanner.Text()
		msg = user + text

		if text == ".exit" {
			conn.Write([]byte(".exit"))
			break
		}
		conn.Write([]byte(msg))
	}
	conn.Close()
}

func waitMessages(conn net.Conn, user string) {
	defer conn.Close()
	buff := make([]byte, 1024)
	for {
		large, err := conn.Read(buff)
		if err != nil {
			break
		}

		msg := string(buff[:large])

		fmt.Println("<- " + msg)
		fmt.Print("->")
	}

}
