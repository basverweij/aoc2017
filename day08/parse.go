package main

import (
	"bufio"
	"io"
)

func parse(r io.Reader) ([]*instruction, error) {
	var instructions []*instruction

	br := bufio.NewReaderSize(r, 64)
	for {
		l, _, err := br.ReadLine()
		if err != nil {
			if err == io.EOF {
				return instructions, nil
			}

			return nil, err
		}

		i, err := newInstructionFromSource(string(l))
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, i)
	}
}
