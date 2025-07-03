package main

import "fmt"

// 如果类名首字母大写，表示其他包也能够访问
type Hero struct {
	// 如果类的属性首字母大写，表示该属性对外能够访问的，否则只能够类的内部访问
	Name  string
	Ad    int
	Level int
}

func (h *Hero) Show() {
	fmt.Println("Name=", h.Name, "Ad=", h.Ad, "Level=", h.Level)
}

func (h *Hero) GetName() string {
	return h.Name
}

func (h *Hero) SetName(newName string) {
	h.Name = newName
}

func main() {
	hero := Hero{Name: "zhangsan", Ad: 100, Level: 20}
	hero.Show()

	hero.SetName("lisi")
	hero.Show()
}
