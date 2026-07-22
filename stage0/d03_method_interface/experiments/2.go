package experiments

import "fmt"

type Saver interface {
	Save() string
}

type Document struct {
	Name string
}

type Image struct {
	Name string
}

func (d *Document) Save() string {
	return "saved: " + d.Name
}

func (i Image) Save() string {
	return "saved image:" + i.Name
}

func printSaved(s Saver) {
	fmt.Println(s.Save())
}

func ExperimentMethodSet() {
	// 初始实验document
	document := Document{Name: "report"}
	// printSaved(document)
	printSaved(&document)

	// 增加类型 Image
	image := Image{Name: "photo.png"}
	printSaved(image)
	printSaved(&image)
}
