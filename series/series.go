package series

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

func NewInt16(periods int) Int16 {
	return Int16(make([]byte, periods*2))
}

type Int16 []byte

func (sb *Int16) Incr(i int) {
	if (*sb)[(i*2)+1] == 255 {
		(*sb)[(i*2)+1] = 0
		(*sb)[i*2] += 1
		return
	}
	(*sb)[(i*2)+1] += 1
}

func (sb Int16) Val(i int) uint16 {
	return uint16(sb[i*2])<<8 + uint16(sb[(i*2)+1])
}

func (sb *Int16) Add(sb2 *Int16) error {
	if len(*sb) != len(*sb2) {
		return errors.New("Can't add two unequal series integers")
	}
	for i := 0; i < len(*sb); i += 2 {
		if (*sb2)[i] > 0 {
			(*sb)[i] += (*sb2)[i]
		}
		if (*sb2)[i+1] > 0 {
			unit := uint16((*sb)[i+1]) + uint16((*sb2)[i+1])
			if unit < 256 {
				(*sb)[i+1] = uint8(unit)
				continue
			}
			(*sb)[i+1] = uint8(unit - 256)
			(*sb)[i]++
		}
	}
	return nil
}

func (sb *Int16) Set(i int, value uint16) error {
	if value > 65535 {
		return errors.New("Value is greater than a SeriesInt16 can handle")
	}
	if value < 256 {
		(*sb)[(i*2)+1] = uint8(value)
		return nil
	}
	(*sb)[(i*2)+1] = uint8(value % 256)
	(*sb)[i*2] = uint8(value >> 8)
	return nil
}

func (sb *Int16) SetRange(start int, end int, value uint16) {
	for i := start; i <= end; i++ {
		sb.Set(i, value)
	}
}

func (sb Int16) String() string {
	b := &bytes.Buffer{}
	for i := 0; i < len(sb)/2; i++ {
		b.WriteString(fmt.Sprintf("%d ", sb.Val(i)))
	}
	return strings.TrimSpace(b.String())
}
