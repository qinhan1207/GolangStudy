/* 多态 */
package main

import "fmt"

// 本质是一个指针
type Animal interface {
	Sleep()
	GetColor() string // 获取动物的颜色
	GetType() string  // 获取动物的种类
}

// 具体的类
type Cat struct {
	color string
}

func (c *Cat) Sleep() {
	fmt.Println("cat is sleeping")
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "Cat"
}

// 具体的类
type Dog struct {
	color string
}

func (d *Dog) Sleep() {
	fmt.Println("dog is sleeping")
}

func (d *Dog) GetColor() string {
	return d.color
}

func (d *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal Animal) {
	animal.Sleep() // 多态
	fmt.Println("color=", animal.GetColor())
	fmt.Println("type=", animal.GetType())
}

func main() {
	/* var animal Animal = &Cat{"Green"} // 接口的数据类型，父类指针
	animal.Sleep()                    // 调用的就是Cat的Sleep()

	animal = &Dog{"Yellow"}
	animal.Sleep() // 调用Dog的Sleep方法 */
	cat := Cat{"Green"}
	dog := Dog{"Yellow"}
	showAnimal(&cat)
	showAnimal(&dog)

}
