package main

import "fmt"

type Dinner2 interface {
	foodEnough() bool
	doShopping()
	beforeCooking()
	doCooking() string
	afterCooking()
	DoDinner(d Dinner2)
}

// AbstractDinner 类似于抽象类
type AbstractDinner struct {
	Dinner2
	Name string
}

func (AbstractDinner) foodEnough() bool {
	return true
}

func (AbstractDinner) doShopping() {
	fmt.Println("门口小贩买菜")
}

func (ad AbstractDinner) DoDinner(d Dinner2) {
	if !d.foodEnough() {
		d.doShopping()
	}

	d.beforeCooking()
	fmt.Println(d.doCooking())
	d.afterCooking()
	ad.eat()
}

func (ad AbstractDinner) eat() {
	fmt.Println(fmt.Sprintf("%sMM说：开吃喽", ad.Name))
}

type HangzhouDinner struct {
	AbstractDinner
}

func NewHangzhouDinner(name string) Dinner2 {
	return &HangzhouDinner{
		AbstractDinner{
			Name: name,
		},
	}
}

func (d *HangzhouDinner) foodEnough() bool {
	return false
}

func (d *HangzhouDinner) beforeCooking() {
	fmt.Println(fmt.Sprintf("%sMM 在洗菜切菜", d.Name))
}

func (d *HangzhouDinner) doCooking() string {
	return fmt.Sprintf("%sMM 在做%s菜", d.Name, d.Name)
}

func (d *HangzhouDinner) afterCooking() {
	fmt.Println(fmt.Sprintf("%sMM 让你去品尝", d.Name))
}

type BeijingDinner struct {
	AbstractDinner
}

func NewBeijingDinner(name string) Dinner2 {
	return &BeijingDinner{
		AbstractDinner{
			Name: name,
		},
	}
}

func (d *BeijingDinner) beforeCooking() {
	fmt.Println(fmt.Sprintf("%sMM 在洗菜切菜", d.Name))
}

func (d *BeijingDinner) doCooking() string {
	return fmt.Sprintf("%sMM 在做%s菜", d.Name, d.Name)
}

func (d *BeijingDinner) afterCooking() {
	fmt.Println(fmt.Sprintf("%sMM 让你去品尝", d.Name))
}

type TaiwanDinner struct {
	AbstractDinner
}

func NewTaiwanDinner(name string) Dinner2 {
	return &TaiwanDinner{
		AbstractDinner{
			Name: name,
		},
	}
}

func (d *TaiwanDinner) foodEnough() bool {
	return false
}

func (d *TaiwanDinner) doShopping() {
	fmt.Println("生鲜超市购买，一定要买茶叶蛋")
}

func (d *TaiwanDinner) beforeCooking() {
	fmt.Println(fmt.Sprintf("%sMM 在洗菜切菜", d.Name))
}

func (d *TaiwanDinner) doCooking() string {
	return fmt.Sprintf("%sMM 在做%s菜", d.Name, d.Name)
}

func (d *TaiwanDinner) afterCooking() {
	fmt.Println(fmt.Sprintf("%sMM 让你去品尝", d.Name))
}
