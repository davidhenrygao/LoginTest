package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

const (
	P uint64 = 0xffffffffffffffc5
	G uint64 = 5
)

const N int = int(unsafe.Sizeof(0))

func LocalEndian() binary.ByteOrder {
	x := 1
	p := unsafe.Pointer(&x)
	p2 := (*[N]byte)(p)
	if p2[0] == 1 {
		fmt.Println("LittleEndian.")
		return binary.LittleEndian
	} else {
		fmt.Println("BigEndian.")
		return binary.BigEndian
	}
}

func mul_mod_p(a uint64, b uint64) uint64 {
	var m uint64 = 0
	var t uint64
	for b != 0 {
		if (b & 1) != 0 {
			t = P - a
			if m >= t {
				m -= t
			} else {
				m += a
			}
		}
		if a >= P-a {
			a = a*2 - P
		} else {
			a = a * 2
		}
		b >>= 1
	}
	return m
}

func pow_mod_p(a uint64, b uint64) uint64 {
	if b == 1 {
		return a
	}
	t := pow_mod_p(a, b>>1)
	t = mul_mod_p(t, t)
	if (b % 2) != 0 {
		t = mul_mod_p(t, a)
	}
	return t
}

func powModP(a uint64, b uint64) uint64 {
	if a > P {
		a = a % P
	}
	return pow_mod_p(a, b)
}

func DHExchange(key []byte) (error, []byte) {
	if len(key) != 8 {
		err := fmt.Errorf("Input key length illegal: expected(8), got(%d)!",
			len(key))
		return err, nil
	}
	endian := LocalEndian()
	e := endian.Uint64(key)
	ret := powModP(G, e)
	exkey := make([]byte, 8)
	endian.PutUint64(exkey, ret)
	return nil, exkey
}

func DHSecret(pub []byte, pri []byte) (error, []byte) {
	if len(pub) != 8 || len(pri) != 8 {
		err := fmt.Errorf("Input keys length illegal: expected(8, 8), got(%d, %d)!",
			len(pub), len(pri))
		return err, nil
	}
	endian := LocalEndian()
	a := endian.Uint64(pub)
	b := endian.Uint64(pri)
	ret := powModP(a, b)
	secret := make([]byte, 8)
	endian.PutUint64(secret, ret)
	return nil, secret
}
