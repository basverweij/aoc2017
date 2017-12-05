package main

import (
	"bufio"
	"io"
	"strings"
)

type passphrase []string

func newPassphrase(s string) passphrase {
	return strings.Split(s, " ")
}

func newPassphrases(r io.Reader) ([]passphrase, error) {
	var ps []passphrase

	br := bufio.NewReaderSize(r, 1024)

	for {
		l, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				return ps, nil
			}

			return nil, err
		}

		ps = append(ps, newPassphrase(string(l)))
	}
}
