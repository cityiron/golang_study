package main

import (
	"fmt"
	"testing"
)

func TestTemplate1(t *testing.T) {
	fmt.Println("---准备台湾餐---")
	d1 := NewTWDinner("台湾")
	d1.DoDinner()
	fmt.Println("---准备杭州餐---")
	d2 := NewHZDinner("杭州")
	d2.DoDinner()
	fmt.Println("---准备北京餐---")
	d3 := NewBJDinner("北京")
	d3.DoDinner()
}

func TestTemplate2(t *testing.T) {
	fmt.Println("---准备台湾餐---")
	d1 := NewTaiwanDinner("台湾")
	d1.DoDinner(d1)
	fmt.Println("---准备杭州餐---")
	d2 := NewHangzhouDinner("杭州")
	d2.DoDinner(d2)
	fmt.Println("---准备北京餐---")
	d3 := NewTaiwanDinner("北京")
	d3.DoDinner(d3)
}
