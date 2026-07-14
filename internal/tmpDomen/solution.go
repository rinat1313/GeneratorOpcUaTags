package tmpDomen

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type Object struct {
	Id  string
	Tag string
}

type DpeTag struct {
	Dpe string
	Tag string
}

type FileSolution struct {
	NameFile string
	Discript string
}

type ObjCommandIdTag struct {
	Id      string
	OpenId  string
	CloseId string
	StopId  string
}

type CommandStruct struct {
	Id             string
	NameTemplate   string
	AfterHeaderTag string
}

type AutomationStruct struct {
	Id             string
	NameTemplate   string
	AfterHeaderTag string
}

type AnchorPoint struct {
	X string `xml:"X,attr"` //ok
	Y string `xml:"Y,attr"` //ok
}

type Automation struct {
	Description        string               `xml:"Description,attr"`   //ok
	ExternalCode       string               `xml:"ExternalCode,attr"`  //ok
	Id                 string               `xml:"Id,attr"`            //ok
	IdFromSource       string               `xml:"IdFromSource,attr"`  //ok
	Name               string               `xml:"Name,attr"`          //ok
	Tag                string               `xml:"Tag,attr"`           //ok
	Target             string               `xml:"Target,attr"`        //ok
	Template           string               `xml:"Template,attr"`      //ok
	TechnologyProperty []TechnologyProperty `xml:"TechnologyProperty"` //ok
}

type Automations struct {
	Automation []Automation `xml:"Automation"` //ok
}

type ControlCommand struct {
	Description        string               `xml:"Description,attr"`   //ok
	ExternalCode       string               `xml:"ExternalCode,attr"`  //ok
	Id                 string               `xml:"Id,attr"`            //ok
	IdFromSource       string               `xml:"IdFromSource,attr"`  //ok
	Name               string               `xml:"Name,attr"`          //ok
	Tag                string               `xml:"Tag,attr"`           //ok
	Target             string               `xml:"Target,attr"`        //ok
	Template           string               `xml:"Template,attr"`      //ok
	TechnologyProperty []TechnologyProperty `xml:"TechnologyProperty"` //ok
}

type ControlCommands struct {
	ControlCommand []ControlCommand `xml:"ControlCommand"` //ok
}

type GraphicLabel struct {
	AutoLocation         string      `xml:"AutoLocation,attr"`         //ok
	Description          string      `xml:"Description,attr"`          //ok
	ExternalCode         string      `xml:"ExternalCode,attr"`         //ok
	Id                   string      `xml:"Id,attr"`                   //ok
	IdFromSource         string      `xml:"IdFromSource,attr"`         //ok
	Name                 string      `xml:"Name,attr"`                 //ok
	Tag                  string      `xml:"Tag,attr"`                  //ok
	VisualDeclarationKey string      `xml:"VisualDeclarationKey,attr"` //ok
	AnchorPoint          AnchorPoint `xml:"AnchorPoint"`               //ok
}

type GraphicLabels struct {
	GraphicLabel []GraphicLabel `xml:"GraphicLabel"` //ok
}

type GraphicObject struct {
	Angle                 string             `xml:"Angle,attr"`                 //ok
	Description           string             `xml:"Description,attr"`           //ok
	ExternalCode          string             `xml:"ExternalCode,attr"`          //ok
	Id                    string             `xml:"Id,attr"`                    //ok
	IdFromSource          string             `xml:"IdFromSource,attr"`          //ok
	IsHorizontalReflected string             `xml:"IsHorizontalReflected,attr"` //ok
	IsLocked              string             `xml:"IsLocked,attr"`              //ok
	IsReversed            string             `xml:"IsReversed,attr"`            //ok
	IsVerticalReflected   string             `xml:"IsVerticalReflected,attr"`   //ok
	IsVisibled            string             `xml:"IsVisibled,attr"`            //ok
	LinkedGraphicScheme   string             `xml:"LinkedGraphicScheme,attr"`   //ok
	Name                  string             `xml:"Name,attr"`                  //ok
	Tag                   string             `xml:"Tag,attr"`                   //ok
	AnchorPoint           []AnchorPoint      `xml:"AnchorPoint"`                //ok
	GraphicLabels         []GraphicLabels    `xml:"GraphicLabels"`              //ok
	Points                []Points           `xml:"Points"`                     //ok
	TechnologyObject      []TechnologyObject `xml:"TechnologyObject"`           //ok
}

type GraphicScheme struct {
	Description   string          `xml:"Description,attr"`  //ok
	ExternalCode  string          `xml:"ExternalCode,attr"` //ok
	Id            string          `xml:"Id,attr"`           //ok
	IdFromSource  string          `xml:"IdFromSource,attr"` //ok
	Name          string          `xml:"Name,attr"`         //ok
	Tag           string          `xml:"Tag,attr"`          //ok
	Template      string          `xml:"Template,attr"`     //ok
	GraphicObject []GraphicObject `xml:"GraphicObject"`     //ok
}

