package main

import "fmt"

//Memanggil fungsi goodAfternoon
//dari dalam good afternoon akan dilakukan print "selamat sore name1 dan name2"
func main() {
	goodAfternoon("adi", "anti")
	goodAfternoon("ado", "suci")

}

// TODO: answer here
func goodAfternoon(val1, val2 string) {
	fmt.Printf("halo %s dan %s \n", val1, val2)
}
