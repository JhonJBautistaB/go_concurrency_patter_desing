package main

import (
	"fmt"
	"log"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
}

// method constructor
func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

// methodo class Memory, evaluates if the finobacci exists for that number
func (m *Memory) Get(key int) (interface{}, error) {
	result, exists := m.cache[key]
	if !exists {
		result.value, result.err = m.f(key)
		m.cache[key] = result
	}
	return result.value, result.err
}

func GetFinobacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {
	cache := NewCache(GetFinobacci)
	fibo := []int{42, 40, 41, 42, 38}
	for _, n := range fibo {
		start := time.Now()
		value, err := cache.Get(n)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%d, %s, %d\n", n, time.Since(start), value)
	}
}
