package main

import (
	"net/http"
	"strconv"
	"testing"
)

func BenchmarkRequestFiber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, _ := http.Get("http://localhost:" + strconv.Itoa(portFiber))
		if resp.StatusCode != http.StatusOK {
			print("error")
		}
		resp.Body.Close()
	}
}

func BenchmarkRequestFastHttpRouting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, _ := http.Get("http://localhost:" + strconv.Itoa(portFastHttpRouting))
		if resp.StatusCode != http.StatusOK {
			print("error")
		}
		resp.Body.Close()
	}
}

func BenchmarkRequestGin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, _ := http.Get("http://localhost:" + strconv.Itoa(portGin))
		if resp.StatusCode != http.StatusOK {
			print("error")
		}
		resp.Body.Close()
	}
}
func BenchmarkRequestEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, _ := http.Get("http://localhost:" + strconv.Itoa(portEcho))
		if resp.StatusCode != http.StatusOK {
			print("error")
		}
		resp.Body.Close()
	}
}
func BenchmarkRequestMux(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, _ := http.Get("http://localhost:" + strconv.Itoa(portMux))
		if resp.StatusCode != http.StatusOK {
			print("error")
		}
		resp.Body.Close()
	}
}
