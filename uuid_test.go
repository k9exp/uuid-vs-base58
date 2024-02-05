package main

import (
	"testing"

	"github.com/google/uuid"
)

func BenchmarkBase58(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := New(21)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = uuid.NewString()
	}
}
