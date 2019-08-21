package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"net/url"
)

func IsPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)

}

func PrivateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err

	}
	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}
		ip := ipnet.IP.To4()
		if IsPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, errors.New("no private ip address")
}

func IPtoI(ip string) uint32 {
	ips := net.ParseIP(ip)

	if len(ips) == 16 {
		return binary.BigEndian.Uint32(ips[12:16])
	} else if len(ips) == 4 {
		return binary.BigEndian.Uint32(ips)
	}
	return 0
}

// ConvertEndian convert bigEndian to littleEndian or littleEndian to bigEndian
func ConvertEndian(num uint32) uint32 {
	return ((num >> 24) & 0xff) | // move byte 3 to byte 0
		((num << 8) & 0xff0000) | // move byte 1 to byte 2
		((num >> 8) & 0xff00) | // move byte 2 to byte 1
		((num << 24) & 0xff000000)
}

// ItoIP convert ip from uint32 to string, like 174343044 to 10.100.67.132, if fail return empty string: ""
func ItoIP(ip uint32) string {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, ip)
	if err != nil {
		return ""
	}

	b := buf.Bytes()
	return fmt.Sprintf("%d.%d.%d.%d", b[0], b[1], b[2], b[3])
}

// AddrtoI convert address (ip:port) from string to uint32, like 10.100.67.132:8080 to 174343044, if fail return 0
func AddrtoI(addr string) uint32 {
	ip, _, err := net.SplitHostPort(addr)
	if err != nil {
		return 0
	}
	return IPtoI(ip)
}

// GetIP get local ip from inteface name like eth1
func GetIP(name string) uint32 {
	ifaces, err := net.Interfaces()
	if err != nil {
		return 0
	}

	for _, v := range ifaces {
		if v.Name == name {
			addrs, err := v.Addrs()
			if err != nil {
				return 0
			}

			for _, addr := range addrs {
				var ip net.IP
				switch val := addr.(type) {
				case *net.IPNet:
					ip = val.IP
				case *net.IPAddr:
					ip = val.IP
				}

				if len(ip) == 16 {
					return binary.BigEndian.Uint32(ip[12:16])
				} else if len(ip) == 4 {
					return binary.BigEndian.Uint32(ip)
				}
			}
		}
	}

	return 0
}

var localIP uint32
var littleEndianIP uint32

// GetLocalIP return local eth1 ip
func GetLocalIP() uint32 {
	if localIP != 0 {
		return localIP
	}
	localIP = GetIP("eth1")

	return localIP
}

// GetLittleEndianLocalIP return little endian local eth1 ip
func GetLittleEndianLocalIP() uint32 {
	if littleEndianIP != 0 {
		return littleEndianIP
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		return 0
	}

	for _, v := range ifaces {
		if v.Name == "eth1" {
			addrs, err := v.Addrs()
			if err != nil {
				return 0
			}

			for _, addr := range addrs {
				var ip net.IP
				switch val := addr.(type) {
				case *net.IPNet:
					ip = val.IP
				case *net.IPAddr:
					ip = val.IP
				}

				if len(ip) == 16 {
					littleEndianIP = binary.LittleEndian.Uint32(ip[12:16])
					return littleEndianIP
				} else if len(ip) == 4 {
					littleEndianIP = binary.LittleEndian.Uint32(ip)
					return littleEndianIP
				}
			}
		}
	}
	return littleEndianIP
}

func UrlEncode(value string) (result string) {
	v := url.Values{}
	v.Add("", value)
	body := v.Encode()
	result = body[1:]
	return
}

func UrlDecode(value string) (result string) {
	urlStr := "param=" + value
	m, _ := url.ParseQuery(urlStr)
	result = m.Get("param")
	return
}
