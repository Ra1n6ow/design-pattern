package simplefactory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircle(t *testing.T) {
	assert := assert.New(t)
	circle := NewShape(ShapeTypeCircle)
	actual := circle.Draw()
	assert.Equal("Drawing Circle", actual)
}

func TestRectangle(t *testing.T) {
	assert := assert.New(t)
	rectangle := NewShape(ShapeTypeRectangle)
	actual := rectangle.Draw()
	assert.Equal("Drawing Rectangle", actual)
}
