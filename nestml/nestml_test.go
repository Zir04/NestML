package nestml

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyNestMLTest(t *testing.T) {
	text := NestMLText("hello")
	assert.Equal(t, "hello", string(text))
}
