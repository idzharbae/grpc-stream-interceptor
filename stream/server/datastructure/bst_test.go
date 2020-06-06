package datastructure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	var bst *BST

	bst = Insert(bst, 1337)

	assert.NotNil(t, bst)
	assert.Equal(t, int32(1337), bst.Data)

	bst = Insert(bst, 1234)
	assert.NotNil(t, bst.Left)
	assert.Nil(t, bst.Right)
	assert.Equal(t, int32(1234), bst.Left.Data)

	bst = Insert(bst, 2)
	assert.NotNil(t, bst.Left.Left)
	assert.Nil(t, bst.Right)
	assert.Nil(t, bst.Left.Right)
	assert.Equal(t, int32(2), bst.Left.Left.Data)

	bst = Insert(bst, 1235)
	assert.NotNil(t, bst.Left.Right)
	assert.Nil(t, bst.Right)
	assert.Equal(t, int32(1235), bst.Left.Right.Data)
}

func TestToArray(t *testing.T) {
	var bst *BST
	bst = Insert(bst, 123)
	bst = Insert(bst, 3)
	bst = Insert(bst, 123123)
	bst = Insert(bst, 1000)
	bst = Insert(bst, 123124)
	bst = Insert(bst, 9)

	assert.Equal(t, []int32{3, 9, 123, 1000, 123123, 123124}, ToArray(bst))
}
