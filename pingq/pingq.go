package pingq

import (
	"fmt"
	"net"
	"time"
)

func Pingq(ip string, timeout time.Duration) (time.Duration, error) {
	const IcmpLen = 8
	msg := [32]byte{
		8, 0, 0, 0, 0, 13, 0, 37,
	}
	check := checkSum(msg[:IcmpLen])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	remoteAddr, err := net.ResolveIPAddr("ip", ip)
	if err != nil {
		return 0, err
	}
	conn, err := net.DialIP("ip:icmp", nil, remoteAddr)
	if err != nil {
		return 0, err
	}
	start := time.Now()
	if _, err := conn.Write(msg[:IcmpLen]); err != nil {
		return 0, err
	}
	conn.SetReadDeadline(time.Now().Add(timeout))
	n, err := conn.Read(msg[:])
	conn.SetReadDeadline(time.Time{})
	if err != nil {
		return 0, err
	} else {
		fmt.Println("PINGR:", msg[:n])
	}
	return time.Since(start), nil
}

func checkSum(msg []byte) uint16 {
	sum := 0
	for n := 0; n < len(msg); n += 2 {
		sum += int(msg[n])<<8 + int(msg[n+1])
	}
	sum = (sum >> 16) + sum&0xffff
	sum += sum >> 16
	return uint16(^sum)
}
