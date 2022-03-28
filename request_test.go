package main

import (
	"net/http"
	"testing"
)

func BenchmarkRequestGin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		http.Get("http://localhost:8081")
	}
}
func BenchmarkRequestEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		http.Get("http://localhost:8082")
	}
}
func BenchmarkRequestMux(b *testing.B) {
	for i := 0; i < b.N; i++ {
		http.Get("http://localhost:8083")
	}
}
