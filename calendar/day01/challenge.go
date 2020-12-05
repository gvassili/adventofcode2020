package day01

import (
	"bufio"
	"errors"
	"io"
	"strconv"
)

type Day01 struct {
	input    []int
	inputMap []int
}

func (d *Day01) Day() int {
	return 1
}

func (d *Day01) Prepare(r io.Reader) error {
	d.input = make([]int, 0, 2048)
	d.inputMap = make([]int, 2048)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		d.input = append(d.input, n)
		d.inputMap[n] = n
	}
	return scanner.Err()
}

func (d *Day01) Part1() (string, error) {
	for _, n := range d.input {
		t := 2020 - n
		if t >= 0 && t == d.inputMap[t] {
			return strconv.Itoa((2020 - n) * n), nil
		}
	}
	return "", errors.New("could not find result")
}

func (d *Day01) Part2() (string, error) {
	for i, n := range d.input {
		for _, m := range d.input[i:] {
			t := 2020 - n - m
			if t > 0 && t == d.inputMap[t] {
				return strconv.Itoa(t * n * m), nil
			}
		}
	}
	return "", errors.New("could not find result")
}