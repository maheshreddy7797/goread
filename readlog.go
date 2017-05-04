package main
import (
    "fmt"
    "io/ioutil"
    "os"
    "path"
    "time"
)
//Reading files requires checking most calls for errors.
func check (e error) {
    if e != nil {
        panic(e)
    }
}
func main() {
//Get the current time and format into a string.
    t := time.Now()
    k :=t.Format("2006-01-02 15:04:05")
//Call the path() function to get the path to a file.
    filepath := dirpath()
//Open the file specified by the path.
    dat, err := ioutil.ReadFile(filepath)
//If the file doesn't exists raise an error.
    check(err)
//Print the contents of the file.
    fmt.Print(string(dat))
//Separate the directory path and filename from the input path.
//Eg: /home/cloudbyte/dir1/test.txt ---> /home/cloudbyte/dir1/ && test.txt.
    filepath= path.Dir(filepath)
//Make a filepath for logs.txt in same directoy.
    outputpath := filepath+"/logs.txt"
//Create a file named logs.txt in the same specified directory in APPEND mode.
//Eg: outputpath="/home/cloudbyte/dir1/logs.txt"
    f2, err := os.OpenFile(outputpath,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
//Raise an error if unable to create a file. 
    check(err)
//Close the file after creating.
    defer f2.Close()
//Format tag to append into logs.txt.
    tag := "Last opened on : " +k+ "\n"
//Append the string into the file.
    n1, err := f2.WriteString(tag)
    _ = n1
    check (err)
}
//Fetch the specified path as an argument from the user. 
func dirpath() string {
    arg :=""
    for i, a := range os.Args[1:] {
        i=i+1
	arg =a
    }
//Return the value of the path in string format.
    return arg
}
