package main

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"
)

func parse(r io.Reader) ([]*program, error) {
	var programs []*program

	var br = bufio.NewReaderSize(r, 128)

	for {
		s, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				return programs, nil
			}

			return nil, err
		}

		p, err := parseProgram(string(s))
		if err != nil {
			return nil, err
		}

		programs = append(programs, p)
	}
}

var (
	programRegexp           = regexp.MustCompile(`([a-z]+) \(([0-9]+)\)( -> ([a-z]+(, [a-z]+)*))?`)
	errInvalidProgramString = errors.New("invalid program string")
)

func parseProgram(s string) (*program, error) {
	m := programRegexp.FindStringSubmatch(s)
	if m == nil {
		return nil, errInvalidProgramString
	}

	w, _ := strconv.Atoi(m[2])

	var sp []string
	if m[4] != "" {
		sp = strings.Split(m[4], ", ")
	}

	return &program{
		Name:        m[1],
		Weight:      w,
		SubPrograms: sp,
	}, nil
}
