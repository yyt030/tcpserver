package testcase

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"tcpserver/server"
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

func A() {
	f, err := os.Open("testdata/gbk")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	seq := make([]byte, 1)
	seq[0] = 0x0a
	s := bytes.Split(data, []byte(seq))

	// saver file
	rf, err := os.Create("testdata/result.utf8")
	if err != nil {
		panic(err)
	}
	defer rf.Close()

	limit := 0
	for _, ss := range s {
		msg := sendAndRecv(ss, len(ss))
		rf.Write(msg)
		rf.WriteString("\n")
		//sendAndRecv2(ss, len(ss))
		limit++
		if limit > 100000 {
			break
		}
	}

}

func sendAndRecv(m []byte, l int) []byte {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	log.Println("WriteLenMsg:", m, l)
	err = server.WriteLenMsg(conn, m, l)
	if err != nil {
		panic(err)
	}

	ll, err := server.ReadLenMsg(conn, 8)
	if err != nil {
		panic(err)
	}

	llen, _ := strconv.Atoi(string(ll))

	msg, err := server.ReadLenMsg(conn, llen)
	if err != nil {
		panic(err)
	}
	log.Printf("got message:[%s]\n", msg)
	return msg
}

func sendAndRecv2(m []byte, l int) {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	log.Println("WriteLenMsg:", m, l)
	conn.Write([]byte("0000abcd"))
}
