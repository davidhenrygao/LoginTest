package main

import (
	"encoding/base64"
	"fmt"
	"github.com/davidhenrygao/LoginTest/proto/battle"
	"github.com/davidhenrygao/LoginTest/proto/card"
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

func send_heartbeat(conn net.Conn) error {
	fmt.Println("Send heartbeat msg.")
	c2s_heartbeat := &common.C2SHeartbeat{}
	data := Marshal(c2s_heartbeat)
	if data == nil {
		fmt.Println("Marshal c2s_heartbeat error.")
		return fmt.Errorf("Marshal c2s_heartbeat error.")
	}
	err := writePackage(conn, uint32(common.Cmd_HEARTBEAT), data)
	if err != nil {
		fmt.Println("c2s_heartbeat writePackage error.")
		return fmt.Errorf("c2s_heartbeat writePackage error.")
	}
	err, pro := readPackage(conn)
	if err != nil {
		return err
	}
	s2c_heartbeat := &common.S2CHeartbeat{}
	err = Unmarshal(pro.data, s2c_heartbeat)
	if err != nil {
		fmt.Printf("Unmarshal s2c_heartbeat error: %+v\n", err)
		return err
	}
	fmt.Printf("s2c_heartbeat = %+v\n", s2c_heartbeat)
	return nil
}

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

func loadCards(conn net.Conn, idx, pz uint32) error {
	fmt.Printf("load player cards(%d, %d).\n", idx, pz)
	c2s_load_cards := &card.C2SLoadCards{}
	if idx != 1 {
		c2s_load_cards.BeginIndex = proto.Uint32(idx)
	}
	if pz != 0 {
		c2s_load_cards.BeginIndex = proto.Uint32(idx)
	}
	data := Marshal(c2s_load_cards)
	if data == nil {
		fmt.Println("Marshal c2s_load_cards error.")
		return fmt.Errorf("Marshal c2s_load_cards error.")
	}
	err := writePackage(conn, uint32(common.Cmd_LOAD_CARDS), data)
	if err != nil {
		fmt.Println("c2s_load_cards writePackage error.")
		return fmt.Errorf("c2s_load_cards writePackage error.")
	}
	err, pro := readPackage(conn)
	if err != nil {
		return err
	}
	s2c_load_cards := &card.S2CLoadCards{}
	err = Unmarshal(pro.data, s2c_load_cards)
	if err != nil {
		fmt.Printf("Unmarshal s2c_load_cards error: %+v\n", err)
		return err
	}
	fmt.Printf("s2c_load_cards = %+v\n", s2c_load_cards)
	return nil
}

func loadCardDecks(conn net.Conn) error {
	fmt.Printf("load player card decks.\n")
	c2s_load_card_decks := &card.C2SLoadCardDecks{}
	data := Marshal(c2s_load_card_decks)
	if data == nil {
		fmt.Println("Marshal c2s_load_card_decks error.")
		return fmt.Errorf("Marshal c2s_load_card_decks error.")
	}
	err := writePackage(conn, uint32(common.Cmd_LOAD_CARD_DECKS), data)
	if err != nil {
		fmt.Println("c2s_load_card_decks writePackage error.")
		return fmt.Errorf("c2s_load_card_decks writePackage error.")
	}
	err, pro := readPackage(conn)
	if err != nil {
		return err
	}
	s2c_load_card_decks := &card.S2CLoadCardDecks{}
	err = Unmarshal(pro.data, s2c_load_card_decks)
	if err != nil {
		fmt.Printf("Unmarshal s2c_load_card_decks error: %+v\n", err)
		return err
	}
	fmt.Printf("s2c_load_card_decks = %+v\n", s2c_load_card_decks)
	return nil
}

func gmGetCard(conn net.Conn, id, amount uint32) error {
	fmt.Printf("gm player get card.\n")
	c2s_gm_get_card := &card.C2SGmGetCard{}
	c2s_gm_get_card.Id = proto.Uint32(id)
	c2s_gm_get_card.Amount = proto.Uint32(amount)
	data := Marshal(c2s_gm_get_card)
	if data == nil {
		fmt.Println("Marshal c2s_gm_get_card error.")
		return fmt.Errorf("Marshal c2s_gm_get_card error.")
	}
	err := writePackage(conn, uint32(common.Cmd_GM_GET_CARD), data)
	if err != nil {
		fmt.Println("c2s_gm_get_card writePackage error.")
		return fmt.Errorf("c2s_gm_get_card writePackage error.")
	}
	for true {
		err, pro := readPackage(conn)
		if err != nil {
			return err
		}
		if pro.cmd == uint32(common.Cmd_GM_GET_CARD) {
			s2c_gm_get_card := &card.S2CGmGetCard{}
			err = Unmarshal(pro.data, s2c_gm_get_card)
			if err != nil {
				fmt.Printf("Unmarshal s2c_gm_get_card error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_gm_get_card = %+v\n", s2c_gm_get_card)
			break
		} else {
			s2c_update_cards := &card.S2CUpdateCards{}
			err = Unmarshal(pro.data, s2c_update_cards)
			if err != nil {
				fmt.Printf("Unmarshal s2c_update_cards error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_update_cards = %+v\n", s2c_update_cards)
		}
	}
	return nil
}

func gmChangePlyaerProperty(conn net.Conn, ctype uint32, val int32) error {
	fmt.Printf("gm change player property(%d:%d).\n", ctype, val)
	c2s_gm_change_player_property := &player.C2SGmChangePlayerProperty{}
	c2s_gm_change_player_property.Opcode = proto.Uint32(ctype)
	c2s_gm_change_player_property.Args = make([]*common.IarrayElem, 1)
	c2s_gm_change_player_property.Args[0] = &common.IarrayElem{
		Value: proto.Int32(val),
		Pos:   proto.Uint32(1),
	}
	/*
		iarray_elem := &common.IarrayElem{}
		iarray_elem.Value = proto.Int32(val)
		iarray_elem.Pos = proto.Uint32(1)
		argsArray := []*common.IarrayElem{iarray_elem}
		//argsArray := make([]*common.IarrayElem, 1)
		//argsArray[0] = iarray_elem
		c2s_gm_change_player_property.Args = argsArray
	*/
	data := Marshal(c2s_gm_change_player_property)
	if data == nil {
		fmt.Println("Marshal c2s_gm_change_player_property error.")
		return fmt.Errorf("Marshal c2s_gm_change_player_property error.")
	}
	err := writePackage(conn, uint32(common.Cmd_GM_CHANGE_PLAYER_PROPERTY), data)
	if err != nil {
		fmt.Println("c2s_gm_change_player_property writePackage error.")
		return fmt.Errorf("c2s_gm_change_player_property writePackage error.")
	}
	now := time.Now()
	deadline := now.Add(5 * time.Second)
	err = conn.SetReadDeadline(deadline)
	if err != nil {
		return err
	}
	defer func() {
		err = conn.SetReadDeadline(time.Time{})
	}()
	for true {
		err, pro := readPackage(conn)
		if err != nil {
			return err
		}
		if pro.cmd == uint32(common.Cmd_GM_CHANGE_PLAYER_PROPERTY) {
			s2c_gm_change_player_property := &player.S2CGmChangePlayerProperty{}
			err = Unmarshal(pro.data, s2c_gm_change_player_property)
			if err != nil {
				fmt.Printf("Unmarshal s2c_gm_change_player_property error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_gm_change_player_property = %+v\n", s2c_gm_change_player_property)
		} else if pro.cmd == uint32(common.Cmd_UPDATE_CARDS) {
			s2c_update_cards := &card.S2CUpdateCards{}
			err = Unmarshal(pro.data, s2c_update_cards)
			if err != nil {
				fmt.Printf("Unmarshal s2c_update_cards error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_update_cards = %+v\n", s2c_update_cards)
		} else {
			s2c_update_player_property := &player.S2CUpdatePlayerProperty{}
			err = Unmarshal(pro.data, s2c_update_player_property)
			if err != nil {
				fmt.Printf("Unmarshal s2c_update_player_property error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_update_player_property = %+v\n", s2c_update_player_property)
		}
	}
	return nil
}

func cardLevelUp(conn net.Conn, id, uplv uint32) error {
	fmt.Printf("Card level up.\n")
	c2s_up_card_level := &card.C2SUpCardLevel{}
	c2s_up_card_level.Id = proto.Uint32(id)
	c2s_up_card_level.UpLevel = proto.Uint32(uplv)
	data := Marshal(c2s_up_card_level)
	if data == nil {
		fmt.Println("Marshal c2s_up_card_level error.")
		return fmt.Errorf("Marshal c2s_up_card_level error.")
	}
	err := writePackage(conn, uint32(common.Cmd_UP_CARD_LEVEL), data)
	if err != nil {
		fmt.Println("c2s_up_card_level writePackage error.")
		return fmt.Errorf("c2s_up_card_level writePackage error.")
	}
	now := time.Now()
	deadline := now.Add(5 * time.Second)
	err = conn.SetReadDeadline(deadline)
	if err != nil {
		return err
	}
	defer func() {
		err = conn.SetReadDeadline(time.Time{})
	}()
	for true {
		err, pro := readPackage(conn)
		if err != nil {
			return err
		}
		if pro.cmd == uint32(common.Cmd_UP_CARD_LEVEL) {
			s2c_up_card_level := &card.S2CUpCardLevel{}
			err = Unmarshal(pro.data, s2c_up_card_level)
			if err != nil {
				fmt.Printf("Unmarshal s2c_up_card_level error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_up_card_level = %+v\n", s2c_up_card_level)
		} else if pro.cmd == uint32(common.Cmd_UPDATE_CARDS) {
			s2c_update_cards := &card.S2CUpdateCards{}
			err = Unmarshal(pro.data, s2c_update_cards)
			if err != nil {
				fmt.Printf("Unmarshal s2c_update_cards error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_update_cards = %+v\n", s2c_update_cards)
		} else {
			s2c_update_player_property := &player.S2CUpdatePlayerProperty{}
			err = Unmarshal(pro.data, s2c_update_player_property)
			if err != nil {
				fmt.Printf("Unmarshal s2c_update_player_property error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_update_player_property = %+v\n", s2c_update_player_property)
		}
	}
	return nil
}

func checkCard(conn net.Conn, id uint32) error {
	fmt.Printf("Check Card.\n")
	c2s_check_card := &card.C2SCheckCard{}
	c2s_check_card.Id = proto.Uint32(id)
	data := Marshal(c2s_check_card)
	if data == nil {
		fmt.Println("Marshal c2s_check_card error.")
		return fmt.Errorf("Marshal c2s_check_card error.")
	}
	err := writePackage(conn, uint32(common.Cmd_CHECK_CARD), data)
	if err != nil {
		fmt.Println("c2s_check_card writePackage error.")
		return fmt.Errorf("c2s_check_card writePackage error.")
	}
	for true {
		err, pro := readPackage(conn)
		if err != nil {
			return err
		}
		if pro.cmd == uint32(common.Cmd_CHECK_CARD) {
			s2c_check_card := &card.S2CCheckCard{}
			err = Unmarshal(pro.data, s2c_check_card)
			if err != nil {
				fmt.Printf("Unmarshal s2c_check_card error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_check_card = %+v\n", s2c_check_card)
			break
		} else {
			s2c_update_cards := &card.S2CUpdateCards{}
			err = Unmarshal(pro.data, s2c_update_cards)
			if err != nil {
				fmt.Printf("Unmarshal s2c_update_cards error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_update_cards = %+v\n", s2c_update_cards)
		}
	}
	return nil
}

func changeDeck(conn net.Conn, index uint32) error {
	fmt.Printf("Change Deck.\n")
	c2s_change_deck := &card.C2SChangeDeck{}
	c2s_change_deck.Index = proto.Uint32(index)
	data := Marshal(c2s_change_deck)
	if data == nil {
		fmt.Println("Marshal c2s_change_deck error.")
		return fmt.Errorf("Marshal c2s_change_deck error.")
	}
	err := writePackage(conn, uint32(common.Cmd_CHANGE_DECK), data)
	if err != nil {
		fmt.Println("c2s_change_deck writePackage error.")
		return fmt.Errorf("c2s_change_deck writePackage error.")
	}
	for true {
		err, pro := readPackage(conn)
		if err != nil {
			return err
		}
		if pro.cmd == uint32(common.Cmd_CHANGE_DECK) {
			s2c_change_deck := &card.S2CChangeDeck{}
			err = Unmarshal(pro.data, s2c_change_deck)
			if err != nil {
				fmt.Printf("Unmarshal s2c_change_deck error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_change_deck = %+v\n", s2c_change_deck)
			break
		}
	}
	return nil
}

func changeCardDeck(conn net.Conn, index, id, pos uint32) error {
	fmt.Printf("Change card Deck.\n")
	c2s_change_card_deck := &card.C2SChangeCardDeck{}
	change_info := &card.ChangeInfo{}
	change_info.CardDeckIndex = proto.Uint32(index)
	change_info.Id = proto.Uint32(id)
	change_info.Pos = proto.Uint32(pos)
	c2s_change_card_deck.Change = change_info
	data := Marshal(c2s_change_card_deck)
	if data == nil {
		fmt.Println("Marshal c2s_change_card_deck error.")
		return fmt.Errorf("Marshal c2s_change_card_deck error.")
	}
	err := writePackage(conn, uint32(common.Cmd_CHANGE_CARD_DECK), data)
	if err != nil {
		fmt.Println("c2s_change_card_deck writePackage error.")
		return fmt.Errorf("c2s_change_card_deck writePackage error.")
	}
	for true {
		err, pro := readPackage(conn)
		if err != nil {
			return err
		}
		if pro.cmd == uint32(common.Cmd_CHANGE_CARD_DECK) {
			s2c_change_card_deck := &card.S2CChangeCardDeck{}
			err = Unmarshal(pro.data, s2c_change_card_deck)
			if err != nil {
				fmt.Printf("Unmarshal s2c_change_card_deck error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_change_card_deck = %+v\n", s2c_change_card_deck)
			break
		}
	}
	return nil
}

func doBattle(battleServerAddr string, battleId uint32, playerId uint64) error {
	fmt.Printf("Connect to battle server(%s).\n", battleServerAddr)
	conn, err := net.Dial("tcp", battleServerAddr)
	if err != nil {
		fmt.Printf("Connect to battle server error: %s.\n", err)
		return err
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var counter uint32 = 1
	c2s_battle_ready := &battle.C2SBattleReady{}
	c2s_battle_ready.BattleId = proto.Uint32(battleId)
	c2s_battle_ready.PlayerId = proto.Uint64(playerId)
	data := Marshal(c2s_battle_ready)
	if data == nil {
		fmt.Println("Marshal c2s_battle_ready error.")
		return fmt.Errorf("Marshal c2s_battle_ready error.")
	}
	err = writePackage(conn, uint32(common.Cmd_BATTLE_READY), data)
	if err != nil {
		fmt.Println("c2s_battle_ready writePackage error.")
		return fmt.Errorf("c2s_battle_ready writePackage error.")
	}

	for true {
		now := time.Now()
		deadline := now.Add(10 * time.Second)
		err = conn.SetReadDeadline(deadline)
		err, pro := readPackage(conn)
		if err != nil {
			return err
		}
		err = conn.SetReadDeadline(time.Time{})
		if pro.cmd == uint32(common.Cmd_BATTLE_READY) {
			s2c_battle_ready := &battle.S2CBattleReady{}
			err = Unmarshal(pro.data, s2c_battle_ready)
			if err != nil {
				fmt.Printf("Unmarshal s2c_battle_ready error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_battle_ready = %+v\n", s2c_battle_ready)
		} else if pro.cmd == uint32(common.Cmd_BATTLE_START) {
			s2c_battle_start := &battle.S2CBattleStart{}
			err = Unmarshal(pro.data, s2c_battle_start)
			if err != nil {
				fmt.Printf("Unmarshal s2c_battle_start error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_battle_start = %+v\n", s2c_battle_start)
		} else if pro.cmd == uint32(common.Cmd_BATTLE_FRAME_UPDATE) {
			s2c_battle_frame_update := &battle.S2CBattleFrameUpdate{}
			err = Unmarshal(pro.data, s2c_battle_frame_update)
			if err != nil {
				fmt.Printf("Unmarshal s2c_battle_frame_update error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_battle_frame_update = %+v\n", s2c_battle_frame_update)
			if counter >= 101 {
				c2s_battle_end := &battle.C2SBattleEnd{}
				c2s_battle_end.Result = proto.Uint32(0)
				data = Marshal(c2s_battle_end)
				if data == nil {
					fmt.Println("Marshal c2s_battle_end error.")
					return fmt.Errorf("Marshal c2s_battle_end error.")
				}
				err = writePackage(conn, uint32(common.Cmd_BATTLE_END), data)
				if err != nil {
					fmt.Println("c2s_battle_end writePackage error.")
					return fmt.Errorf("c2s_battle_end writePackage error.")
				}
			}
			if counter%4 == 0 {
				c2s_battle_action := &battle.C2SBattleAction{}
				c2s_battle_action.ClassId = proto.Uint32(counter)
				c2s_battle_action.Action = proto.String("fuck you!")
				data = Marshal(c2s_battle_action)
				if data == nil {
					fmt.Println("Marshal c2s_battle_action error.")
					return fmt.Errorf("Marshal c2s_battle_action error.")
				}
				err = writePackage(conn, uint32(common.Cmd_BATTLE_ACTION), data)
				if err != nil {
					fmt.Println("c2s_battle_action writePackage error.")
					return fmt.Errorf("c2s_battle_action writePackage error.")
				}
			}
			counter++
		} else if pro.cmd == uint32(common.Cmd_BATTLE_ACTION) {
			s2c_battle_action := &battle.S2CBattleAction{}
			err = Unmarshal(pro.data, s2c_battle_action)
			if err != nil {
				fmt.Printf("Unmarshal s2c_battle_action error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_battle_action = %+v\n", s2c_battle_action)
		} else if pro.cmd == uint32(common.Cmd_BATTLE_END) {
			s2c_battle_end := &battle.S2CBattleEnd{}
			err = Unmarshal(pro.data, s2c_battle_end)
			if err != nil {
				fmt.Printf("Unmarshal s2c_battle_end error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_battle_end = %+v\n", s2c_battle_end)
			break
		} else {
			fmt.Println("Unknown command = %+v.\n", pro.cmd)
			break
		}
	}
	return nil
}

func test_match_and_battle(conn net.Conn, playerId uint64) error {
	fmt.Printf("Test match and battle.\n")
	c2s_match_start := &battle.C2SMatchStart{}
	c2s_match_start.Type = proto.Uint32(1)
	data := Marshal(c2s_match_start)
	if data == nil {
		fmt.Println("Marshal c2s_match_start error.")
		return fmt.Errorf("Marshal c2s_match_start error.")
	}
	err := writePackage(conn, uint32(common.Cmd_MATCH_START), data)
	if err != nil {
		fmt.Println("c2s_up_card_level writePackage error.")
		return fmt.Errorf("c2s_up_card_level writePackage error.")
	}
	now := time.Now()
	deadline := now.Add(10 * time.Second)
	err = conn.SetReadDeadline(deadline)
	if err != nil {
		return err
	}
	defer func() {
		err = conn.SetReadDeadline(time.Time{})
	}()
	for true {
		now = time.Now()
		deadline = now.Add(10 * time.Second)
		err = conn.SetReadDeadline(deadline)
		err, pro := readPackage(conn)
		if err != nil {
			return err
		}
		err = conn.SetReadDeadline(time.Time{})
		if pro.cmd == uint32(common.Cmd_MATCH_START) {
			s2c_match_start := &battle.S2CMatchStart{}
			err = Unmarshal(pro.data, s2c_match_start)
			if err != nil {
				fmt.Printf("Unmarshal s2c_match_start error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_match_start = %+v\n", s2c_match_start)
		} else if pro.cmd == uint32(common.Cmd_MATCH_UPDATE) {
			s2c_match_update := &battle.S2CMatchUpdate{}
			err = Unmarshal(pro.data, s2c_match_update)
			if err != nil {
				fmt.Printf("Unmarshal s2c_match_update error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_match_update = %+v\n", s2c_match_update)
		} else if pro.cmd == uint32(common.Cmd_BATTLE_INIT) {
			s2c_battle_init := &battle.S2CBattleInit{}
			err = Unmarshal(pro.data, s2c_battle_init)
			if err != nil {
				fmt.Printf("Unmarshal s2c_battle_init error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_battle_init = %+v\n", s2c_battle_init)
			battleId := s2c_battle_init.GetBattleId()
			battleServerAddr := s2c_battle_init.GetServerAddr()
			err = doBattle(battleServerAddr, battleId, playerId)
			if err != nil {
				return err
			}
			break
		} else {
			fmt.Println("Unknown command = %+v.\n", pro.cmd)
			break
		}
	}
	return nil
}

func test_match_and_cancel(conn net.Conn) error {
	fmt.Printf("Test match and cancel.\n")
	c2s_match_start := &battle.C2SMatchStart{}
	c2s_match_start.Type = proto.Uint32(1)
	data := Marshal(c2s_match_start)
	if data == nil {
		fmt.Println("Marshal c2s_match_start error.")
		return fmt.Errorf("Marshal c2s_match_start error.")
	}
	err := writePackage(conn, uint32(common.Cmd_MATCH_START), data)
	if err != nil {
		fmt.Println("c2s_up_card_level writePackage error.")
		return fmt.Errorf("c2s_up_card_level writePackage error.")
	}
	now := time.Now()
	deadline := now.Add(10 * time.Second)
	err = conn.SetReadDeadline(deadline)
	if err != nil {
		return err
	}
	defer func() {
		err = conn.SetReadDeadline(time.Time{})
	}()
	for true {
		now = time.Now()
		deadline = now.Add(10 * time.Second)
		err = conn.SetReadDeadline(deadline)
		err, pro := readPackage(conn)
		if err != nil {
			return err
		}
		err = conn.SetReadDeadline(time.Time{})
		if pro.cmd == uint32(common.Cmd_MATCH_START) {
			s2c_match_start := &battle.S2CMatchStart{}
			err = Unmarshal(pro.data, s2c_match_start)
			if err != nil {
				fmt.Printf("Unmarshal s2c_match_start error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_match_start = %+v\n", s2c_match_start)
		} else if pro.cmd == uint32(common.Cmd_MATCH_UPDATE) {
			s2c_match_update := &battle.S2CMatchUpdate{}
			err = Unmarshal(pro.data, s2c_match_update)
			if err != nil {
				fmt.Printf("Unmarshal s2c_match_update error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_match_update = %+v\n", s2c_match_update)
			c2s_battle_cancel := &battle.C2SMatchCancel{}
			data = Marshal(c2s_battle_cancel)
			if data == nil {
				fmt.Println("Marshal c2s_battle_cancel error.")
				return fmt.Errorf("Marshal c2s_battle_cancel error.")
			}
			err = writePackage(conn, uint32(common.Cmd_MATCH_CANCEL), data)
			if err != nil {
				fmt.Println("c2s_battle_cancel writePackage error.")
				return fmt.Errorf("c2s_battle_cancel writePackage error.")
			}
		} else if pro.cmd == uint32(common.Cmd_MATCH_CANCEL) {
			s2c_match_cancel := &battle.S2CMatchCancel{}
			err = Unmarshal(pro.data, s2c_match_cancel)
			if err != nil {
				fmt.Printf("Unmarshal s2c_match_cancel error: %+v\n", err)
				return err
			}
			fmt.Printf("s2c_match_cancel = %+v\n", s2c_match_cancel)
			break
		} else {
			fmt.Println("Unknown command = %+v.\n", pro.cmd)
			break
		}
	}
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

	err = send_heartbeat(conn)
	if err != nil {
		fmt.Printf("launch send_heartbeat err = %+v\n", err)
		return err
	}

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
		/*
			for i := 0; i < 10; i++ {
				fmt.Printf("send echo request(%d).\n", i)
				err = doEcho(conn, echo)
				if err != nil {
					return err
				}
				time.Sleep(1 * time.Second)
			}
		*/
	}

	err = loadCards(conn, 1, 0)
	if err != nil {
		return err
	}
	/*
		err = loadCards(conn, 3, 0)
		if err != nil {
			return err
		}
		err = loadCards(conn, 2, 100)
		if err != nil {
			return err
		}
		err = loadCards(conn, 100, 100)
		if err != nil {
			return err
		}
	*/

	err = loadCardDecks(conn)
	if err != nil {
		return err
	}

	err = send_heartbeat(conn)
	if err != nil {
		fmt.Printf("launch send_heartbeat err = %+v\n", err)
		return err
	}

	err = gmGetCard(conn, 1004, 100)
	if err != nil {
		return err
	}

	err = gmGetCard(conn, 4005, 100000)
	if err != nil {
		return err
	}

	err = gmChangePlyaerProperty(conn, 3, 1000000)
	if err != nil {
		fmt.Printf("gmChangePlyaerProperty err = %+v\n", err)
	}

	err = cardLevelUp(conn, 4005, 2)
	if err != nil {
		fmt.Printf("cardLevelUp err = %+v\n", err)
	}

	err = cardLevelUp(conn, 1001, 5)
	if err != nil {
		fmt.Printf("cardLevelUp err = %+v\n", err)
	}

	err = cardLevelUp(conn, 1001, 14)
	if err != nil {
		fmt.Printf("cardLevelUp err = %+v\n", err)
	}

	err = checkCard(conn, 1001)
	if err != nil {
		return err
	}

	err = checkCard(conn, 4005)
	if err != nil {
		return err
	}

	err = changeDeck(conn, 2)
	if err != nil {
		return err
	}

	err = changeCardDeck(conn, 0, 1001, 5)
	if err != nil {
		return err
	}

	err = changeCardDeck(conn, 0, 4003, 5)
	if err != nil {
		return err
	}

	err = changeCardDeck(conn, 0, 4005, 5)
	if err != nil {
		return err
	}

	/*
		err = gmChangePlyaerProperty(conn, 1, 1)
		if err != nil {
			fmt.Printf("gmChangePlyaerProperty err = %+v\n", err)
		}

		err = gmChangePlyaerProperty(conn, 2, 1000)
		if err != nil {
			fmt.Printf("gmChangePlyaerProperty err = %+v\n", err)
		}
	*/

	//err = test_match_and_cancel(conn)
	err = test_match_and_battle(conn, player.GetId())
	if err != nil {
		return err
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

	err = send_heartbeat(conn)
	if err != nil {
		fmt.Printf("send_heartbeat err = %+v\n", err)
		return
	}

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

	err = send_heartbeat(conn)
	if err != nil {
		fmt.Printf("send_heartbeat err = %+v\n", err)
		return
	}

	err = send_heartbeat(conn)
	if err != nil {
		fmt.Printf("send_heartbeat err = %+v\n", err)
		return
	}

	//login
	c2s_login := &login.C2SLogin{}
	openid := "1234567890"
	//openid := "abcde"
	c2s_login.Platformid = proto.Uint32(1)
	c2s_login.Openid = proto.String(openid)
	c2s_login.Unionid = proto.String("abcde")
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
