package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func readData(r io.Reader) ([][]int, error) {
	var data [][]int

	br := bufio.NewReaderSize(r, 128)

	for {
		l, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				return data, nil
			}

			return nil, err
		}

		d, err := parseLine(string(l))
		if err != nil {
			return nil, err
		}

		data = append(data, d)
	}
}

func parseLine(s string) ([]int, error) {
	ss := strings.Split(s, "\t")

	d := make([]int, len(ss))

	var err error
	for i, v := range ss {
		d[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
	}

	return d, nil
}
