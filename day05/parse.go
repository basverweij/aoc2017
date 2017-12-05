package main

import (
	"bufio"
	"io"
	"strconv"
)

func parse(r io.Reader) (jumps, error) {
	var js jumps

	br := bufio.NewReaderSize(r, 16)

	for {
		l, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				return js, nil
			}

			return nil, err
		}

		j, err := strconv.Atoi(string(l))
		if err != nil {
			return nil, err
		}

		js = append(js, j)
	}
}
