package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const maxMsgLen = (1 << 16) - 20

type protocol struct {
	session  uint16
	cmd      uint32
	len      uint16
	data     []byte
	checksum uint16
}

func (p *protocol) Serialize() (error, []byte) {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, p.session)
	if err != nil {
		return err, nil
	}
	err = binary.Write(&buf, binary.BigEndian, p.cmd)
	if err != nil {
		return err, nil
	}
	err = binary.Write(&buf, binary.BigEndian, p.len)
	if err != nil {
		return err, nil
	}
	err = binary.Write(&buf, binary.BigEndian, p.data)
	if err != nil {
		return err, nil
	}
	err = binary.Write(&buf, binary.BigEndian, p.checksum)
	if err != nil {
		return err, nil
	}
	return nil, buf.Bytes()
}

func (p *protocol) Unserialize(b []byte) error {
	buf := bytes.NewReader(b)
	err := binary.Read(buf, binary.BigEndian, p.session)
	if err != nil {
		return err
	}
	err = binary.Read(buf, binary.BigEndian, p.cmd)
	if err != nil {
		return err
	}
	err = binary.Read(buf, binary.BigEndian, p.len)
	if err != nil {
		return err
	}
	p.data = make([]byte, p.len)
	err = binary.Read(buf, binary.BigEndian, p.data)
	if err != nil {
		return err
	}
	err = binary.Read(buf, binary.BigEndian, p.checksum)
	if err != nil {
		return err
	}
	return nil
}

func readPackage(conn net.Conn) (error, protocol) {
	len := make([]byte, 2)
	_, err := io.ReadFull(conn, len)
	if err != nil {
		fmt.Printf("ReadPackage read msg len error: %s.\n", err)
		return err, protocol{}
	}
	ul := binary.BigEndian.Uint16(len)
	pkg := make([]byte, ul)
	_, err = io.ReadFull(conn, pkg)
	if err != nil {
		fmt.Printf("ReadPackage read msg pkg error: %s.\n", err)
		return err, protocol{}
	}
	p := new(protocol)
	err = p.Unserialize(pkg)
	if err != nil {
		fmt.Printf("ReadPackage unserialize msg error: %s.\n", err)
		return err, protocol{}
	}
	return err, *p
}

var sessionCnt uint16 = 1

func writePackage(conn net.Conn, cmd uint32, b []byte) error {
	l := len(b)
	if l >= maxMsgLen {
		var err = fmt.Errorf("length(%d) exceeds %d", l, maxMsgLen)
		fmt.Printf("WritePackage length error: %s.\n", err)
		return err
	}
	p := &protocol{
		session:  sessionCnt,
		cmd:      cmd,
		len:      uint16(l),
		data:     b,
		checksum: 0,
	}
	err, pkg := p.Serialize()
	if err != nil {
		fmt.Printf("writePackage serialize error: %#v.\n", err)
		return err
	}
	ul := len(pkg)
	wb := make([]byte, 2+ul)
	binary.BigEndian.PutUint16(wb[:2], uint16(ul))
	copy(wb[2:], pkg)
	_, err = conn.Write(wb)
	if err != nil {
		fmt.Printf("WritePackage conn write error: %s.\n", err)
		return err
	}
	sessionCnt += 1
	if sessionCnt == 0 {
		sessionCnt = 1
	}
	return nil
}
