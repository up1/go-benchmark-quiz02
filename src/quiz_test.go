package main

import (
	"testing"
)

var dataSet = []string{
	"xxxxxxxxxxx{}xxxxxxxxxxxxxxx",
	"xxxxxxxxxxx{}xxxxxxx{}xxxxx",
	"xxxxxxxxxxx{xxxxxxx{}xxxxx",
	"xxxxxxxxxxx{}xxxxxxx{x}xxxxx",
	"xxxxxxxxxxx{x}xxxxxxxxxxxxxxx",
	"xxxxxxxxxxx}{xxxxxxxxxxxxxxx",
	"{xxxxxxxxxxxxxxx",
	"xxxxxxxxxxxxxxx}x",
}

func BenchmarkChonlasith(b *testing.B) {
	for j, n := 0, len(dataSet); j < n; j++ {
		for i := 0; i < b.N; i++ {
			chonlasith(dataSet[j])
		}
	}
}

func BenchmarkChonlasith2(b *testing.B) {
	for j, n := 0, len(dataSet); j < n; j++ {
		for i := 0; i < b.N; i++ {
			chonlasith2(dataSet[j])
		}
	}
}
