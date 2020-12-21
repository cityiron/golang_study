package main

import (
	"fmt"
)

type Dinner interface {
	DoDinner()
}

type AbstractCooking struct {
	foodEnough    func() bool
	doShopping    func()
	beforeCooking func()
	doCooking     func() string
	afterCooking  func()
	Name          string
}

func (d *AbstractCooking) DoDinner() {
	if !d.foodEnough() {
		d.doShopping()
	}
	d.beforeCooking()
	fmt.Println(d.doCooking())
	d.afterCooking()
	d.eat()
}

func (d *AbstractCooking) eat() {
	fmt.Println(fmt.Sprintf("%sMM说：开吃喽", d.Name))
}

func foodEnough() bool {
	return true
}

func doShopping() {
	fmt.Println("门口小贩买菜")
}

type HZDinner struct {
	AbstractCooking
}

func NewHZDinner(name string) Dinner {
	c := new(HZDinner)
	c.Name = name
	// 选择实现的
	c.AbstractCooking.foodEnough = c.foodEnough
	c.AbstractCooking.doShopping = doShopping
	// 必须实现的
	c.AbstractCooking.beforeCooking = c.beforeCooking
	c.AbstractCooking.doCooking = c.doCooking
	c.AbstractCooking.afterCooking = c.afterCooking
	return c
}

func (c *HZDinner) foodEnough() bool {
	return false
}

func (c *HZDinner) beforeCooking() {
	fmt.Println(fmt.Sprintf("%sMM 在洗菜切菜", c.Name))
}

func (c *HZDinner) doCooking() string {
	return fmt.Sprintf("%sMM 在做%s菜", c.Name, c.Name)
}

func (c *HZDinner) afterCooking() {
	fmt.Println(fmt.Sprintf("%sMM 让你去品尝", c.Name))
}

type TWDinner struct {
	AbstractCooking
}

func NewTWDinner(name string) Dinner {
	c := new(TWDinner)
	c.Name = name
	// 选择实现的
	c.AbstractCooking.foodEnough = c.foodEnough
	c.AbstractCooking.doShopping = c.doShopping
	// 必须实现的
	c.AbstractCooking.beforeCooking = c.beforeCooking
	c.AbstractCooking.doCooking = c.doCooking
	c.AbstractCooking.afterCooking = c.afterCooking
	return c
}

func (c *TWDinner) foodEnough() bool {
	return false
}

func (c *TWDinner) doShopping() {
	fmt.Println("生鲜超市购买，一定要买茶叶蛋")
}

func (c *TWDinner) beforeCooking() {
	fmt.Println(fmt.Sprintf("%sMM 在洗菜切菜", c.Name))
}

func (c *TWDinner) doCooking() string {
	return fmt.Sprintf("%sMM 在做%s菜", c.Name, c.Name)
}

func (c *TWDinner) afterCooking() {
	fmt.Println(fmt.Sprintf("%sMM 让你去品尝", c.Name))
}

type BJDinner struct {
	AbstractCooking
}

func NewBJDinner(name string) Dinner {
	c := new(BJDinner)
	c.Name = name
	// 选择实现的
	c.AbstractCooking.foodEnough = foodEnough
	c.AbstractCooking.doShopping = doShopping
	// 必须实现的
	c.AbstractCooking.beforeCooking = c.beforeCooking
	c.AbstractCooking.doCooking = c.doCooking
	c.AbstractCooking.afterCooking = c.afterCooking
	return c
}

func (c *BJDinner) beforeCooking() {
	fmt.Println(fmt.Sprintf("%sMM 在洗菜切菜", c.Name))
}

func (c *BJDinner) doCooking() string {
	return fmt.Sprintf("%sMM 在做%s菜", c.Name, c.Name)
}

func (c *BJDinner) afterCooking() {
	fmt.Println(fmt.Sprintf("%sMM 让你去品尝", c.Name))
}