type ObjectReference struct {
	Description  string `xml:"Description,attr"`
	ExternalCode string `xml:"ExternalCode,attr"`
	Id           string `xml:"Id,attr"`
	IdFromSource string `xml:"IdFromSource,attr"`
	Name         string `xml:"Name,attr"`
	Tag          string `xml:"Tag,attr"`
	Target       string `xml:"Target,attr"`
}

type ObjectReferences struct {
	ObjectReference []ObjectReference `xml:"ObjectReference"` // ??????????????????????????????????????????????????
}

type Point struct {
	X string `xml:"X,attr"` //ok
	Y string `xml:"Y,attr"` //ok
}

type Points struct {
	Point []Point `xml:"Point"` //ok
}

type Project struct {
	Description   string          `xml:"Description,attr"`  //ok
	ExternalCode  string          `xml:"ExternalCode,attr"` //ok
	Id            string          `xml:"Id,attr"`           //ok
	IdFromSource  string          `xml:"IdFromSource,attr"` //ok
	Name          string          `xml:"Name,attr"`         //ok
	Tag           string          `xml:"Tag,attr"`          //ok
	GraphicScheme []GraphicScheme `xml:"GraphicScheme"`     //ok
}

type Projects struct {
	Project []Project `xml:"Project"`
}

type Scenarios struct {
}

type Solution struct {
	Company                     string                      `xml:"Company,attr"`                //ok
	Description                 string                      `xml:"Description,attr"`            //ok
	ExternalCode                string                      `xml:"ExternalCode,attr"`           //ok
	Id                          string                      `xml:"Id,attr"`                     //ok
	Name                        string                      `xml:"Name,attr"`                   //ok
	StepIterationsCount         string                      `xml:"StepIterationsCount,attr"`    //ok
	Tag                         string                      `xml:"Tag,attr"`                    //ok
	Automations                 Automations                 `xml:"Automations"`                 // ok
	ControlCommands             ControlCommands             `xml:"ControlCommands"`             // ok
	Projects                    Projects                    `xml:"Projects"`                    // ok
	Scenarios                   Scenarios                   `xml:"Scenarios"`                   // ok
	TargetAreas                 TargetAreas                 `xml:"TargetAreas"`                 // ok
	TechnologyObjectsLinks      TechnologyObjectsLinks      `xml:"TechnologyObjectsLinks"`      // ok
	TechnologyPassports         TechnologyPassports         `xml:"TechnologyPassports"`         // ok
	TechnologyStandaloneObjects TechnologyStandaloneObjects `xml:"TechnologyStandaloneObjects"` // ok
	Workspaces                  Workspaces                  `xml:"Workspaces"`                  // ok
}

type TargetArea struct {
	Description  string `xml:"Description,attr"`  //ok
	ExternalCode string `xml:"ExternalCode,attr"` //ok
	Id           string `xml:"Id,attr"`           //ok
	IdFromSource string `xml:"IdFromSource,attr"` //ok
	Index        string `xml:"Index,attr"`        //ok
	Name         string `xml:"Name,attr"`         //ok
	Tag          string `xml:"Tag,attr"`          //ok
}

type TargetAreas struct {
	TargetArea []TargetArea `xml:"TargetArea"` //ok
}

type TechnologyListItem struct {
	Description        string               `xml:"Description,attr"`  //ok
	ExternalCode       string               `xml:"ExternalCode,attr"` //ok
	Id                 string               `xml:"Id,attr"`           //ok
	IdFromSource       string               `xml:"IdFromSource,attr"` //ok
	Name               string               `xml:"Name,attr"`         //ok
	Tag                string               `xml:"Tag,attr"`          //ok
	Template           string               `xml:"Template,attr"`     //ok
	TechnologyProperty []TechnologyProperty `xml:"TechnologyProperty"`
}

type TechnologyObject struct {
	Description        string               `xml:"Description,attr"`  //ok
	ExternalCode       string               `xml:"ExternalCode,attr"` //ok
	Id                 string               `xml:"Id,attr"`           //ok
	IdFromSource       string               `xml:"IdFromSource,attr"` //ok
	Name               string               `xml:"Name,attr"`         //ok
	Tag                string               `xml:"Tag,attr"`          //ok
	Template           string               `xml:"Template,attr"`     //ok
	TechnologyProperty []TechnologyProperty `xml:"TechnologyProperty"`
}

type TechnologyObjectLink struct {
	Description      string `xml:"Description,attr"`      //ok
	ExternalCode     string `xml:"ExternalCode,attr"`     //ok
	Id               string `xml:"Id,attr"`               //ok
	IdFromSource     string `xml:"IdFromSource,attr"`     //ok
	Name             string `xml:"Name,attr"`             //ok
	Tag              string `xml:"Tag,attr"`              //ok
	TechnologyObject string `xml:"TechnologyObject,attr"` //ok
}

