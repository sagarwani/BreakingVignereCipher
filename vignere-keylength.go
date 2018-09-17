package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

var (
	fileInfo os.FileInfo
	err      error
)

//To fetch the starting position of the first array element of the key in sequence array.
func getstartpos(n int) int{
	var sum int
	for a := 2; a < n; a++{
		sum = sum + a
	}
	return sum
}

//Get the sequence of strings for the particular key "n".
func getsequence(xarray []string, n int) []string{
	pos := getstartpos(n)
	var harray []string
	if n == 2 {
		harray = append(harray, xarray[0])
		harray = append(harray, xarray[1])
	}
	if n == 3{
		harray = append(harray, xarray[2])
		harray = append(harray, xarray[3])
		harray = append(harray, xarray[4])
	}
	if n > 3{
		for b := pos; b < (pos + n);b++{
			harray = append(harray, xarray[b])
		}
	}
	return harray
}

//To return the count of a character in the sequence string
func getcount(a string, b string) int{
	var count int
	for i:=0;i<len(a);i++{
		if string(a[i]) == b{
			count = count + 1
		}
	}
	return count
}

func calculate_ic(a string) float64{
	var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//fmt.Println("Inside Sum")
	a = strings.ToUpper(a)
	//fmt.Println("The value of sequence is ", a)
	var n = float64(len(a))
	n = n * (n - 1)
	var sum float64 = 0
	for i:=0;i<26;i++{
		var count = float64(getcount(a,string(alphabet[i])))
		//fmt.Println("The value of count is ", count)
		if count != 0 {
			sum = sum + ((count * (count - 1))/n)
			//fmt.Println("The value of length of string is ", n)
			//fmt.Println("The value of countx is ", count)
			//fmt.Println("Inside sum iteration and sum is ", sum)
		}
	}
	return sum
}


func main() {

	//fmt.Println(" Vignere Cipher v1 Keylength Estimation ")
	//fmt.Println(" ============================ ")
	//fmt.Println()

	//Simply reading the text file and taking contents into a variable.
	fname := os.Args[1]
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
	ciphertext := string(data)
	fmt.Println()
	file.Close()

	//Refining the plaintext
	re := regexp.MustCompile("[a-zA-Z]")
	match := re.FindAllString(ciphertext,-1)
	ciphertext = strings.Join(match,"") //Joining string array to become a single string.
	ciphertext = strings.Replace(ciphertext, " ", "", -1) //Removing whitespaces if any.
	//fmt.Println("The Ciphertext to be decrypted: ",ciphertext)

	//var alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//vcipher := ciphertext

	max_keylength := 32
	min_keylength := 2

	var seq string
	var arrayseq []string

	for n := min_keylength; n < max_keylength; n++ {
		for m := 0; m < n; m++ {
			for p := m; p < len(ciphertext);{
				seq =  seq + string(ciphertext[p])
				p = p + n
			}
			arrayseq = append(arrayseq, seq)
			seq = ""
		}
	}
	//fmt.Println("The sequences of the keys are: ",arrayseq)
	//fmt.Println("The list of sequence for the given key is ", seqx)
	//fmt.Println("The I.C. for the given sequence is ", ic)

	//Calculating IC for each key.
	var m map[int]float64
	m = make(map[int]float64)
	for r := min_keylength; r < max_keylength; r++{
		var avg_ic float64 = 0
		for _, element := range getsequence(arrayseq, r){
			//fmt.Println("The value of element is ", element)
			avg_ic = avg_ic + calculate_ic(element)
		}
		m[r] = avg_ic/float64(r)
	}

	//================Print MAP orderly====================
	// To store the keys in slice in sorted order
	var keys []float64
	for _,k := range m {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	// To perform the opertion you want
	var sorted_map_values []float64
	for _, k := range keys {
		sorted_map_values =  append(sorted_map_values, k)
	}
	//fmt.Println(sorted_map_values)
	//======================================================

	//Extract 5 largest values.
	j := len(sorted_map_values)
	j1 := sorted_map_values[j-1]
	j2 := sorted_map_values[j-2]
	j3 := sorted_map_values[j-3]
	j4 := sorted_map_values[j-4]
	j5 := sorted_map_values[j-5]

	for g,h := range m {
		if h == j1{
			fmt.Println("The highest value", j1, " has keylength of ", g)
		}
		if h == j2{
			fmt.Println("The 2nd highest value ", j2, " has keylength of ", g)
		}
		if h == j3{
			fmt.Println("The 3rd highest value ", j3, " has keylength of ", g)
		}
		if h == j4{
			fmt.Println("The 4th highest value ", j4, " has keylength of ", g)
		}
		if h == j5{
			fmt.Println("The 5th highest value ", j5, " has keylength of ", g)
		}
	}

}