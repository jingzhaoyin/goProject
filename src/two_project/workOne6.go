package main

import (
	"fmt"
	"strings"
)

// Person 结构体
type Person struct {
	Name string
	Age  int
}

// Person 的方法
func (p Person) GetName() string {
	return p.Name
}

func (p Person) GetAge() int {
	return p.Age
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

// Employee 结构体，组合了 Person
type Employee struct {
	Person     // 匿名嵌入，相当于继承了 Person 的字段和方法
	EmployeeID string
	Department string
	Position   string
	Salary     float64
	Skills     []string
}

// Employee 的构造函数
func NewEmployee(name string, age int, employeeID, department, position string, salary float64) *Employee {
	return &Employee{
		Person: Person{
			Name: name,
			Age:  age,
		},
		EmployeeID: employeeID,
		Department: department,
		Position:   position,
		Salary:     salary,
		Skills:     []string{},
	}
}

// Employee 实现 PrintInfo 方法
func (e *Employee) PrintInfo() {
	fmt.Println("=== 员工信息 ===")
	fmt.Printf("员工ID: %s\n", e.EmployeeID)
	fmt.Printf("姓名: %s\n", e.Name) // 直接访问嵌入结构的字段
	fmt.Printf("年龄: %d岁\n", e.Age)
	fmt.Printf("部门: %s\n", e.Department)
	fmt.Printf("职位: %s\n", e.Position)
	fmt.Printf("薪资: ￥%.2f\n", e.Salary)
	fmt.Printf("成年状态: %v\n", e.IsAdult()) // 调用嵌入结构的方法

	if len(e.Skills) > 0 {
		fmt.Printf("技能: %s\n", strings.Join(e.Skills, ", "))
	} else {
		fmt.Println("技能: 暂无")
	}
	fmt.Println()
}

// Employee 的其他方法
func (e *Employee) AddSkill(skill string) {
	e.Skills = append(e.Skills, skill)
}

func (e *Employee) AddSkills(skills []string) {
	e.Skills = append(e.Skills, skills...)
}

func (e *Employee) GetSkills() []string {
	return e.Skills
}

func (e *Employee) Promote(newPosition string, newSalary float64) {
	oldPosition := e.Position
	e.Position = newPosition
	e.Salary = newSalary
	fmt.Printf("%s 从 %s 晋升为 %s，新薪资: ￥%.2f\n",
		e.Name, oldPosition, newPosition, newSalary)
}

func (e *Employee) Transfer(newDepartment string) {
	oldDept := e.Department
	e.Department = newDepartment
	fmt.Printf("%s 从 %s 部门调转到 %s 部门\n",
		e.Name, oldDept, newDepartment)
}

// Employee 的 String 方法
func (e *Employee) String() string {
	return fmt.Sprintf("Employee{ID: %s, Name: %s, Position: %s, Dept: %s}",
		e.EmployeeID, e.Name, e.Position, e.Department)
}

// Manager 结构体，继承 Employee
type Manager struct {
	Employee             // 嵌入 Employee
	TeamSize    int      // 团队人数
	TeamMembers []string // 团队成员
}

// Manager 的构造方法
func NewManager(emp *Employee, teamSize int) *Manager {
	return &Manager{
		Employee:    *emp,
		TeamSize:    teamSize,
		TeamMembers: []string{},
	}
}

// Manager 重写 PrintInfo 方法
func (m *Manager) PrintInfo() {
	fmt.Println("=== 经理信息 ===")
	m.Employee.PrintInfo() // 调用父类的方法
	fmt.Printf("团队规模: %d人\n", m.TeamSize)
	if len(m.TeamMembers) > 0 {
		fmt.Printf("团队成员: %s\n", strings.Join(m.TeamMembers, ", "))
	}
	fmt.Println()
}

// Manager 特有的方法
func (m *Manager) AddTeamMember(member string) {
	m.TeamMembers = append(m.TeamMembers, member)
	fmt.Printf("%s 加入了 %s 的团队\n", member, m.Name)
}

func (m *Manager) ConductMeeting() {
	fmt.Printf("经理 %s 正在主持会议，团队规模: %d人\n", m.Name, m.TeamSize)
}

// 接口定义
type InformationPrinter interface {
	PrintInfo()
}

// 工具函数：处理所有实现了 InformationPrinter 接口的类型
func PrintAllInfo(printers []InformationPrinter) {
	fmt.Println("\n=== 批量打印信息 ===")
	for i, printer := range printers {
		fmt.Printf("第 %d 个信息:\n", i+1)
		printer.PrintInfo()
	}
}

func main() {
	fmt.Println("=== Go 组合与结构体示例 ===\n")

	// 创建 Employee 实例
	emp1 := NewEmployee("张三", 28, "E1001", "技术部", "高级工程师", 25000.0)
	emp1.AddSkill("Go语言")
	emp1.AddSkill("分布式系统")
	emp1.AddSkill("Docker")

	emp2 := NewEmployee("李四", 22, "E1002", "市场部", "市场专员", 12000.0)
	emp2.AddSkills([]string{"市场分析", "PPT制作", "客户沟通"})

	emp3 := NewEmployee("王五", 35, "E1003", "财务部", "财务经理", 35000.0)
	emp3.AddSkills([]string{"财务分析", "预算管理", "税务筹划"})

	// 打印员工信息
	emp1.PrintInfo()
	emp2.PrintInfo()
	emp3.PrintInfo()

	// 测试晋升和调转
	fmt.Println("=== 员工变动 ===")
	emp1.Promote("技术专家", 30000.0)
	emp2.Transfer("销售部")
	fmt.Println()

	// 创建 Manager 实例
	manager := NewManager(emp3, 5)
	manager.AddTeamMember("赵六")
	manager.AddTeamMember("钱七")
	manager.PrintInfo()
	manager.ConductMeeting()

	// 使用接口多态
	fmt.Println("=== 使用接口多态 ===")
	var printers []InformationPrinter
	printers = append(printers, emp1, emp2, emp3, manager)
	PrintAllInfo(printers)

	// 演示组合的访问方式
	fmt.Println("=== 组合访问方式演示 ===")

	// 直接访问嵌入结构的字段
	fmt.Printf("员工1姓名: %s\n", emp1.Name)
	fmt.Printf("员工1年龄: %d\n", emp1.Age)

	// 通过嵌入结构名访问
	fmt.Printf("员工1 Person: %v\n", emp1.Person)

	// 调用嵌入结构的方法
	fmt.Printf("员工1是否成年: %v\n", emp1.IsAdult())
	fmt.Printf("员工1的姓名(通过方法): %s\n", emp1.GetName())

	// 类型断言示例
	fmt.Println("\n=== 类型检查 ===")
	checkTypes(emp1)
	checkTypes(manager)
}

// 类型检查函数
func checkTypes(obj interface{}) {
	switch v := obj.(type) {
	case *Employee:
		fmt.Printf("%s 是 Employee 类型\n", v.Name)
	case *Manager:
		fmt.Printf("%s 是 Manager 类型\n", v.Name)
	case *Person:
		fmt.Printf("%s 是 Person 类型\n", v.Name)
	default:
		fmt.Printf("未知类型: %T\n", v)
	}
}

// 部门结构体
type Department struct {
	Name      string
	Manager   *Manager
	Employees []*Employee
	Budget    float64
}

func (d *Department) PrintDepartmentInfo() {
	fmt.Printf("\n=== %s 部门信息 ===\n", d.Name)
	fmt.Printf("部门经理: %s\n", d.Manager.Name)
	fmt.Printf("员工数量: %d\n", len(d.Employees))
	fmt.Printf("部门预算: ￥%.2f\n", d.Budget)
	fmt.Printf("团队成员:\n")
	for _, emp := range d.Employees {
		fmt.Printf("  - %s (%s)\n", emp.Name, emp.Position)
	}
}

// 在 main 函数中添加部门示例
func demonstrateDepartment() {
	fmt.Println("\n=== 部门管理示例 ===")

	techManager := NewManager(NewEmployee("技术总监", 40, "M1001", "技术部", "技术总监", 50000.0), 10)
	techDept := &Department{
		Name:    "技术研发部",
		Manager: techManager,
		Employees: []*Employee{
			NewEmployee("程序员A", 25, "E2001", "技术部", "初级工程师", 15000.0),
			NewEmployee("程序员B", 28, "E2002", "技术部", "中级工程师", 20000.0),
			NewEmployee("程序员C", 32, "E2003", "技术部", "高级工程师", 28000.0),
		},
		Budget: 1000000.0,
	}

	techDept.PrintDepartmentInfo()
}
