package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"

	"tcpserver/server"
)

func main() {
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
