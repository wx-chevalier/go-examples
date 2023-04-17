package fibonaccis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	assert.Equal(t, Fibonacci(0), 0)
	assert.Equal(t, Fibonacci(1), 1)
	assert.Equal(t, Fibonacci(2), 1)
	assert.Equal(t, Fibonacci(3), 2)
	assert.Equal(t, Fibonacci(4), 3)
	assert.Equal(t, Fibonacci(5), 5)
	assert.Equal(t, Fibonacci(6), 8)
	assert.Equal(t, Fibonacci(7), 13)
	assert.Equal(t, Fibonacci(8), 21)
	assert.Equal(t, Fibonacci(9), 34)
	assert.Equal(t, Fibonacci(10), 55)
}

func TestFibonacciRecursively(t *testing.T) {
	assert.Equal(t, FibonacciRecursively(0), 0)
	assert.Equal(t, FibonacciRecursively(1), 1)
	assert.Equal(t, FibonacciRecursively(2), 1)
	assert.Equal(t, FibonacciRecursively(3), 2)
	assert.Equal(t, FibonacciRecursively(4), 3)
	assert.Equal(t, FibonacciRecursively(5), 5)
	assert.Equal(t, FibonacciRecursively(6), 8)
	assert.Equal(t, FibonacciRecursively(7), 13)
	assert.Equal(t, FibonacciRecursively(8), 21)
	assert.Equal(t, FibonacciRecursively(9), 34)
	assert.Equal(t, FibonacciRecursively(10), 55)
}
