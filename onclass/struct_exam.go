package main

import "fmt"

func main()  {
	realFish := Fish{Name: "golden"}
	realFish.Swim()

	fake := FakeFish{Name: "fake"}
	fake.FakeSwim()

	// 类型强转
	td := Fish(fake)
	td.Swim()

	sFake := StrongFakeFish{Name: "strong"}
	sFake.Swim()

	tds := Fish(sFake)
	// 真的变成了鱼
	tds.Swim()
}

type Fish struct {
	Name string
}

func (f Fish) Swim() {
	fmt.Printf("I'm a %v fish.\n", f.Name)
}

// FakeFish 定义新类型, 实现新方法
type FakeFish Fish

func (f FakeFish) FakeSwim() {
	fmt.Printf("我是山寨鱼，嘎嘎嘎\n")
}

type StrongFakeFish Fish
func (f StrongFakeFish) Swim() {
	fmt.Printf("我是华强北山寨鱼，嘎嘎嘎\n")
}