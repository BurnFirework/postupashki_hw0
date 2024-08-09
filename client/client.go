package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatalln("Ошибка подключения:", err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalln("Ошибка чтения:", err)
	}

	message := string(buf[:n])
	if message == "OK" {
		fmt.Println(message)
	} else {
		fmt.Println("Получено другое сообщение:", message)
	}
}
