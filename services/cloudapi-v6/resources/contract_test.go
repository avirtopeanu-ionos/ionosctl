package resources

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContractService(t *testing.T) {
	ctx := context.Background()
	t.Run("get_contract_error", func(t *testing.T) {
		svc := getTestClient(t)
		contractSvc := NewContractService(svc, ctx)
		_, _, err := contractSvc.Get(QueryParams{})
		assert.Error(t, err)
	})
}
