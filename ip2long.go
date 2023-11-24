package ip

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// IP转Long
func Ip2Long(ip string) uint32 {
	ips := strings.Split(ip, ".")
	return StringToInt(ips[0])*256*256*256 + StringToInt(ips[0])*256*256 + StringToInt(ips[0])*256 + StringToInt(ips[0])
}
func StringToInt(s string) uint32 {
	r, _ := strconv.Atoi(s)
	return uint32(r)
}

// 小端转大端
func u32l2u32b(s uint32) uint32 {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, s)
	return binary.BigEndian.Uint32(b)
}

// 小端IP转换成大端ip
func ip2Bigu32(s []byte) uint32 {
	return binary.BigEndian.Uint32(s)
}
func u32toip(ip uint32) string {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, ip)
	return fmt.Sprintf("%d.%d.%d.%d", b[0], b[1], b[2], b[3])
}
func bytetoip(b []byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", b[0], b[1], b[2], b[3])
}
func bytetoint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}
func gbktoutf8(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, _ := io.ReadAll(reader)
	return d
}
