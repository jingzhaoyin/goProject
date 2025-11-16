package main

import (
	"fmt"
	"math"
)

// Shape 接口定义
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 结构体定义
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle 实现 Shape 接口的 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle 实现 Shape 接口的 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 额外的便捷方法
func (r Rectangle) IsSquare() bool {
	return r.Width == r.Height
}

func (r Rectangle) String() string {
	if r.IsSquare() {
		return fmt.Sprintf("正方形(边长: %.2f)", r.Width)
	}
	return fmt.Sprintf("矩形(宽: %.2f, 高: %.2f)", r.Width, r.Height)
}

// Circle 结构体定义
type Circle struct {
	Radius float64
}

// Circle 实现 Shape 接口的 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle 实现 Shape 接口的 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("圆形(半径: %.2f)", c.Radius)
}

// Triangle 结构体（额外示例）
type Triangle struct {
	A, B, C float64 // 三条边
}

// Triangle 实现 Shape 接口
func (t Triangle) Area() float64 {
	// 使用海伦公式计算三角形面积
	s := t.Perimeter() / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func (t Triangle) String() string {
	return fmt.Sprintf("三角形(边: %.2f, %.2f, %.2f)", t.A, t.B, t.C)
}

// 工具函数：打印形状信息
func printShapeInfo(s Shape) {
	fmt.Printf("形状: %v\n", s)
	fmt.Printf("  面积: %.2f\n", s.Area())
	fmt.Printf("  周长: %.2f\n", s.Perimeter())
	fmt.Println("---")
}

// 工具函数：处理 Shape 切片
func processShapes(shapes []Shape) {
	fmt.Println("=== 处理形状集合 ===")
	totalArea := 0.0
	totalPerimeter := 0.0

	for i, shape := range shapes {
		fmt.Printf("形状 %d:\n", i+1)
		printShapeInfo(shape)
		totalArea += shape.Area()
		totalPerimeter += shape.Perimeter()
	}

	fmt.Printf("总面积: %.2f\n", totalArea)
	fmt.Printf("总周长: %.2f\n", totalPerimeter)
	fmt.Printf("平均面积: %.2f\n", totalArea/float64(len(shapes)))
	fmt.Printf("平均周长: %.2f\n", totalPerimeter/float64(len(shapes)))
}

// 类型断言示例
func describeShape(s Shape) {
	fmt.Printf("类型: %T\n", s)

	// 类型断言，检查具体类型
	switch shape := s.(type) {
	case Rectangle:
		if shape.IsSquare() {
			fmt.Println("  这是一个正方形")
		} else {
			fmt.Println("  这是一个矩形")
		}
	case Circle:
		fmt.Println("  这是一个圆形")
	case Triangle:
		fmt.Println("  这是一个三角形")
	default:
		fmt.Println("  未知形状")
	}
}

func main() {
	fmt.Println("=== Go 接口示例：形状计算 ===\n")

	// 创建 Rectangle 实例
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Println("1. 矩形测试:")
	printShapeInfo(rect)
	describeShape(rect)

	// 创建正方形
	square := Rectangle{Width: 4, Height: 4}
	fmt.Println("2. 正方形测试:")
	printShapeInfo(square)
	describeShape(square)

	// 创建 Circle 实例
	circle := Circle{Radius: 3}
	fmt.Println("3. 圆形测试:")
	printShapeInfo(circle)
	describeShape(circle)

	// 创建 Triangle 实例
	triangle := Triangle{A: 3, B: 4, C: 5}
	fmt.Println("4. 三角形测试:")
	printShapeInfo(triangle)
	describeShape(triangle)

	// 使用接口类型的切片
	fmt.Println("\n=== 使用接口切片 ===")
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Circle{Radius: 2.5},
		Rectangle{Width: 4, Height: 4}, // 正方形
		Triangle{A: 5, B: 12, C: 13},   // 直角三角形
		Circle{Radius: 1},
	}

	processShapes(shapes)

	// 接口类型检查示例
	fmt.Println("\n=== 接口类型检查 ===")
	checkInterfaceImplementation()

	// 空接口示例
	fmt.Println("\n=== 空接口示例 ===")
	useEmptyInterface()
}

// 检查接口实现
func checkInterfaceImplementation() {
	var s Shape

	// Rectangle 实现了 Shape 接口
	rect := Rectangle{Width: 3, Height: 4}
	s = rect // 这是有效的
	fmt.Printf("Rectangle 实现了 Shape: %v\n", s.Area())

	// 编译时检查
	var _ Shape = (*Rectangle)(nil) // 这行代码会在编译时检查 Rectangle 是否实现了 Shape
	var _ Shape = (*Circle)(nil)
	var _ Shape = (*Triangle)(nil)

	fmt.Println("所有类型都正确实现了 Shape 接口")
}

// 空接口示例（可以存储任何类型）
func useEmptyInterface() {
	// 空接口可以存储任何值
	var anything interface{}

	anything = Rectangle{Width: 2, Height: 3}
	fmt.Printf("存储矩形: %v\n", anything)

	anything = Circle{Radius: 5}
	fmt.Printf("存储圆形: %v\n", anything)

	anything = "这是一个字符串"
	fmt.Printf("存储字符串: %v\n", anything)

	anything = 42
	fmt.Printf("存储整数: %v\n", anything)

	// 类型断言从空接口获取具体值
	if shape, ok := anything.(Shape); ok {
		fmt.Printf("这是一个形状，面积: %.2f\n", shape.Area())
	} else {
		fmt.Println("这不是一个形状")
	}
}

// 工厂函数示例
func CreateShape(shapeType string, params ...float64) Shape {
	switch shapeType {
	case "rectangle":
		if len(params) >= 2 {
			return Rectangle{Width: params[0], Height: params[1]}
		}
	case "circle":
		if len(params) >= 1 {
			return Circle{Radius: params[0]}
		}
	case "triangle":
		if len(params) >= 3 {
			return Triangle{A: params[0], B: params[1], C: params[2]}
		}
	}
	return nil
}

// 在 main 函数末尾添加工厂函数测试
func testFactoryFunction() {
	fmt.Println("\n=== 工厂函数测试 ===")

	shapes := []Shape{
		CreateShape("rectangle", 3, 4),
		CreateShape("circle", 2.5),
		CreateShape("triangle", 3, 4, 5),
	}

	for i, shape := range shapes {
		if shape != nil {
			fmt.Printf("工厂创建的形状 %d: ", i+1)
			printShapeInfo(shape)
		}
	}
}
