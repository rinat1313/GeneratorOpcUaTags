package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	//"generatorOPCUA/internal/controller"
	"generatorOPCUA/internal/service"
)

func main() {
	TestRidingSolution()
	//control := controller.Controller{}
	//control.StartWork()
}

func TestRidingSolution() {
	mapTypeObj, err := ParseCSV("tmpTestData/objects.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	PathSol := "TestSolution.xml"
	sol, err := service.ParsingSolution(PathSol)
	if err != nil {
		fmt.Printf("Ошибка парсинга файла: %s\n", err)
		return
	}

	mapObj, err := service.GetSolutionObject(sol)
	if err != nil {
		fmt.Printf("Ошибка перебора решения: %s\n", err)
		return
	}
	fmt.Printf("Количество типов Распарсило - Должно быть: %d - %d.\n", len(mapObj), len(mapTypeObj))
	if len(mapObj) == len(mapTypeObj) {
		fmt.Println("Количество типов совпадает.")
	} else {
		fmt.Println("Количество типов НЕ совпадает.")
		fmt.Printf("\nНашло: %d шт\n", len(mapObj))
		for nameType, _ := range mapObj {
			fmt.Println(nameType)
		}
		fmt.Printf("\nДолжно быть: %d шт\n", len(mapTypeObj))
		for nameType, _ := range mapTypeObj {
			fmt.Println(nameType)
		}
	}

}

// ParseCSV читает CSV-файл и возвращает map[тип]map[id]true.
// Первая строка файла пропускается (считается заголовком).
func ParseCSV(filePath string) (map[string]map[string]bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'

	// Пропускаем строку заголовков
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("не удалось прочитать заголовок: %w", err)
	}

	result := make(map[string]map[string]bool)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("ошибка чтения строки: %w", err)
		}

		// Пропускаем строки, в которых меньше двух колонок
		if len(record) < 2 {
			continue
		}

		objType := record[0]
		objID := record[1]

		// Если для этого типа ещё нет внутренней карты — создаём
		if _, exists := result[objType]; !exists {
			result[objType] = make(map[string]bool)
		}

		// Добавляем ID (значение true просто обозначает наличие)
		result[objType][objID] = true
	}

	return result, nil
}
