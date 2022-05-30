package main

import (
	"fmt"
	"log"
)

func main() {
	// panic example
	// a, b := 1,0
	// ans := a / b
	// fmt.Println(ans)

	//panic behavior
	// fmt.Println("start")
	// panic("something bad happened")
	// fmt.Println("end")

	// panic real world example
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	fmt.Println("start")
	panicker()
	fmt.Println("end")

}

// panic and recover example
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("somenthing bad happened")
	fmt.Println("done panicking")
}
