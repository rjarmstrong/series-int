package series

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

// intervals of 4
func NewInt1664(periods int) Int1664 {
	return Int1664(make([]uint64, periods/4))
}

type Int1664 []uint64

func (sb *Int1664) Periods() int {
	return len(*sb) * 4
}

func (sb *Int1664) Incr(i uint16) {
}

func (sb Int1664) Val(i int) uint16 {
	return uint16(sb[i*2])<<8 + uint16(sb[(i*2)+1])
}

func (sb *Int1664) Add(sb2 *Int1664) error {
	if len(*sb) != len(*sb2) {
		return errors.New("Can't add two unequal series integers")
	}
	for i := 0; i < len(*sb); i += 1 {
		(*sb)[i] += (*sb)[i]

	}
	return nil
}

func (sb *Int1664) AddRange(i int, i2 int, value uint16) {
}

func (sb *Int1664) Set(i int, value uint16) error {
	if value > 65535 {
		return errors.New("Value is greater than a SeriesInt16 can handle")
	}
	// chunk index
	chunkIx := i / 4
	seg := (*sb)[chunkIx]
	// bit offset
	shift := uint8(i % chunkIx)
	// clear
	seg &^= 0x1111111111111111 << shift
	// set
	(*sb)[chunkIx] |= uint64(value << shift)
	return nil
}

func (sb *Int1664) SetRange(start int, end int, value uint16) {
}

func (sb Int1664) String() string {
	b := &bytes.Buffer{}
	for i := 0; i < len(sb)/2; i++ {
		b.WriteString(fmt.Sprintf("%d ", sb.Val(i)))
	}
	return strings.TrimSpace(b.String())
}
