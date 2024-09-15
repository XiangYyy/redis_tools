package redisclient

import (
	"encoding/binary"
	"strings"
)

// 二进制解码
func GodDecode(binaryStr string) (string, error) {
	var decoded uint16
	err := binary.Read(strings.NewReader(binaryStr), binary.BigEndian, &decoded)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
	// return fmt.Sprint(decoded), nil

}
