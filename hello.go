package module

import "fmt"

func HelloWorld(name string) string {
	fmt.Println("Hello world ", name)
	return VERSION
}
