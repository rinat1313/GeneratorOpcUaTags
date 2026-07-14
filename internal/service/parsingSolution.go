package service

import (
	"encoding/xml"
	"generatorOPCUA/internal/tmpDomen"
	"os"
)

func ParsingSolution(filename string) (tmpDomen.Solution, error) {
	filename = "data/" + filename
	file, err := os.Open(filename)
	var sol tmpDomen.Solution
	if err != nil {
		return sol, err
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)

	for {
		t, err := decoder.Token()
		if err != nil {
			break
		}
		if se, ok := t.(xml.StartElement); ok {
			if err = decoder.DecodeElement(&sol, &se); err != nil {
				break
			}
		}
	}
	return sol, err
}

func GetSolutionObject(solution tmpDomen.Solution) (map[string]map[string]tmpDomen.TechnologyObject, error) {
	result := make(map[string]map[string]tmpDomen.TechnologyObject)

	for _, project := range solution.Projects.Project {
		for _, graphicScheme := range project.GraphicScheme {
			for _, graphicObject := range graphicScheme.GraphicObject {
				for _, technologyObject := range graphicObject.TechnologyObject {
					if _, ok := result[technologyObject.Template]; !ok {
						result[technologyObject.Template] = make(map[string]tmpDomen.TechnologyObject)
					}
					if _, ok := result[technologyObject.Template][technologyObject.Id]; !ok {
						result[technologyObject.Template][technologyObject.Id] = technologyObject
					}
				}
			}
		}
	}
	return result, nil
}

//	func ParsingSolution(path string) (map[string][]tmpDomen.TechnologyObject, error) {
//		regTemplate := `(?s)<TechnologyObject\s+[^>]*?(?:/>|>(?:.*?)</TechnologyObject>)`
//		fi, err := os.Open("data/" + path)
//		if err != nil {
//			fmt.Println(err)
//		}
//		defer fi.Close()
//		text, err := ioutil.ReadAll(fi)
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		//var objects []tmpDomen.TechProp
//		mapObjectsIsType := make(map[string][]tmpDomen.TechnologyObject)
//		var object tmpDomen.TechnologyObject
//
//		march := regexp.MustCompile(regTemplate)
//		matches := march.FindAll(text, -1)
//		for _, m := range matches {
//
//			err := xml.Unmarshal(m, &object)
//			if err != nil {
//				fmt.Printf("Ошибка при unmarshal: %s\n", err)
//			}
//
//			if _, ok := mapObjectsIsType[object.Template]; !ok {
//				mapObjectsIsType[object.Template] = make([]tmpDomen.TechnologyObject, 0)
//			}
//
//			mapObjectsIsType[object.Template] = append(mapObjectsIsType[object.Template], object)
//		}
//
//		return mapObjectsIsType, nil
//	}
func GetNameSolution(path string) (string, error) {
	//regSolution := "<Solution\\s+[^>]+>.*?</Solution>"
	////regSolution := "<Solution[^>]*\">"
	//fi, err := os.Open(path)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer fi.Close()
	//text, err := ioutil.ReadAll(fi)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//var solution tmpDomen.Solution
	//match := regexp.MustCompile(regSolution)
	//matches := match.FindAll(text, -1)
	//err = xml.Unmarshal(matches[0], &solution)
	//if err != nil {
	//	fmt.Printf("Ошибка парсинга : %s\n", err)
	//}
	//path = "data/" + path

	solution, err := ParsingSolution(path)
	if err != nil {
		return "", err
	}
	return solution.Name + " (" + solution.Description + ")", nil
}

func GetFilesIsDir(path string) ([]tmpDomen.FileSolution, error) {

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var result []tmpDomen.FileSolution

	for _, file := range files {
		if !file.IsDir() {
			var sol tmpDomen.FileSolution
			sol.NameFile = file.Name()
			result = append(result, sol)
		}
	}

	return result, nil
}
