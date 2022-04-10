package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_Humans(t *testing.T) {
	var humans Humans
	humans = append(humans, Human{
		ID:   primitive.ObjectID{},
		Name: "John Doe",
		Age:  35,
	}, Human{
		ID:   primitive.ObjectID{},
		Name: "Jane Doe",
		Age:  35,
	})
	assert.Equal(t, 2, len(humans))
}
