package applicationlog

import "fmt"

func LogHelloWeb() {
	fmt.Println("path execution: controller/controller.go -> HelloWeb")
}

func LogCreate() {
	fmt.Println("path execution: controller/controller.go -> Create")
}