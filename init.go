package ip

import (
	_ "embed"
	"encoding/binary"
	"fmt"
)

//go:embed qqwry.dat
var _ipv4Data []byte
var (
	startIP int32
	endIP   int32
	totalIP int32
)

func Init() {
	startIP = int32(binary.LittleEndian.Uint32(_ipv4Data[:4]))
	endIP = int32(binary.LittleEndian.Uint32(_ipv4Data[4:8]))
	totalIP = (endIP - startIP) / 7
	fmt.Printf("[%d ]  %d -> %d\n", totalIP, startIP, endIP)
}
func FindIP(ip string) {
	iplong := StringToInt(ip)
	iplong = u32l2u32b(iplong)
	s := int32(0)
	e := totalIP
	i := int32(0)
	ffip := int32(0)
	for s <= e {
		i = (e + s) / 2
		fip := readIP(i)
		bip := ip2Bigu32(fip[:4])
		if iplong < bip {
			fmt.Printf("%s e %d : %d %v\n", ip, e, i, fip)
			e = i - 1
		} else {
			eip := readIP(i + 7)
			eeip := ip2Bigu32(eip[:4])
			if iplong > eeip {
				s = i + 1
			} else {
				ffip = startIP + i*7
				break
			}
		}
	}
	dd := readIP(ffip)
	bips := bytetoip(dd[:4])

	dd2 := readIP(ffip + 7)
	eeips := bytetoip(dd2[:4])
	fmt.Printf(" %s -> %s\n", bips, eeips)
	fmt.Println(dd2, dd2[4])
	switch dd2[4] {
	case 1:
		fmt.Printf("标志字节为1，表示国家和区域信息都被同时重定向\n")
	case 2:
		fmt.Printf("标志字节为2，表示国家信息被重定向\n")
	default:
		cname, ii := readStr(ffip + 7 + 5)
		fmt.Printf("否则，表示国家信息没有被重定向 %s %d\n", cname, ii)
	}
}

// 读取IP信息
func readIP(s int32) []byte {
	return _ipv4Data[s : s+7]
}

func readStr(i int32) (string, int32) {
	ret := []byte("")
	for _ipv4Data[i] != 0 {
		ret = append(ret, _ipv4Data[i])
		i++
	}
	return string(gbktoutf8(ret)), i
}
