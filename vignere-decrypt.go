package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
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

	fmt.Println(" Vignere Cipher v1 Decryption ")
	fmt.Println(" ============================ ")
	fmt.Println()

	var alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var newkey string

	//Simply reading the text file and taking contents into a variable.
	fname := os.Args[2]
	file, err := os.Open(string(fname))
	if err != nil {
		log.Fatal(err)
	}

	//Check if file size if less than 100KB.
	fileInfo, err := os.Stat(fname)
	if fileInfo.Size() < 100000 {
		fmt.Println("Size of the file is {", fileInfo.Size(), "bytes } which is less than 100KB. Proceeding. . .")
	} else {
		fmt.Println("The file size is greater than 100KB. Please try again")
		os.Exit(3)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Contents of the file are: ", string(data))
	fmt.Println()
	file.Close()

	key := os.Args[1]
	ciphertextfile := string(data)

	//Refining the plaintext
	re := regexp.MustCompile("[a-zA-Z]")
	match := re.FindAllString(ciphertextfile,-1)
	ciphertextfile = strings.Join(match,"") //Joining string array to become a single string.
	ciphertextfile = strings.Replace(ciphertextfile, " ", "", -1) //Removing whitespaces if any.
	fmt.Println("Here is the extracted text from file: ",match)
	fmt.Println("The key to be used for decryption is: ",key)
	fmt.Println("The Ciphertext to be decrypted: ",ciphertextfile)

	//Make the size of key equal to plaintext size
	j := 0
	if len(ciphertextfile) > len(key) {
		for i := 0; i < len(ciphertextfile); i++ {
			if j == len(key){
				j = 0
			}
			newkey =  newkey + string(key[j])
			j++
		}
	}
	fmt.Println("The key to be used in the operation: ",newkey)

	//Decrypting the ciphertext into plaintext
	m:=0
	var plaintext string
	for k := 0; k < len(ciphertextfile); k++ {
		newpos := pos(strings.ToUpper(string(ciphertextfile[k]))) - pos(strings.ToUpper(string(newkey[m])))
		newpos = int(newpos)
		if newpos > 0 {
			newpos = newpos % 26
			plaintext = plaintext + string(alphabet[newpos])
		}else {
			newpos = (newpos + 26) % 26
			plaintext = plaintext + string(alphabet[newpos])
		}
		m++
	}
	fmt.Println("The Plaintext is ", plaintext)
}
