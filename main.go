package main

import (
	"fmt"
	"github.com/davidhenrygao/LoginTest/proto/common"
	"github.com/davidhenrygao/LoginTest/proto/login"
	//"github.com/davidhenrygao/LoginTest/proto/player"
	"encoding/base64"
	"github.com/golang/protobuf/proto"
	"net"
)

func Marshal(v interface{}) []byte {
	data, err := proto.Marshal(v.(proto.Message))
	if err != nil {
		fmt.Printf("Protobuf Marshal error: %s.\n", err)
		return nil
	}
	return data
}

func Unmarshal(b []byte, v interface{}) error {
	err := proto.Unmarshal(b, v.(proto.Message))
	return err
}

func base64Decode(s string) []byte {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Printf("base64 decode error: %#v.\n", err)
		return nil
	}
	return b
}

func main() {
	fmt.Println("Connect to server(192.168.2.188:10086).")
	conn, err := net.Dial("tcp", "192.168.2.188:10086")
	if err != nil {
		fmt.Printf("Connect to server error: %s.\n", err)
		return
	}

	//challenge
	var pro protocol
	err, pro = readPackage(conn)
	if err != nil {
		return
	}
	if common.Cmd(pro.cmd) != common.Cmd_LOGIN_CHALLENGE {
		fmt.Printf("Unexpected cmd(%d): %#v.\n", pro.cmd, common.Cmd_name[int32(pro.cmd)])
		return
	}
	s2c_challenge := &login.S2CChallenge{}
	err = Unmarshal(pro.data, s2c_challenge)
	if err != nil {
		fmt.Printf("Unmarshal challenge error: %+v\n", err)
		return
	}
	challengeStr := s2c_challenge.GetChallenge()
	challenge := base64Decode(challengeStr)
	if challenge == nil || len(challenge) != 8 {
		fmt.Printf("challenge(%#v) illegal!\n", challenge)
		return
	}
	fmt.Printf("challenge: %#v\n", challenge)

	//exchangekey
	c2s_exchangekey := &login.C2SExchangekey{}
	clientkey := []byte("12345678")
	err, exclientkey := DHExchange(clientkey)
	if err != nil {
		return
	}
	clientkeyStr := base64.StdEncoding.EncodeToString(exclientkey)
	fmt.Println("base64(clientkey) string: ", clientkeyStr)
	c2s_exchangekey.Clientkey = proto.String(clientkeyStr)
	data := Marshal(c2s_exchangekey)
	if data == nil {
		fmt.Println("Marshal c2s_exchangekey error.")
		return
	}
	err = writePackage(conn, uint32(common.Cmd_LOGIN_EXCHANGEKEY), data)
	if err != nil {
		fmt.Println("c2s_exchangekey writePackage error.")
		return
	}
	err, pro = readPackage(conn)
	if err != nil {
		return
	}
	s2c_exchangekey := &login.S2CExchangekey{}
	err = Unmarshal(pro.data, s2c_exchangekey)
	if err != nil {
		fmt.Printf("Unmarshal s2c exchangekey error: %+v\n", err)
		return
	}
	serverkeyStr := s2c_exchangekey.GetServerkey()
	serverkey := base64Decode(serverkeyStr)
	if serverkey == nil || len(serverkey) != 8 {
		fmt.Printf("serverkey(%#v) illegal!\n", serverkey)
		return
	}
	fmt.Printf("serverkey: %#v\n", serverkey)
	err, secret := DHSecret(serverkey, clientkey)
	if err != nil {
		return
	}
	fmt.Printf("secret = %+v\n", secret)

	//handshake
}
