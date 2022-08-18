package tis620

import "unicode/utf8"

const OFFSET = 0xD60
const WIDTH = 3 // because thai character in utf-8 use 3 bytes (start at : 0xE01 to 0xE5B)

func ToUTF8(tis620bytes []byte) []byte {
	output := make([]byte, 0)
	buffer := make([]byte, WIDTH)
	for _, c := range tis620bytes {
		if !isTIS620ThaiChar(c) {
			output = append(output, c)
		} else {
			utf8.EncodeRune(buffer, rune(c)+OFFSET)
			output = append(output, buffer...)
		}
	}
	return output
}

func ToTIS620(utf8bytes []rune) []byte {
	output := make([]byte, 0)

	for _, chr := range utf8bytes {
		if isUTFEngChar(chr) {
			output = append(output, byte(chr))
		}

		if isUTFThaiChar(chr) {
			output = append(output, byte(chr-OFFSET))
		}
	}

	return output
}

func CheckTIS620(data []byte) bool {

	for _, chr := range data {
		if !isTIS620Char(chr) {
			return false
		}
	}

	return true
}

func isUTFThaiChar(chr rune) bool {
	return (chr >= 0xE01 && chr <= 0xE3A) || (chr >= 0xE3F && chr <= 0xE5B)
}

func isUTFEngChar(chr rune) bool {
	return (chr >= 0x20) && (chr <= 0x7E)
}

func isTIS620ThaiChar(c byte) bool {
	return (c >= 0xA1 && c <= 0xDA) || (c >= 0xDF && c <= 0xFB)
}

func isTIS620Char(c byte) bool {
	return (c >= 0x20 && c <= 0x7E) || (c >= 0xA1 && c <= 0xDA) || (c >= 0xDF && c <= 0xFB)
}
