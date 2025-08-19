package simplefactory

const (
	ShapeTypeCircle    string = "Circle"
	ShapeTypeRectangle string = "Rectangle"
)

// Shape 接口定义
type Shape interface {
	Draw() string
}

type Circle struct{}

func (c *Circle) Draw() string {
	return "Drawing Circle"
}

type Rectangle struct{}

func (r *Rectangle) Draw() string {
	return "Drawing Rectangle"
}

func NewShape(shapeType string) Shape {
	switch shapeType {
	case ShapeTypeCircle:
		return &Circle{}
	case ShapeTypeRectangle:
		return &Rectangle{}
	}
	return nil
}
