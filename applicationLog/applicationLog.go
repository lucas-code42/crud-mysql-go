package applicationlog

import "fmt"

// LogHelloWeb
func LogHelloWeb() {
	fmt.Println("path execution: controller/controller.go -> HelloWeb")
}

// LogCreate
func LogCreate() {
	fmt.Println("path execution: controller/controller.go -> Create")
}

// LogDelete
func LogDelete() {
	fmt.Println("path execution: controller/controller.go -> Delete")
}

// LogListAllUsers
func LogListAllUsers() {
	fmt.Println("path execution: controller/controller.go -> LogListAllUsers")
}
