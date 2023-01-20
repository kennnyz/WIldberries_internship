package main

import "fmt"

// Отвечает за сохранение обработанных данных

type DataSaveService struct {
	Next Service
}

func (d *DataSaveService) Execute(data *Data) {
	if !data.GetSource {
		fmt.Println("Data not update") // Проверяется обработались ли данные или нет
		return
	}
	fmt.Println("Data saved")
}

func (d *DataSaveService) SetNext(service Service) {
	d.Next = service
}
