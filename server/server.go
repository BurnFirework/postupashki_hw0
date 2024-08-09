package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "0.0.0.0:8081")
	if err != nil {
		log.Println(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		_, err = conn.Write([]byte("OK"))
		if err != nil {
			log.Println(err)
		}

		conn.Close()
	}
}
