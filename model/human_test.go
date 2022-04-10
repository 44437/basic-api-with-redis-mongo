package model

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_Human(t *testing.T) {
	var err error
	ID := strings.Repeat("1", 24)
	objectID, err := primitive.ObjectIDFromHex(ID)
	require.Nil(t, err)

	human := Human{
		ID:   objectID,
		Name: "John Doe",
		Age:  35,
	}

	humanByte, err := json.Marshal(human)

	require.Nil(t, err)
	assert.Equal(t,
		string(humanByte),
		fmt.Sprintf(
			`{"_id":"%s","name":"John Doe","age":35}`,
			ID,
		),
	)
}
