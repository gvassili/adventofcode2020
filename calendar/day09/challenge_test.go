package day09

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const input = ``

var fullInput = func() []byte {
	r, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return b
}()

func TestChallenge_Prepare(t *testing.T) {
	var c Challenge
	err := c.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
}

func BenchmarkChallenge_Prepare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := bytes.NewReader(fullInput)
		var c Challenge
		c.Prepare(buf)
	}
}

func TestChallenge_Part1(t *testing.T) {
	var c Challenge
	err := c.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	r, err := c.Part1()
	assert.NoError(t, err)
	assert.Equal(t, "", r)
}

func BenchmarkChallenge_Part1(b *testing.B) {
	buf := bytes.NewReader(fullInput)
	var c Challenge
	c.Prepare(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Part1()
	}
}

func TestChallenge_Part2(t *testing.T) {
	var c Challenge
	err := c.Prepare(bytes.NewReader([]byte(input)))
	assert.NoError(t, err)
	r, err := c.Part2()
	assert.NoError(t, err)
	assert.Equal(t, "", r)
}

func BenchmarkChallenge_Part2(b *testing.B) {
	buf := bytes.NewReader(fullInput)
	var c Challenge
	c.Prepare(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Part2()
	}
}
