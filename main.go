package main

import (
	"fmt"
	"github.com/davidhenrygao/LoginTest/proto/common"
	"github.com/davidhenrygao/LoginTest/proto/login"
	//"github.com/davidhenrygao/LoginTest/proto/player"
	"crypto/des"
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
	//fmt.Println("Connect to server(192.168.2.188:10086).")
	//conn, err := net.Dial("tcp", "192.168.2.188:10086")
	fmt.Println("Connect to server(192.168.0.168:10086).")
	conn, err := net.Dial("tcp", "192.168.0.168:10086")
	if err != nil {
		fmt.Printf("Connect to server error: %s.\n", err)
		return
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

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
		fmt.Printf("challenge(%#x) illegal!\n", challenge)
		return
	}
	fmt.Printf("challenge: %#x\n", challenge)

	//exchangekey
	c2s_exchangekey := &login.C2SExchangekey{}
	clientkey := []byte("12345678")
	err, exclientkey := DHExchange(clientkey)
	if err != nil {
		return
	}
	clientkeyStr := base64.StdEncoding.EncodeToString(exclientkey)
	fmt.Println("base64(dh-exchange(clientkey)) string: ", clientkeyStr)
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
		fmt.Printf("serverkey(%#x) illegal!\n", serverkey)
		return
	}
	fmt.Printf("serverkey: %#x\n", serverkey)
	err, secret := DHSecret(serverkey, clientkey)
	if err != nil {
		return
	}
	fmt.Printf("secret = %#x\n", secret)

	//handshake
	c2s_handshake := &login.C2SHandshake{}
	encodeChallengeStr := base64.StdEncoding.EncodeToString(hmac64_md5(challenge, secret))
	fmt.Println("base64(hmac64_md5(challenge, secret)) string: ", encodeChallengeStr)
	c2s_handshake.EncodeChallenge = proto.String(encodeChallengeStr)
	data = Marshal(c2s_handshake)
	if data == nil {
		fmt.Println("Marshal c2s_handshake error.")
		return
	}
	err = writePackage(conn, uint32(common.Cmd_LOGIN_HANDSHAKE), data)
	if err != nil {
		fmt.Println("c2s_handshake writePackage error.")
		return
	}
	err, pro = readPackage(conn)
	if err != nil {
		return
	}
	s2c_handshake := &login.S2CHandshake{}
	err = Unmarshal(pro.data, s2c_handshake)
	if err != nil {
		fmt.Printf("Unmarshal s2c handshake error: %+v\n", err)
		return
	}
	code := s2c_handshake.GetCode()
	if code != 0 {
		fmt.Printf("handshake error: code = %+v\n", code)
		return
	}

	//login
	c2s_login := &login.C2SLogin{}
	token := "david"
	platform := "finyin"
	tokenStr := base64.StdEncoding.EncodeToString([]byte(token))
	platformStr := base64.StdEncoding.EncodeToString([]byte(platform))
	tpStr := tokenStr + "@" + platformStr
	codec, err := des.NewCipher(secret)
	if err != nil {
		fmt.Printf("DES new cipher err = %+v\n", err)
		return
	}
	etp := make([]byte, len(tpStr))
	codec.Encrypt(etp, []byte(tpStr))
	etpStr := base64.StdEncoding.EncodeToString(etp)
	fmt.Println("base64(DES(secret, token)) string: ", etpStr)
	c2s_login.Token = proto.String(etpStr)
	data = Marshal(c2s_login)
	if data == nil {
		fmt.Println("Marshal c2s_login error.")
		return
	}
	err = writePackage(conn, uint32(common.Cmd_LOGIN), data)
	if err != nil {
		fmt.Println("c2s_login writePackage error.")
		return
	}
	err, pro = readPackage(conn)
	if err != nil {
		return
	}
	s2c_login := &login.S2CLogin{}
	err = Unmarshal(pro.data, s2c_login)
	if err != nil {
		fmt.Printf("Unmarshal s2c_login error: %+v\n", err)
		return
	}
	code = s2c_login.GetCode()
	if code != 0 {
		fmt.Printf("handshake error: code = %+v\n", code)
		return
	}
	fmt.Println("Login success!")
	serverinfo := s2c_login.GetInfo()
	subid, err := base64.StdEncoding.DecodeString(serverinfo.GetSubid())
	if err != nil {
		fmt.Printf("login get subid base64 decode error: %#v.\n", err)
		return
	}
	serveraddr, err := base64.StdEncoding.DecodeString(serverinfo.GetServerAddr())
	if err != nil {
		fmt.Printf("login get serveraddr base64 decode error: %#v.\n", err)
		return
	}
	fmt.Printf("subid: %s.\n", string(subid))
	fmt.Printf("serveraddr: %s.\n", string(serveraddr))

	//index := 1
	//launch 1
	//index += 1
	//launch 2
	//logout
}
