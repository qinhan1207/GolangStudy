/* 继承 */

package main

import "fmt"

// 定义父类
type Human struct {
	name string
	sex  string
}

func (h *Human) Eat() {
	fmt.Println("人吃饭")
}

func (h *Human) Walk() {
	fmt.Println("人走路")
}

// 定义子类
type SuperMan struct {
	Human // 表示SuperMan类继承了Human类
	level int
}

// 重定义父类的方法
func (s *SuperMan) Eat() {
	fmt.Println("超人也要吃饭")
}

// 子类的新方法
func (s *SuperMan) Fly() {
	fmt.Println("超人会飞")
}

func (s *SuperMan) Print() {
	fmt.Println("name=", s.name, "sex=", s.sex, "level=", s.level)
}

func main() {
	human := Human{"qinhan", "男"}
	human.Eat()
	human.Walk()

	// 定义一个子类对象
	// s := SuperMan{Human{"lisi", "male"}, 88}
	var s SuperMan
	s.name = "lisi"
	s.sex = "male"
	s.level = 88

	s.Walk()
	s.Eat()
	s.Fly()

	s.Print()
}
