package join

import (
	"fmt"
	"strings"
	"testing"
)

var (
	testData    = []string{"a", "b", "c", "d", "e"}
	testString1 = "example.com/"
	testString2 = "login"
)

func BenchmarkStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := testString1 + testString2
		_ = s
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strings.Join([]string{testString1, testString2}, "")
		_ = s
	}
}

func BenchmarkStringSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("lakfjlkwajflkwajlkfjalkjfwalj%s/", testString2)
		_ = s
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strings.Join(testData, ":")
		_ = s
	}
}

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := fmt.Sprintf("%s:%s:%s:%s:%s", testData[0], testData[1], testData[2], testData[3], testData[4])
		_ = s
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := testData[0] + ":" + testData[1] + ":" + testData[2] + ":" + testData[3] + ":" + testData[4]
		_ = s
	}
}
