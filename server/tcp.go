package server

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"tcpserver/config"
	"tcpserver/conv"
)

func NewTcpServer(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("start tcp server errror", err)
	}
	defer listener.Close()
	log.Println("tcp server starting on ", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error:%v from %s\n", err, conn.RemoteAddr())
			continue
		}

		go HandleConn(conn)
	}
}

func HandleConn(conn net.Conn) {
	log.Printf("--- start handle connection from %v ---", conn.RemoteAddr())
	defer conn.Close()

	// Read header
	//log.Println("<<< read header message")
	header, err := ReadLenMsg(conn, config.HeaderLen)
	if err != nil {
		log.Printf("<<< return because read header message:%v, error:%v\n", header, err)
		return
	}
	// log.Printf("<<< header:%v,", header)
	bodyLen, err := strconv.Atoi(string(header))
	if err != nil || bodyLen <= 0 {
		log.Printf("<<< return because got header message:%v, bodyLen:%d, error:%v", string(header), bodyLen, err)
		return
	}

	// Read body
	// log.Println("<<< read body message")
	body, err := ReadLenMsg(conn, bodyLen)
	if err != nil || len(body) == 0 {
		log.Printf("<<< return because got body message:%v, error:%v, expected bodyLen:%v", body, err, bodyLen)
		return
	}
	log.Printf("<<< got message bodyLen:%d, body:%v\n", bodyLen, body)

	// Convert message
	resp, err := conv.ConvertMsg(body, bodyLen)
	if err != nil {
		log.Printf("<<< return because conv message:%v, error:%v\n", resp, err)
		return
	}

	// TODO:
	// Parse request message
	// Check token value
	// Build resp message

	// Write resp message
	w := bufio.NewWriter(conn)
	defer w.Flush()
	err = WriteLenMsg(w, resp, len(resp))
	if err != nil {
		log.Printf(">>> return because write resp message:%v, error:%v", resp, err)
		return
	}
	log.Printf(">>> write response message:%v, length:%d, over", string(resp), len(resp))
}

// Read message of l length from r
func ReadLenMsg(c net.Conn, l int) ([]byte, error) {
	msg := make([]byte, l)
	total := 0
	for total < l {
		buf := make([]byte, l-total)
		c.SetReadDeadline(time.Now().Add(config.TimeoutDuration))
		n, err := c.Read(buf)
		if n >= 0 {
			//log.Printf("NNN, n:%v, l:%v, total:%v, buf:%v", n, l, total, string(buf))
			//copy(msg[total:], buf[:n])
			msg = append(msg[:total], buf[:n]...)
			total += n
		}

		if err != nil {
			if err == io.EOF {
				log.Println("<<< read EOF:", err)
				break
			}
			return nil, err
		}
	}
	// Check got message whether is complete
	if total != l {
		//log.Println("MMM", total, l, string(msg))
		return msg[:total], errors.New("got not complete message")
	}

	return msg, nil
}

// Write message of l length to w
func WriteLenMsg(w io.Writer, m []byte, l int) error {
	if l <= 0 {
		return nil
	}
	// Header
	w.Write([]byte(fmt.Sprintf("%08s", strconv.Itoa(l))))
	//log.Printf(">>> write header:%v, body:%v", fmt.Sprintf("%08s", strconv.Itoa(l)), m)

	// Write Body
	total := 0
	for total < l {
		n, err := w.Write(m[total:])
		if n > 0 {
			total += n
		}
		if err != nil {
			return err
		}
	}

	return nil
}
