package main

// Строитель — это порождающий паттерн проектирования,
//который позволяет создавать сложные объекты пошагово.
//Строитель даёт возможность использовать один и тот же
//код строительства для получения разных
//представлений объектов

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}

	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}
