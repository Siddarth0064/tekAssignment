package main

import (
	"fmt"
	"os"
)

func main() {
	filename := []string{"a.txt", "b.txt"} // this is a slice string which contains file names
	for _, filename := range filename {
		if err := removeFile(filename); err != nil { // pass file name to removeFile func and checking err is occurd are not
			fmt.Printf("error while removeing filr %s: %v \n", filename, err)
		}
	}

}
func removeFile(filename string) error { // this is func which accepts file name as a string
	err := os.Remove(filename) // importing remove func and passing the file name
	if err != nil {            // checking err occured or not
		if os.IsNotExist(err) {
			fmt.Printf("file %s does not exist \n", filename)
		} else {
			return err
		}
	} else {
		fmt.Printf("file %s has been removed\n", filename)
	}
	return nil
}
