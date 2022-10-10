package tea

import (
	"encoding/binary"
)

const (
	op = uint32(0xffffffff)
)

type teaCipher struct {
	k1 uint32
	k2 uint32
	k3 uint32
	k4 uint32
}

func code(v uint64, cipher *teaCipher) []byte {
	n := 16
	delta := uint32(0x9e3779b9)
	y := uint32(v >> 32)
	z := uint32(v)
	s := uint32(0)
	for i := 0; i < n; i++ {
		s += delta
		y += (z + s) ^ (z<<4 + cipher.k1) ^ (z>>5 + cipher.k2)
		z += (y + s) ^ (y<<4 + cipher.k3) ^ (y>>5 + cipher.k4)
	}
	r := make([]byte, 8)
	binary.BigEndian.PutUint32(r[0:], y)
	binary.BigEndian.PutUint32(r[4:], z)
	return r
}

func decode(v uint64, cipher *teaCipher) []byte {
	n := 16
	delta := uint32(0x9E3779B9)
	y := uint32(v >> 32)
	z := uint32(v)
	s := uint32((delta << 4) & op)
	for i := 0; i < n; i++ {
		z -= (y + s) ^ (y<<4 + cipher.k3) ^ (y>>5 + cipher.k4)
		y -= (z + s) ^ (z<<4 + cipher.k1) ^ (z>>5 + cipher.k2)
		s -= delta
	}
	r := make([]byte, 8)
	binary.BigEndian.PutUint32(r[0:], y)
	binary.BigEndian.PutUint32(r[4:], z)
	return r
}

func (cipher *teaCipher) Encrypt(v []byte) []byte {
	length := len(v)
	filln := 10 - (length+1)%8
	encrypted := make([]byte, filln+length+7)
	encrypted[0] = byte(filln-3) | 0xF8
	copy(encrypted[filln:], v)
	var tr, to, o uint64
	for i := 0; i < len(encrypted); i += 8 {
		o = binary.BigEndian.Uint64(encrypted[i:i+8]) ^ tr
		tr = binary.BigEndian.Uint64(code(o, cipher)) ^ to
		to = o
		binary.BigEndian.PutUint64(encrypted[i:i+8], tr)
	}
	return encrypted
}

func (cipher *teaCipher) Decrypt(v []byte) []byte {
	if len(v) < 16 || len(v)%8 != 0 {
		return nil
	}
	preCrypt := binary.BigEndian.Uint64(v[0:8])
	prePlain := binary.BigEndian.Uint64(decode(preCrypt, cipher))
	pos := byte(prePlain>>56)&0x07 + 3
	decrypted := make([]byte, len(v))
	binary.BigEndian.PutUint64(decrypted[0:], prePlain)
	var x uint64
	for i := 8; i < len(decrypted); i += 8 {
		x = binary.BigEndian.Uint64(decode(binary.BigEndian.Uint64(v[i:])^prePlain, cipher)) ^ preCrypt
		prePlain = x ^ preCrypt
		preCrypt = binary.BigEndian.Uint64(v[i : i+8])
		binary.BigEndian.PutUint64(decrypted[i:], x)
	}
	return decrypted[pos : len(decrypted)-7]
}

func NewTeaCipher(k []byte) *teaCipher {
	return &teaCipher{k1: binary.BigEndian.Uint32(k[0:4]), k2: binary.BigEndian.Uint32(k[4:8]), k3: binary.BigEndian.Uint32(k[8:12]), k4: binary.BigEndian.Uint32(k[12:16])}
}
