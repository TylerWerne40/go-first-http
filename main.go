package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// listen on a port
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receiver the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg)
	}

	err = c.Close()
	if err != nil {
		return
	}
}

func client() {
	// connect to the server
	c, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {

		}
	}(c)
	msg := "Hello World!"
	fmt.Println("Sending: ", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	go server()
	go client()

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}
}
