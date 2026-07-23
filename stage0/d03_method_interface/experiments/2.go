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
	var _ Saver = (*Document)(nil)
	var _ Saver = Image{}
	var _ Saver = (*Image)(nil)

	// 下面这行取消注释后应编译失败：
	// var _ Saver = Document{}

	// 增加类型 Image
	image := Image{Name: "photo.png"}
	printSaved(image)
	printSaved(&image)
}
