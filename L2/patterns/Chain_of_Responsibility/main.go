package main

//Цепочка обязанностей — это поведенческий паттерн проектирования, который позволяет передавать запросы последовательно по цепочке обработчиков.
//Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

func main() {
	device := &Device{Name: "Device 1"}
	updateSvc := &UpdateDataService{Name: "Update 1"}
	saveService := &DataSaveService{}
	device.SetNext(updateSvc)
	updateSvc.SetNext(saveService)

	Data := &Data{}
	device.Execute(Data)
}
