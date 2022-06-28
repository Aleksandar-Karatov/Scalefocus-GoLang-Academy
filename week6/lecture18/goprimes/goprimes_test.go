package goprimes_test

import (
	"testing"
	"time"
	"week6Lecture18Tasks/week6/lecture18/goprimes"
)

func Benchmark100PrimesWith0MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goprimes.PrimesAndSleep(100, 0)
	}
}

func Benchmark100PrimesWith5MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goprimes.PrimesAndSleep(100, 5*time.Millisecond)
	}
}
func Benchmark100PrimesWith10MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goprimes.PrimesAndSleep(100, 10*time.Millisecond)
	}
}

func Benchmark100GoPrimesWith0MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goprimes.GoprimesAndSleep(100, 0)
	}
}
func Benchmark100GoPrimesWith5MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goprimes.GoprimesAndSleep(100, 5*time.Millisecond)
	}
}
func Benchmark100GoPrimesWith10MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goprimes.GoprimesAndSleep(100, 10*time.Millisecond)
	}
}
