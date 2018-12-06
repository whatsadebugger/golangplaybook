package join

import (
	"errors"
	"fmt"
	"testing"
)

var ger error

func BenchmarkErrorsNew(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		err = errors.New("wowerror")
	}
	ger = err
}

func BenchmarkErrorf(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		err = fmt.Errorf("wowerror")
	}
	ger = err
}
