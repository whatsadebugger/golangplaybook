package days_test

import (
	"log"
	"testing"
	"time"

	"github.com/whatsadebugger/golangplaybook/days"
)

func TestTomorrow(t *testing.T) {
	tom := days.Tomorrow()
	log.Println(tom.Format(time.RFC1123))
}
