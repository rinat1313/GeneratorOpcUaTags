package service

import (
	"encoding/xml"
	"fmt"
	"generatorOPCUA/internal/tmpDomen"
	"os"
	"regexp"
	"strings"
)

func ParseCommandXML() map[string]tmpDomen.ObjCommandIdTag {
	var result []tmpDomen.ControlCommand
	data, err := os.ReadFile("data/Solution.xml")
	if err != nil {
		panic(err)
	}

	templateCommand := "<ControlCommand\\s+[^>]*?/>"
	re := regexp.MustCompile(templateCommand)
	matches := re.FindAllString(string(data), -1)

	fmt.Printf("Найдено объектов: %d\n", len(matches))

	for _, command := range matches {
		var tmp tmpDomen.ControlCommand
		err := xml.Unmarshal([]byte(command), &tmp)
		if err != nil {
			panic(err)
		}
		result = append(result, tmp)
	}
	mappingCommand := map[string]tmpDomen.ObjCommandIdTag{}

	for _, command := range result {
		if command.Template != "ValveOpen" && command.Template != "ValveClose" && command.Template != "ValveStop" {
			continue
		}

		tmp := tmpDomen.ObjCommandIdTag{}
		tmp.Id = command.Target

		if value, ok := mappingCommand[command.Target]; ok {
			tmp = value
		}

		if command.Template == "ValveOpen" {
			if len(tmp.OpenId) == 0 {
				tmp.OpenId = command.Id
			}
		} else if command.Template == "ValveClose" {
			if len(tmp.CloseId) == 0 {
				tmp.CloseId = command.Id
			}
		} else if command.Template == "ValveStop" {
			if len(tmp.StopId) == 0 {
				tmp.StopId = command.Id
			}
		}
		mappingCommand[command.Target] = tmp
	}

	fmt.Printf("Итоговое количество задвижек: %d\n", len(mappingCommand))
	return mappingCommand
}

func ParseCommandToNameObj(nameSolution, body string, comma rune) map[string]map[string]tmpDomen.CommandStruct {

	solution, err := ParsingSolution(nameSolution)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	mapComm, err := MakeObjectCommandToString(body, comma)
	if err != nil {
		fmt.Println(err)
	}
	var commands = make(map[string]map[string]tmpDomen.CommandStruct, len(solution.ControlCommands.ControlCommand))
	fmt.Printf("Всего распарсеных комманд: %d\n", len(mapComm))
	fmt.Printf("Найдено объектов: %d\n", len(solution.ControlCommands.ControlCommand))

	for _, command := range solution.ControlCommands.ControlCommand {
		if _, ok := mapComm[strings.ToLower(command.Template)]; !ok {
			continue
		}
		tmp, _ := mapComm[strings.ToLower(command.Template)]
		tmp.Id = command.Id
		m, ok := commands[command.Target]
		if !ok {
			m = make(map[string]tmpDomen.CommandStruct)
		}
		m[command.Id] = tmp
		commands[command.Target] = m
	}

	fmt.Printf("Итоговое количество комманд: %d\n", len(commands))
	return commands
}

func ParseAutomationToBody(nameTemplate, nameSolution, body string, comma rune) map[string]map[string]tmpDomen.AutomationStruct {
	solution, err := ParsingSolution(nameSolution)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//
	//
	//fmt.Printf("Старт генерации защит для типа объектов: %s\n", nameTemplate)
	//var result []tmpDomen.Automation
	//
	//templateAutomation := "<Automation\\b[^>]*>.*?</Automation>"
	//matches, err := GetBodySolution(nameSolution, templateAutomation)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//for _, automation := range matches {
	//	var tmp tmpDomen.Automation
	//	err := xml.Unmarshal([]byte(automation), &tmp)
	//
	//	if err != nil {
	//		fmt.Printf("Ошибка с текстом: %s\n", automation)
	//	}
	//	result = append(result, tmp)
	//}
	//
	//// start
	//
	inf, err := MakeObjectsToString(body, comma)
	//
	//fmt.Printf("Количество объектов объектов: %d\n", len(inf))

	unicMap := make(map[string]tmpDomen.Automation)

	for _, automation := range solution.Automations.Automation {
		for _, value := range inf {
			if automation.Target == value.Id {
				if _, ok := unicMap[automation.Template]; !ok {
					unicMap[automation.Template] = automation
				}
			}
		}
	}

	mapAutomation, err := MakeObjectsAutomationToString(body, comma)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Размер мапы: %d\n", len(mapAutomation))
	var automations = make(map[string]map[string]tmpDomen.AutomationStruct)
	for _, automation := range solution.Automations.Automation {
		if _, ok := mapAutomation[automation.Template]; !ok {
			continue
		}
		tmp, _ := mapAutomation[automation.Template]
		tmp.Id = automation.Id
		m, ok := automations[automation.Target]
		if !ok {
			m = make(map[string]tmpDomen.AutomationStruct)
		}
		m[automation.Id] = tmp
		automations[automation.Target] = m
	}

	return automations
}
