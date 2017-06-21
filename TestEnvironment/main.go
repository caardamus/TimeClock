package main

import ("fmt"
	"os"
	"bufio"
	"time"

)

func main() {

	var myname string
	now := time.Now()

	str1 := "Please enter your WVNCC username:"
	fmt.Println(str1)
	fmt.Scanf("%s", &myname) //user enters username
	fmt.Println("User clock-in: ", myname)
	fmt.Println("Clock in: ",now.Format(time.ANSIC)) //current time stamp

	fileHandle, _ := os.OpenFile(myname +".txt", os.O_CREATE|os.O_APPEND, 0755) //creates and appends txt file
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, myname, ":", now.Format(time.ANSIC), "\r\n") //appends txt file with username and timestamp
	writer.Flush()
}