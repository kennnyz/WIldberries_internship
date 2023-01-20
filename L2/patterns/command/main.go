package main

import "fmt"

//	Команда - поведенческий паттерн, в котором запросы или операции являются отдельными объектами
/*
	Плюсы: нет связью между отправителем и исполнителем комманды, удобно манипулировать операциями:
		повтор, выстраивание в очередь, отложенный запуск запросов
	Минусы: Усложнение кода из-за добавления доп интерфейса и структур
*/
type Command interface {
	Execute()
}

type Light struct {
	status bool
}

func (l *Light) On() {
	l.status = true
	fmt.Println("Light is on! ")
}

func (l *Light) Off() {
	l.status = false
	fmt.Println("Light is off")
}

type LightOnCommand struct {
	light *Light
}

func NewLightOnCmd(light *Light) *LightOnCommand {
	return &LightOnCommand{light}
}

func (loc *LightOnCommand) Execute() {
	loc.light.On()
}

type LightOfCommand struct {
	light *Light
}

func NewLightOfCmd(light *Light) *LightOfCommand {
	return &LightOfCommand{light}
}

func (loc *LightOfCommand) Execute() {
	loc.light.Off()
}

type Button struct {
	cmd Command
}

func newButton(cmd Command) *Button {
	return &Button{cmd}
}
func (b *Button) press() {
	b.cmd.Execute()
}

func main() {
	light := &Light{}

	lightOn := NewLightOnCmd(light)
	lightOf := NewLightOfCmd(light)

	startButton := newButton(lightOn)
	shutdownbutton := newButton(lightOf)

	fmt.Println(light.status)
	startButton.press()
	fmt.Println(light.status)
	shutdownbutton.press()
	fmt.Println(light.status)
}
