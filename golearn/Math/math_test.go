package golearn

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var num, a, b, c int

func TestMain(m *testing.M) {
	//before
	num = 0
	a = 0
	b = 0
	c = 0
	fmt.Printf("-------- Prefix is declared -------- \nnum = %d, \na = %d, \nb = %d, \nc = %d, \n", num, a, b, c)
	fmt.Println("------------------------------------")

	//main
	m.Run() // eksekusi semua unit test

	//after
	num = 0
	a = 0
	b = 0
	c = 0
	fmt.Printf("-------- Postfix is declared -------- \nnum = %d, \na = %d, \nb = %d, \nc = %d, \n", num, a, b, c)
	fmt.Println("-------------------------------------")
}

func TestAbsolute(t *testing.T) {

	t.Run("negative test case", func(t *testing.T) {
		c := Absolute(-5)
		assert.Equal(t, 5, c, "expect 5, got %d", c)
	})
	t.Run("positive test case", func(t *testing.T) {
		c := Absolute(5)
		assert.Equal(t, 5, c, "expect 5, got %d", c)
	})
}

func TestAbsolute_Table(t *testing.T) {
	testCases := []struct {
		name     string
		c        int
		expected int
	}{
		{
			name: "first negative",
			c:    -1, expected: 1,
		},
		{
			name: "second negative",
			c:    -5, expected: 5,
		},
		{
			name: "first positive",
			c:    2, expected: 2,
		},
		{
			name: "second positive",
			c:    7, expected: 7,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := Absolute(tc.c)
			assert.Equal(t, tc.expected, c)
		})
	}
}

func TestMultiply(t *testing.T) {

	t.Run("negative-negative test case", func(t *testing.T) {
		c := Multiply(-1, -2)
		assert.Equal(t, 2, c, "expect 2, got %d", c)
	})
	t.Run("negative-positive test case", func(t *testing.T) {
		c := Multiply(-2, 3)
		assert.Equal(t, -6, c, "expect -6, got %d", c)
	})
	t.Run("positive-negative test case", func(t *testing.T) {
		c := Multiply(3, -4)
		assert.Equal(t, -12, c, "expect -12, got %d", c)
	})
	t.Run("positive-positive test case", func(t *testing.T) {
		c := Multiply(4, 5)
		assert.Equal(t, 20, c, "expect 20, got %d", c)
	})

}

func TestMultiply_Table(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "negative and negative",
			a:        -1,
			b:        -1,
			expected: 1,
		},
		{
			name:     "negative and positive",
			a:        -5,
			b:        2,
			expected: -10,
		},
		{
			name:     "positive and positive",
			a:        3,
			b:        2,
			expected: 6,
		},
		{
			name:     "positive and negative",
			a:        2,
			b:        -6,
			expected: -12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := Multiply(tc.a, tc.b)
			assert.Equal(t, tc.expected, c)
		})
	}
}
