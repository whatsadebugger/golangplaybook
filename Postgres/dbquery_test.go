package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomItem(t *testing.T) {

	db := getDatabaseConnection()
	rows, err := RandomItem(db)

	assert.NoError(t, err)
	assert.True(t, rows.Next())

}
