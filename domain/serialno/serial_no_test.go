package serialno

import (
	"chat-server/persistence"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextSerialNumber(t *testing.T) {
	persistence.InitMongoDB()
	defer func() {
		persistence.PostMongoDB()
	}()

	coll := persistence.Database().Collection(coll)
	err := coll.Drop(context.Background())
	assert.NoError(t, err)
	{
		for i := 1; i < 10; i++ {
			no, err := NextSerialNo("u1", "u2")
			assert.NoError(t, err)
			assert.Equal(t, int64(i), no)
		}
	}

}
