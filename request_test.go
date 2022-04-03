package main

import (
	"net/http"
	"strconv"
	"testing"
)

func BenchmarkRequestGin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, _ := http.Get("http://localhost:" + strconv.Itoa(portGin))
		if resp.StatusCode != http.StatusOK {
			print("error")
		}
	}
}
func BenchmarkRequestEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, _ := http.Get("http://localhost:" + strconv.Itoa(portEcho))
		if resp.StatusCode != http.StatusOK {
			print("error")
		}
	}
}
func BenchmarkRequestMux(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, _ := http.Get("http://localhost:" + strconv.Itoa(portMux))
		if resp.StatusCode != http.StatusOK {
			print("error")
		}
	}
}
