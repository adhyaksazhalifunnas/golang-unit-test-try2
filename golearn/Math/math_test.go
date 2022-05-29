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

func TestAdd(t *testing.T) {

	t.Run("negative test case", func(t *testing.T) {
		c := Add(-1, -2)
		assert.Equal(t, -3, c, "expect -3, got %d", c)
	})
	t.Run("positive test case", func(t *testing.T) {
		c := Add(1, -2)
		assert.Equal(t, -1, c, "expect -1, got %d", c)
	})

}

func TestAdd_Table(t *testing.T) {
	testCases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{
			name:     "negative and negative",
			a:        -1,
			b:        -1,
			expected: -2,
		},
		{
			name:     "negative and positive",
			a:        -1,
			b:        1,
			expected: 0,
		},
		{
			name:     "positive and positive",
			a:        1,
			b:        1,
			expected: 2,
		},
		{
			name:     "positive and negative",
			a:        1,
			b:        -1,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := Add(tc.a, tc.b)
			assert.Equal(t, tc.expected, c)
		})
	}
}
