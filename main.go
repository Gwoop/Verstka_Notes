package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"

	"fyne.io/fyne/v2/widget"
)

func main() {

	//Server.Getdata()

	a := app.New()
	w := a.NewWindow("Notepad")
	w.Resize(fyne.NewSize(300, 500))

	title := widget.NewLabel("Ежедневник")

	//name_label := widget.NewLabel("Название дела")
	//name := widget.NewEntry()

	//case_description := widget.NewLabel("Описание дела")
	//description := widget.NewEntry()

	//date := widget.NewLabel("Дата")
	//date_r := widget.NewEntry()

	//reminder := widget.NewLabel("Напомнить?")

	formStruct := struct {
		Тема      string `json:"Theme"`
		Описание  string `json:"Discribe"`
		Дата      string `json:"Data"`
		Напомнить bool   `json:"remember"`
	}{}

	formData := binding.BindStruct(&formStruct)
	form := newFormWithData(formData)
	form.OnSubmit = func() {
		fmt.Println("Struct:\n", formStruct)
		naebka := formStruct.Напомнить
		if naebka == true {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Fyne Demo",
				Content: "Testing notifications...",
			})
		}

	}

	result := widget.NewLabel("")

	/*btn := widget.NewButton("Назначить", func() {
		username := name.Text
		description_r := description.Text
		male_val := date_r.Text
		//rem_arr := reminder_r.Selected
		result.SetText(username + "\n" + description_r + "\n" + male_val + "\n" + "\n")

	})*/

	w.SetContent(container.NewVBox(
		title,
		//name_label,
		//name,
		//case_description,
		//description,
		//date,
		//date_r,
		form,
		//reminder,
		//reminder_r,
		//btn,
		result,
	))

	w.ShowAndRun()
}

func newFormWithData(data binding.DataMap) *widget.Form {
	keys := data.Keys()
	items := make([]*widget.FormItem, len(keys))
	for i, k := range keys {
		data, err := data.GetItem(k)
		if err != nil {
			items[i] = widget.NewFormItem(k, widget.NewLabel(err.Error()))
		}
		items[i] = widget.NewFormItem(k, createBoundItem(data))
	}

	return widget.NewForm(items...)
}

func createBoundItem(v binding.DataItem) fyne.CanvasObject {
	switch val := v.(type) {
	case binding.Bool:
		return widget.NewCheckWithData("", val)
	case binding.Float:
		s := widget.NewSliderWithData(0, 1, val)
		s.Step = 0.01
		return s
	case binding.Int:
		return widget.NewEntryWithData(binding.IntToString(val))
	case binding.String:
		return widget.NewEntryWithData(val)
	default:
		return widget.NewLabel("")
	}
}
