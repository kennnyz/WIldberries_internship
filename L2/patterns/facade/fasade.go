package facade

import "strings"

//PC реализует pc и фасад.

type PC struct {
	monitor     *Monitor
	systemBlock *SystemBlock
	keyBoard    *Keyboard
}

func NewPC() *PC {
	return &PC{
		monitor:     &Monitor{},
		systemBlock: &SystemBlock{},
		keyBoard:    &Keyboard{},
	}
}

// Start возвращает то, что нужно чтобы играть в компьютер

func (p *PC) Start() string {
	result := []string{
		p.systemBlock.build(),
		p.monitor.turnOn(),
		p.keyBoard.connect(),
	}
	return strings.Join(result, "\n")
}

// Monitor реализует субсистему "Monitor"
type Monitor struct {
}

// Реализация включения
func (m *Monitor) turnOn() string {
	return "Monitor turn on"
}

// SystemBlock реализует субсистему "SystemBlock"
type SystemBlock struct {
}

// Реализация сборки компьютера
func (s *SystemBlock) build() string {
	return "Build System Block"
}

// Keyboard реализует субсистему "Keyboard"
type Keyboard struct {
}

// Реализация подключения
func (k *Keyboard) connect() string {
	return "Keyboard connect"
}
