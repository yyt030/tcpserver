package server

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

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

	r := bufio.NewReader(conn)

	// Read header
	header, err := ReadLenMsg(r, config.HeaderLen)
	if err != nil {
		log.Printf("<<< return and read header:%v, error:%v\n", header, err)
		return
	}
	//log.Printf("<<< header:%v,", header)
	bodyLen, err := strconv.Atoi(string(header))
	if err != nil || bodyLen <= 0 {
		log.Printf("<<< return and got header message:%v, bodyLen:%d, error:%v", header, bodyLen, err)
		return
	}

	// Read body
	body, err := ReadLenMsg(r, bodyLen)
	if err != nil || len(body) == 0 {
		log.Printf("<<< return and got body message:%v, error:%v", body, err)
		return
	}
	log.Printf("<<< got message bodyLen:%d, body:%v\n", bodyLen, body)

	// Convert message
	resp, err := conv.ConvertMsg(body, bodyLen)
	if err != nil {
		log.Printf("<<< return and conv message:%v, error:%v\n", resp, err)
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
		log.Printf(">>> return and write resp message:%v, error:%v", resp, err)
		return
	}
	log.Printf("--- write response message over!!! ---")
}

// Read message of l length from r
func ReadLenMsg(r io.Reader, l int) ([]byte, error) {
	msg := make([]byte, l)
	buf := make([]byte, l)
	total := 0
	for total < l {
		n, err := r.Read(buf)
		if n > 0 {
			copy(msg[total:], buf[:n])
			total += n
		}

		if err != nil {
			if err == io.EOF {
				log.Println("read EOF:", err)
				break
			}
			log.Println("read header error:", err)
			return nil, err
		}
	}

	return msg, nil
}

// Write message of l length to w
func WriteLenMsg(w io.Writer, m []byte, l int) error {
	// Header
	//binary.BigEndian.PutUint64(h, uint64(l))
	w.Write([]byte(fmt.Sprintf("%08s", strconv.Itoa(l))))
	log.Printf(">>> write header:%v, body:%v", fmt.Sprintf("%08s", strconv.Itoa(l)), m)

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