type TechnologyObjectsLink struct {
	Description       string `xml:"Description,attr"`       //ok
	ExternalCode      string `xml:"ExternalCode,attr"`      //ok
	Id                string `xml:"Id,attr"`                //ok
	IdFromSource      string `xml:"IdFromSource,attr"`      //ok
	Name              string `xml:"Name,attr"`              //ok
	Tag               string `xml:"Tag,attr"`               //ok
	TechnologyObject1 string `xml:"TechnologyObject1,attr"` //ok
	TechnologyObject2 string `xml:"TechnologyObject2,attr"` //ok
	Template          string `xml:"Template,attr"`          //ok
}

type TechnologyObjectsLinks struct {
	TechnologyObjectsLink []TechnologyObjectsLink `xml:"TechnologyObjectsLink"` //ok
}

type TechnologyPassport struct {
	Description        string               `xml:"Description,attr"`   //ok
	ExternalCode       string               `xml:"ExternalCode,attr"`  //ok
	Id                 string               `xml:"Id,attr"`            //ok
	IdFromSource       string               `xml:"IdFromSource,attr"`  //ok
	Name               string               `xml:"Name,attr"`          //ok
	Tag                string               `xml:"Tag,attr"`           //ok
	Template           string               `xml:"Template,attr"`      //ok
	TechnologyProperty []TechnologyProperty `xml:"TechnologyProperty"` //ok
}

type TechnologyPassports struct {
	TechnologyPassport []TechnologyPassport `xml:"TechnologyPassport"` //ok
}

type TechnologyProperty struct {
	BooleanValue         string                 `xml:"BooleanValue,attr"`    //ok
	Description          string                 `xml:"Description,attr"`     //ok
	EnumItemValue        string                 `xml:"EnumItemValue,attr"`   //ok
	ExternalCode         string                 `xml:"ExternalCode,attr"`    //ok
	Id                   string                 `xml:"Id,attr"`              //ok
	IdFromSource         string                 `xml:"IdFromSource,attr"`    //ok
	IntegerValue         int                    `xml:"IntegerValue,attr"`    //ok
	LinkValue            string                 `xml:"LinkValue,attr"`       //ok
	Name                 string                 `xml:"Name,attr"`            //ok
	PointsValue          string                 `xml:"PointsValue,attr"`     //ok
	RealValue            CommaFloat64           `xml:"RealValue,attr"`       //ok
	StringValue          string                 `xml:"StringValue,attr"`     //ok
	Tag                  string                 `xml:"Tag,attr"`             //ok
	Template             string                 `xml:"Template,attr"`        //ok
	TechnologyListItem   []TechnologyListItem   `xml:"TechnologyListItem"`   //ok
	TechnologyObjectLink []TechnologyObjectLink `xml:"TechnologyObjectLink"` //ok
	TechnologyPassport   []TechnologyPassport   `xml:"TechnologyPassport"`   //ok
}

type CommaFloat64 float64

func (cf *CommaFloat64) UnmarshalXMLAttr(attr xml.Attr) error {
	s := strings.Replace(attr.Value, ",", ".", 1)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*cf = CommaFloat64(f)
	return nil
}

type TechnologyStandaloneObject struct {
	Description        string               `xml:"Description,attr"`   //ok
	ExternalCode       string               `xml:"ExternalCode,attr"`  //ok
	Id                 string               `xml:"Id,attr"`            //ok
	IdFromSource       string               `xml:"IdFromSource,attr"`  //ok
	Name               string               `xml:"Name,attr"`          //ok
	Tag                string               `xml:"Tag,attr"`           //ok
	Template           string               `xml:"Template,attr"`      //ok
	TechnologyProperty []TechnologyProperty `xml:"TechnologyProperty"` //ok
}

type TechnologyStandaloneObjects struct {
	TechnologyStandaloneObject []TechnologyStandaloneObject `xml:"TechnologyStandaloneObject"` //ok
}

type Workspace struct {
	Description      string           `xml:"Description,attr"`  //ok
	ExternalCode     string           `xml:"ExternalCode,attr"` //ok
	Id               string           `xml:"Id,attr"`           //ok
	IdFromSource     string           `xml:"IdFromSource,attr"` //ok
	Index            string           `xml:"Index,attr"`        //ok
	Name             string           `xml:"Name,attr"`         //ok
	Tag              string           `xml:"Tag,attr"`          //ok
	ObjectReferences ObjectReferences `xml:"ObjectReferences"`  //ok
}

type Workspaces struct {
	Workspace []Workspace `xml:"Workspace"` //ok
}
