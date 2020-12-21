package main

import (
	"fmt"
	"testing"
)

func TestTemplate1(t *testing.T) {
	fmt.Println("---准备台湾餐---")
	d1 := NewTWDinner()
	d1.DoDinner()
	fmt.Println("---准备杭州餐---")
	d2 := NewHZDinner()
	d2.DoDinner()
	fmt.Println("---准备北京餐---")
	d3 := NewBJDinner()
	d3.DoDinner()
}

func TestTemplate2(t *testing.T) {
	fmt.Println("---准备台湾餐---")
	d1 := &TaiwanDinner{}
	DoDinner(d1)
	fmt.Println("---准备杭州餐---")
	d2 := &HangzhouDinner{}
	DoDinner(d2)
	fmt.Println("---准备北京餐---")
	d3 := &BeijingDinner{}
	DoDinner(d3)
}
