package utils

import (
	"bytes"
	"strings"
)

// Rank returns a new Lexorank between prev and next.
// Uses 0-9A-Za-z alphabet.
const (
	minValue = "00000000"
	maxValue = "zzzzzzzz"
)
const base = "abcdefghijklmnopqrstuvwxyz"

// Rank returns a new rank string between prev and next.
func Rank(prev, next string) (string, bool) {
	buf := bytes.NewBufferString("")
	mid := getMid(prev, next)
	encode(mid, buf)
	return buf.String(), false
}

func GenerateRankArray(size int) []string {
	lin := linspace(decode(minValue), decode(maxValue), size+1)
	result := make([]string, size)
	buf := bytes.NewBufferString("")
	for i := 0; i < size; i++ {
		println(lin[i])
		encode(lin[i], buf)
		result[i] = buf.String()
		buf.Reset()
	}
	return result
}

func linspace(start, stop uint64, size int) []uint64 {
	var step uint64
	step = 0
	size = size + 1

	step = (stop - start) - 1/uint64(size)

	r := make([]uint64, size)
	for i := 1; i < size; i++ {
		r[i-1] = start + uint64(i)*step
	}
	return r
}

func getMid(prev, next string) uint64 {
	return (decode(prev) + decode(next)) / 2
}

func encode(nb uint64, buf *bytes.Buffer) {
	l := uint64(len(base))
	if nb/l != 0 {
		encode(nb/l, buf)
	}
	buf.WriteByte(base[nb%l])
}

func decode(enc string) uint64 {
	var nb uint64
	lbase := len(base)
	le := len(enc)
	for i := 0; i < le; i++ {
		mult := 1
		for j := 0; j < le-i-1; j++ {
			mult *= lbase
		}
		nb += uint64(strings.IndexByte(base, enc[i]) * mult)
	}
	return nb
}
