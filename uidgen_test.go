package uidgen

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateGenerator(t *testing.T) {
	a := CreateGenerator(11)
	assert.NotNil(t, a)
}

func TestGenerator(t *testing.T) {
	a := CreateGenerator(11)
	assert.NotNil(t, a)
	id := a.GenUID()
	assert.True(t, id > 0)
	fmt.Printf("newid : %#v\n", id)
	id = a.GenUID()
	assert.True(t, id > 0)
	fmt.Printf("newid : %#v\n", id)
	assert.Equal(t, uint64(2), id&0xffffff)
	assert.Equal(t, uint64(11<<24), id&0xff000000)
	assert.True(t, (id&0xffffffff00000000) > 0)
}
