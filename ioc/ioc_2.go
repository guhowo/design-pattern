package ioc

import "fmt"

/*
	本文件是面向接口编程。
	题目：
	模拟组装2台电脑，
	--- 抽象层 ---有显卡Card 方法display，有内存Memory 方法storage，有处理器CPU 方法calculate
	--- 实现层层 ---有 Intel因特尔公司、产品有(显卡、内存、CPU)，有 Kingston 公司， 产品有(内存3)，有 NVIDIA 公司， 产品有(显卡)
	--- 逻辑层 ---1. 组装一台Intel系列的电脑，并运行，2. 组装一台 Intel CPU Kingston内存 NVIDIA显卡的电脑，并运行
*/

//显卡接口，对显卡对象进行抽象
type ICard interface {
	Display()
}

//内存接口
type IMemory interface {
	Storage()
}

//CPU接口
type ICPU interface {
	Calculate()
}

type IntelCard struct{}
type IntelMemory struct{}
type IntelCPU struct{}
type KingstonMemory struct{}
type NVIDIACard struct{}

func (c *IntelCard) Display() {
	fmt.Println("Intel Card. Display ")
}

func (m *IntelMemory) Storage() {
	fmt.Println("Intel Memory. Storage ")
}

func (c *IntelCPU) Calculate() {
	fmt.Println("Intel CPU. Calculate ")
}

func (m *KingstonMemory) Storage() {
	fmt.Println("Kingston Memory. Storage ")
}

func (c *NVIDIACard) Display() {
	fmt.Println("NVIDIA Card. Display ")
}

type Computer struct {
	Card   ICard
	Memory IMemory
	CPU    ICPU
}

func NewComputer(card ICard, memory IMemory, cpu ICPU) *Computer {
	return &Computer{card, memory, cpu}
}

func (c *Computer) Work() {
	c.Card.Display()
	c.CPU.Calculate()
	c.Memory.Storage()
}

func Demo() {
	intelCard, nVIDIACard := &IntelCard{}, &NVIDIACard{}
	intelCPU := &IntelCPU{}
	intelMemory, kingstonMemory := &IntelMemory{}, &KingstonMemory{}
	pc1 := NewComputer(intelCard, intelMemory, intelCPU)
	pc2 := NewComputer(nVIDIACard, kingstonMemory, intelCPU)
	pc1.Work()
	pc2.Work()
}
