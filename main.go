package main

import (
	"fmt"

	gotrans "github.com/salihkemaloglu/localization-golang/services"
)

func main() {
	err := gotrans.InitLocales("langs")
	if err != nil {
		fmt.Println("Error happened when langs file loaded", err.Error())
	}

	lang := gotrans.DetectLanguage("en")
	fmt.Println("", gotrans.Tr(lang, "hello_world"))
	lang = gotrans.DetectLanguage("tr")
	fmt.Println("", gotrans.Tr(lang, "hello_world"))
	lang = gotrans.DetectLanguage("fr")
	fmt.Println("", gotrans.Tr(lang, "hello_world"))
}
