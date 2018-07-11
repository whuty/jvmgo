package classfile

import (
	"fmt"
	"unicode/utf16"
)

// ConstantUtf8Info struct
type ConstantUtf8Info struct {
	str string
}

func (mConstantUtf8Info *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	//mConstantUtf8Info.str = string(bytes)
	mConstantUtf8Info.str = decodeMUTF8(bytes)
}

// Str getter
func (mConstantUtf8Info *ConstantUtf8Info) Str() string {
	return mConstantUtf8Info.str
}

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
func decodeMUTF8(bytearr []byte) string {
	utflen := len(bytearr)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararrCount := 0

	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararrCount] = c
		chararrCount++
	}

	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararrCount] = c
			chararrCount++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararrCount] = c&0x1F<<6 | char2&0x3F
			chararrCount++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararrCount] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararrCount++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararrCount]
	runes := utf16.Decode(chararr)
	return string(runes)
}
