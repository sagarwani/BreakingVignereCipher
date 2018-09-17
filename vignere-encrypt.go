package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	fileInfo os.FileInfo
	err      error
)

func pos(x string) int{
	var posreturn int
	var alphabetx string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for l := 0; l < len(alphabetx); l++ {
		if string(alphabetx[l]) == x{
			posreturn = l
		}
	}
	return posreturn
}

func main() {

	//fmt.Println(" Vignere Cipher v1 Encryption ")
	//fmt.Println(" ============================ ")
	//fmt.Println()

	key := os.Args[1]
	var alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var newkey string

	//Simply reading the text file and taking contents into a variable.
	fname := os.Args[2]
	file, err := os.Open(string(fname))
	if err != nil {
		log.Fatal(err)
	}

	//Check if file size if less than 100KB.
	fileInfo, err = os.Stat(fname)
	if fileInfo.Size() < 100000 {
		//fmt.Println("Size of the file is {", fileInfo.Size(), "bytes } which is less than 100KB. Proceeding. . .")
	} else {
		fmt.Println("The file size is greater than 100KB. Please try again")
		os.Exit(3)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Contents of the file are: ", string(data))
	textfile := string(data)
	fmt.Println()
	file.Close()

	//Refining the plaintext
	re := regexp.MustCompile("[a-zA-Z]")
	match := re.FindAllString(textfile,-1)
	textfile = strings.Join(match,"") //Joining string array to become a single string.
	textfile = strings.Replace(textfile, " ", "", -1) //Removing whitespaces if any.
	//fmt.Println("Here is the extracted text from file: ",match)
	//fmt.Println("The key to be used for encryption is: ",key)
	//fmt.Println("The Plaintext to be encrypted: ",textfile)

	//Make the size of key equal to plaintext size
	j := 0
	if len(textfile) > len(key) {
		for i := 0; i < len(textfile); i++ {
			if j == len(key){
				j = 0
			}
			newkey =  newkey + string(key[j])
			j++
		}
	}
	//fmt.Println("The key to be used in the operation: ",newkey)

	//Enciphering plaintext into ciphertext
	m:=0
	var ciphertext string
	for k := 0; k < len(textfile); k++ {
		newpos := (pos(strings.ToUpper(string(textfile[k]))) + pos(strings.ToUpper(string(newkey[m])))) % 26
		ciphertext = ciphertext + string(alphabet[int(newpos)])
		m++
	}
	fmt.Println("The Encrypted text is: ", ciphertext)
}
