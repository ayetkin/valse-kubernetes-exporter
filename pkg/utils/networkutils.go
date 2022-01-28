package utils

import "crypto/tls"

func TlsDial(url string) (func() tls.ConnectionState, error) {
	conn, err := tls.Dial("tcp", url, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		return nil, err
	}
	return conn.ConnectionState, err
}
