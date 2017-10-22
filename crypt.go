package main

import (
	"crypto/md5"
	//"fmt"
)

func hmac64_md5(data []byte, key []byte) []byte {
	m := make([]byte, 48)
	for i := 0; i < 3; i++ {
		copy(m[i*16:i*16+8], data)
		copy(m[i*16+8:(i+1)*16], key)
	}
	//fmt.Printf("m = %#x\n", m)
	h := md5.New()
	_, _ = h.Write(m)
	r := h.Sum(nil)
	//fmt.Printf("r = %#x\n", r)
	result := make([]byte, 8)
	for i := 0; i < 8; i++ {
		result[i] = r[i] ^ r[i+8]
	}
	return result
}

func hashkey(data []byte) []byte {
	djb_hash := uint32(5381)
	js_hash := uint32(1315423911)

	len := len(data)
	for i := 0; i < len; i++ {
		c := (uint32)(data[i])
		djb_hash += (djb_hash << 5) + c
		js_hash ^= ((js_hash << 5) + c + (js_hash >> 2))
	}

	key := make([]byte, 8)
	key[0] = byte(djb_hash & 0xff)
	key[1] = byte((djb_hash >> 8) & 0xff)
	key[2] = byte((djb_hash >> 16) & 0xff)
	key[3] = byte((djb_hash >> 24) & 0xff)

	key[4] = byte(js_hash & 0xff)
	key[5] = byte((js_hash >> 8) & 0xff)
	key[6] = byte((js_hash >> 16) & 0xff)
	key[7] = byte((js_hash >> 24) & 0xff)
	return key
}
