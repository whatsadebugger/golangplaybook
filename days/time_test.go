package days_test

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/whatsadebugger/golangplaybook/days"
)

func TestTomorrow(t *testing.T) {
	tom := days.Tomorrow()
	log.Println(tom.Format(time.RFC1123))
}

func TestTimeAssertion(t *testing.T) {
	got := time.Unix(1729825337, 0).UTC().Format(time.RFC3339)
	expected := "2024-10-25T03:02:17Z"
	assert.Equal(t, expected, got)
}
