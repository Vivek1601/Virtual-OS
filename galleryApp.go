package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
	"fmt"
    "io/ioutil"
    "log"
	"strings"
	"fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/theme"
)

func showGalleryApp() {
	a := app.New()
	w := a.NewWindow("Hello")
    w.Resize(fyne.NewSize(800,600))
	root_src:= "C:\\Users\\hp\\Pictures\\Screenshots"

    files, err := ioutil.ReadDir(root_src)
    if err != nil {
        log.Fatal(err)
    }
    tabs := container.NewAppTabs()
	// var picsArr []string;

    for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() == false{
		extension:=	strings.SplitAfter(file.Name(), ".")[1];
		if extension == "png" || extension == "jpeg"{
            // picsArr = append(picsArr,root_src + "\\" + file.Name());
			image := canvas.NewImageFromFile(root_src + "\\" + file.Name())
		image.FillMode = canvas.ImageFillOriginal
		tabs.Append(container.NewTabItem(file.Name(), image))
		}
		}
    } 

	// tabs := container.NewAppTabs(
	// 	container.NewTabItem("Image", canvas.NewImageFromFile(picsArr[0])),
	// 	// container.NewTabItem("Image 2", canvas.NewImageFromFile(picsArr[1])),

	// )
    //  for i:=1;i<len(picsArr);i++{
	// 	image := canvas.NewImageFromFile(picsArr[i])
	// 	image.FillMode = canvas.ImageFillOriginal
	// 	// tabs.Append(container.NewTabItem("Image", canvas.NewImageFromFile(picsArr[i])))
	// 	tabs.Append(container.NewTabItem("Image", image))
	//  }
	 tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(tabs)
	w.Show()
}