package server

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"
)

func TestNullMsg(*testing.T) {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatal("dial tcp server failed")
	}
	defer conn.Close()

	// case 1: send 0 header and 0 body
	msgHeader := "0000000000"
	for _, v := range msgHeader {
		conn.Write([]byte(string(v)))
	}

}

func TestNotCompleteMsg(*testing.T) {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatal("dial tcp server failed")
	}
	defer conn.Close()

	// send not complete message
	msg := "0000000412"
	conn.Write([]byte(string(msg)))
}

func TestConnEnd(*testing.T) {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatal("dial tcp server failed")
	}
	defer conn.Close()
}

func TestSendTimeout(*testing.T) {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatal("dial tcp server failed")
	}
	defer conn.Close()
	time.Sleep(time.Second * 6)
}

func TestSendTimeout2(*testing.T) {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		log.Fatal("dial tcp server failed")
	}
	defer conn.Close()

	msg := "0000000812345678"
	for _, v := range msg {
		//time.Sleep(time.Second * time.Duration(rand.Intn(2)))
		fmt.Println(">>>", string(v))
		conn.Write([]byte(string(v)))
	}
}
