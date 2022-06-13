package utils

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"
)

func conv(i int) string {

	if i <= 0 {

		return ""

	}

	j := (i - 1) % 26

	l := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	s := conv((i - j) / 26)

	return s + string(l[j])
}

func GetAxis(row, col int) (string, error) {

	if col <= 0 {
		return "", fmt.Errorf("column value was invalid")
	}

	if row <= 0 {
		return "", fmt.Errorf("row value was invalid")

	}
	cstr := conv(col)
	rstr := strconv.Itoa(row)
	// log.Printf("%s = %d\n %s = %d", cstr, col, rstr, row)
	return cstr + rstr, nil
}

func GetIPFromCustomDNS(dns string, host string) (out string, err error) {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, dns)
		},
	}
	ip, err := r.LookupHost(context.Background(), host)
	if err != nil {
		return "", err
	}
	return ip[0], nil
}
