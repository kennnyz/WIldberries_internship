package main

import "fmt"

// Отвечает за обработку полученных данных

type UpdateDataService struct {
	Name string
	Next Service
}

func (d *UpdateDataService) Execute(data *Data) {
	if data.UpdateStatus {
		fmt.Printf("Data from device [%s] already updated\n", d.Name)
		d.Next.Execute(data)
		return
	}
	fmt.Printf("Update data from device [%s]\n", d.Name)
	data.UpdateStatus = true
	d.Next.Execute(data)
}

func (d *UpdateDataService) SetNext(service Service) {
	d.Next = service
}
