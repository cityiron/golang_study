package main

import "fmt"

type Dinner2 interface {
	foodEnough() bool
	doShopping()
	beforeCooking()
	doCooking() string
	afterCooking()
}

// AbstractDinner 类似于抽象类
type AbstractDinner struct {
}

func (AbstractDinner) foodEnough() bool {
	return true
}

func (AbstractDinner) doShopping() {
	fmt.Println("门口小贩买菜")
}

func (AbstractDinner) beforeCooking() {
}

func (AbstractDinner) doCooking() string {
	return ""
}

func (AbstractDinner) afterCooking() {
}

func DoDinner(d Dinner2) {
	if !d.foodEnough() {
		d.doShopping()
	}

	d.beforeCooking()
	fmt.Println(d.doCooking())
	d.afterCooking()
}

type HangzhouDinner struct {
	AbstractDinner
}

func (*HangzhouDinner) foodEnough() bool {
	return false
}

func (*HangzhouDinner) beforeCooking() {
	fmt.Println("杭州MM 在洗菜切菜")
}

func (*HangzhouDinner) doCooking() string {
	return "杭州MM 在做杭州菜"
}

func (*HangzhouDinner) afterCooking() {
	fmt.Println("杭州MM 让你去品尝")
}

type BeijingDinner struct {
	AbstractDinner
}

func (*BeijingDinner) beforeCooking() {
	fmt.Println("北京MM 在洗菜切菜")
}

func (*BeijingDinner) doCooking() string {
	return "北京MM 在做北京菜"
}

func (*BeijingDinner) afterCooking() {
	fmt.Println("北京MM 让你去品尝")
}

type TaiwanDinner struct {
	AbstractDinner
}

func (*TaiwanDinner) foodEnough() bool {
	return false
}

func (*TaiwanDinner) doShopping() {
	fmt.Println("生鲜超市购买，一定要买茶叶蛋")
}

func (*TaiwanDinner) beforeCooking() {
	fmt.Println("台湾MM 在洗菜切菜")
}

func (*TaiwanDinner) doCooking() string {
	return "台湾MM 在做台湾菜"
}

func (*TaiwanDinner) afterCooking() {
	fmt.Println("台湾MM 让你去品尝")
}
