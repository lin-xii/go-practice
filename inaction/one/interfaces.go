package one

import "fmt"

type Animals interface {
	Sleep()
	Eat()
	Run()
}

type tiger struct{}

func (t *tiger) Sleep() {
	fmt.Println("Sleep")
}

func (t *tiger) Eat() {
	fmt.Println("Eat")
}

// interface 的实现，是严格的命名、params_type、return_type完全一致
// 和TS不太像，TS是只要类型声明一致，命名反而不重要。
// 相比之下，Go的更严格
// func (t *tiger) Run(s string) {
func (t *tiger) Run() {
	fmt.Println("Run")
}

type bird struct{}

func (b *bird) Sleep() {
	fmt.Println("bird Sleep")
}

func (b *bird) Eat() {
	fmt.Println("bird Eat")
}

func (b *bird) Run() {
	fmt.Println("bird can't run, but fly")
}

func RunInterface() {
	var animal Animals = &bird{}
	animal.Run()
}
