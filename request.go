package goWHOIS

import (
	"net"
)

func (req *Req) whoisRequest(host string, modifier string) (string, error) {
	conn, err := net.Dial("tcp", host+":43")
	if err != nil {
		return "", err
	}
	conn.Write([]byte(modifier + req.Object + "\r\n"))
	buf := make([]byte, 1024)
	res := []byte{}
	for {
		numbytes, err := conn.Read(buf)
		sbuf := buf[0:numbytes]
		res = append(res, sbuf...)
		if err != nil {
			break
		}
	}
	conn.Close()
	return string(res), nil
}
