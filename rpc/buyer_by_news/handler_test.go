package buyer_by_news

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	res := inCompaniesAffected([]string{"ImportCraft International"}, "Oranges/ImportCraft International")

	assert.Equal(t, true, res)
}
