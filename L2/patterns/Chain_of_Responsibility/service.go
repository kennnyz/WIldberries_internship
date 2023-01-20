package main

type Service interface {
	Execute(*Data)
	SetNext(Service)
}

type Data struct {
	GetSource    bool // Флаг который показывает были ли приняты данные обработчиком или нет
	UpdateStatus bool // Ставит статус тот
}
