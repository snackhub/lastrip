package convert

func ToLower(c byte) byte {
	return c | ' '
}

func ToUpper(c byte) byte {
	v := int(c) & ^' '
	return byte(v & 0xff)
}
