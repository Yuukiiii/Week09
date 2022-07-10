package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
)

const (
	packLen      = 4
	headerLen    = 2
	versionLen   = 2
	operationLen = 4
	seqIdLen     = 4
	nonBody      = packLen + headerLen + versionLen + operationLen + seqIdLen
)

func processGOIM(conn net.Conn) {
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	var (
		body      []byte
		err       error
		socketBuf []byte
	)
	reader := bufio.NewReader(conn)
	if socketBuf, err = reader.Peek(nonBody); err != nil {
		fmt.Println("read non body package failed, err: " + err.Error())
		return
	}
	// 网络通讯协议采用的是Big-Endian
	// 前 4 个是 package length
	packageLength := binary.BigEndian.Uint32(socketBuf[0:packLen])
	// 第 4 - 6 是头
	headerLength := binary.BigEndian.Uint16(socketBuf[packLen : packLen+headerLen])
	// 第 6 - 8 是版本号
	version := binary.BigEndian.Uint16(socketBuf[packLen+headerLen : packLen+headerLen+versionLen])
	// 8 - 12 是 operation
	operation := binary.BigEndian.Uint32(socketBuf[packLen+headerLen+versionLen : packLen+headerLen+versionLen+operationLen])
	// 12 - 16 是 seqId
	seqId := binary.BigEndian.Uint32(socketBuf[packLen+headerLen+versionLen+operationLen:])
	bodyLen := packLen - headerLen
	if body, err = reader.Peek(bodyLen); err != nil {
		fmt.Println("read body package failed, err: " + err.Error())
		return
	}
	fmt.Println(fmt.Sprintf("packageLength = %v, headerLength = %v, version = %v, operation = %v, seqId = %v, body = %v", packageLength, headerLength, version, operation, seqId, string(body[0:10])))
	return
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		panic("tcp listener closed, err: " + err.Error())
	}
	for {
		conn, err := listen.Accept()
		fmt.Println("Conn Accepted")
		if err != nil {
			fmt.Println("Conn err: " + err.Error())
			continue
		}
		go processGOIM(conn)
	}
}
