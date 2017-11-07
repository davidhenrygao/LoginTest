package main

import (
	"encoding/base64"
	"fmt"
	"github.com/davidhenrygao/LoginTest/proto/common"
	"github.com/davidhenrygao/LoginTest/proto/login"
	"github.com/davidhenrygao/LoginTest/proto/player"
	"github.com/golang/protobuf/proto"
	"net"
	"strconv"
	"time"
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

var index int = 1

func doEcho(conn net.Conn, msg string) error {
	fmt.Println("Send echo msg: ", msg)
	c2s_echo := &player.C2SEcho{}
	c2s_echo.Msg = proto.String(msg)
	data := Marshal(c2s_echo)
	if data == nil {
		fmt.Println("Marshal c2s_echo error.")
		return fmt.Errorf("Marshal c2s_echo error.")
	}
	err := writePackage(conn, uint32(common.Cmd_ECHO), data)
	if err != nil {
		fmt.Println("c2s_echo writePackage error.")
		return fmt.Errorf("c2s_echo writePackage error.")
	}
	err, pro := readPackage(conn)
	if err != nil {
		return err
	}
	s2c_echo := &player.S2CEcho{}
	err = Unmarshal(pro.data, s2c_echo)
	if err != nil {
		fmt.Printf("Unmarshal s2c_echo error: %+v\n", err)
		return err
	}
	msgback := s2c_echo.GetMsg()
	fmt.Printf("msgback = %+v\n", msgback)
	return nil
}

func doLogout(conn net.Conn, uid uint64) error {
	fmt.Printf("logout player: id(%d).\n", uid)
	c2s_logout := &player.C2SLogout{}
	c2s_logout.Uid = proto.Uint64(uid)
	data := Marshal(c2s_logout)
	if data == nil {
		fmt.Println("Marshal c2s_logout error.")
		return fmt.Errorf("Marshal c2s_logout error.")
	}
	err := writePackage(conn, uint32(common.Cmd_LOGOUT), data)
	if err != nil {
		fmt.Println("c2s_logout writePackage error.")
		return fmt.Errorf("c2s_logout writePackage error.")
	}
	err, pro := readPackage(conn)
	if err != nil {
		return err
	}
	s2c_logout := &player.S2CLogout{}
	err = Unmarshal(pro.data, s2c_logout)
	if err != nil {
		fmt.Printf("Unmarshal s2c_logout error: %+v\n", err)
		return err
	}
	code := s2c_logout.GetCode()
	fmt.Printf("logout code = %+v\n", code)
	return nil
}

func launch(serveraddr string, secret, openid, subid []byte, echo string, logout bool) error {
	fmt.Printf("Connect to server(%s).\n", serveraddr)
	conn, err := net.Dial("tcp", serveraddr)
	if err != nil {
		fmt.Printf("Connect to gate server error: %s.\n", err)
		return err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println(err)
		}
		index++
	}()

	c2s_launch := &login.C2SLaunch{}
	t := base64.StdEncoding.EncodeToString(openid)
	s := base64.StdEncoding.EncodeToString(subid)
	idx := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(index)))
	etoken := t + "@" + s + ":" + idx
	hetoken := hashkey([]byte(etoken))
	hmac := hmac64_md5(hetoken, secret)
	hmacStr := base64.StdEncoding.EncodeToString(hmac)
	/*
		fmt.Printf("etoken = %+v\n", etoken)
		fmt.Printf("secret = %#x\n", secret)
		fmt.Printf("hashkey = %#x\n", hetoken)
		fmt.Printf("hmac = %#x\n", hmac)
		fmt.Printf("hmacStr = %s\n", hmacStr)
	*/
	signature := etoken + ":" + hmacStr
	fmt.Println("signature: ", signature)
	c2s_launch.Signature = proto.String(signature)
	data := Marshal(c2s_launch)
	if data == nil {
		fmt.Println("Marshal c2s_launch error.")
		return fmt.Errorf("Marshal c2s_launch error.")
	}
	err = writePackage(conn, uint32(common.Cmd_LOGIN_LAUNCH), data)
	if err != nil {
		fmt.Println("c2s_launch writePackage error.")
		return fmt.Errorf("c2s_launch writePackage error.")
	}
	err, pro := readPackage(conn)
	if err != nil {
		return err
	}
	s2c_launch := &login.S2CLaunch{}
	err = Unmarshal(pro.data, s2c_launch)
	if err != nil {
		fmt.Printf("Unmarshal s2c_launch error: %+v\n", err)
		return err
	}
	code := s2c_launch.GetCode()
	if code != 0 {
		fmt.Printf("launch error: code = %+v\n", code)
		return fmt.Errorf("launch error.")
	}
	player := s2c_launch.GetPlayer()
	fmt.Printf("player = %+v\n", player)

	if echo != "" {
		err = doEcho(conn, echo)
		if err != nil {
			return err
		}
	}

	if logout {
		err = doLogout(conn, player.GetId())
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	loginserveradd := "192.168.2.188:10086"
	//loginserveradd := "192.168.0.168:10086"
	//loginserveradd := "192.168.36.64:10086"
	fmt.Printf("Connect to server(%s).", loginserveradd)
	conn, err := net.Dial("tcp", loginserveradd)
	if err != nil {
		fmt.Printf("Connect to login server error: %s.\n", err)
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

	//time.Sleep(10 * time.Second)

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
	openid := "1234567890"
	c2s_login.Platformid = proto.Uint32(1)
	c2s_login.Openid = proto.String(openid)
	c2s_login.Unionid = proto.String("abcdefghijk")
	c2s_login.Nickname = proto.String("david")
	c2s_login.Headimgurl = proto.String("https://www.wx.com/davidhenry.jpg")
	/*
		token := "david"
		platform := "finyin"
		tokenStr := base64.StdEncoding.EncodeToString([]byte(token))
		platformStr := base64.StdEncoding.EncodeToString([]byte(platform))
		tpStr := tokenStr + "@" + platformStr
		err, etp := EncryptDES_ECB([]byte(tpStr), secret)
		if err != nil {
			fmt.Printf("EncryptDES_ECB error = %+v\n", err)
			return
		}
		etpStr := base64.StdEncoding.EncodeToString(etp)
		fmt.Println("base64(DES(secret, token)) string: ", etpStr)
		c2s_login.Token = proto.String(etpStr)
	*/
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
	fmt.Printf("base64(subid): %s.\n", base64.StdEncoding.EncodeToString(subid))
	fmt.Printf("serveraddr: %s.\n", string(serveraddr))

	//launch 1
	err = launch(string(serveraddr), secret, []byte(openid), subid, "Hi! Server", false)
	if err != nil {
		fmt.Printf("launch error = %+v\n", err)
		return
	}
	//launch 2
	err = launch(string(serveraddr), secret, []byte(openid), subid, "Good luck!", true)
	if err != nil {
		fmt.Printf("launch error = %+v\n", err)
		return
	}
	/*
		//launch 3
		err = launch(string(serveraddr), secret, []byte(token), subid, "Good Bye!", true)
		if err != nil {
			fmt.Printf("launch error = %+v\n", err)
			return
		}
	*/

	time.Sleep(10 * time.Second)

	fmt.Println("All test pass!")
}
