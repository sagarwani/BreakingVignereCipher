package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
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

//Get the sequence of strings for the particular key "n".
func getsequence(ciphertext string, n int) []string{
	var harray []string

	for i := 0; i < len(ciphertext);{
		harray = append(harray, string(ciphertext[i]))
		i = i + n
	}
	return harray
}

func vdecrypt(fkey string, fnamearg string ) string{

	var alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//Decrypting the ciphertext into plaintext
	m:=0
	var plaintext string
	newkey := fkey
	for k := 0; k < len(fnamearg); k++ {
		newpos := pos(strings.ToUpper(string(fnamearg[k]))) - pos(strings.ToUpper(string(newkey[m])))
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
	return plaintext
}

func calculate_chisq(seq string) float64{
	var f map[string]float64
	f = make(map[string]float64)
	var h map[string]float64
	h = make(map[string]float64)
	var c map[string]int
	c = make(map[string]int)
	f["A"] = 8.17
	f["B"] = 1.49
	f["C"] = 2.78
	f["D"] = 4.25
	f["E"] = 12.70
	f["F"] = 2.23
	f["G"] = 2.02
	f["H"] = 6.09
	f["I"] = 7.00
	f["J"] = 0.15
	f["K"] = 0.77
	f["L"] = 4.03
	f["M"] = 2.41
	f["N"] = 6.75
	f["O"] = 7.51
	f["P"] = 1.93
	f["Q"] = 0.10
	f["R"] = 5.99
	f["S"] = 6.33
	f["T"] = 9.06
	f["U"] = 2.76
	f["V"] = 0.98
	f["W"] = 2.36
	f["X"] = 0.15
	f["Y"] = 1.97
	f["Z"] = 0.07

	//Find expected count of each alphabet and store it in the map
	var alphabetx string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for l := 0; l < len(alphabetx); l++ {
		expected_count := (f[string(alphabetx[l])] * float64(len(seq)))/100
		h[string(alphabetx[l])] = expected_count
	}

	//Find count of all alphabets in the given text
	var count int = 0
	for v := 0; v <len(alphabetx);v++{
		for w := 0; w < len(seq); w++{
			if seq[w] == alphabetx[v]{
				count = count + 1
			}
		}
		c[string(alphabetx[v])] = count
		count = 0
	}

	//Calculating the value of Chi-Squared statistic
	var chisq float64 = 0
	for t := 0; t < 26; t++{
		chisq = chisq + ((float64(c[string(alphabetx[t])]) - h[string(alphabetx[t])]) * (float64(c[string(alphabetx[t])]) - h[string(alphabetx[t])]))/(h[string(alphabetx[t])])
	}
	return chisq
}

func main () {

	//fmt.Println(" Vignere Cipher v1 Finding Key ")
	//fmt.Println(" ============================ ")
	//fmt.Println()

	//var keylength int
	//fmt.Scanf("%d", &keylength)

	// convert input (type string) to integer
	keylength, err := strconv.ParseInt(os.Args[1], 10, 0)

	if err != nil {
		fmt.Println("First input parameter must be integer")
		os.Exit(1)
	}
	//ciphertext := os.Args[2]

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
	ciphertext := string(data)
	fmt.Println()
	file.Close()

	//Refining the plaintext
	re := regexp.MustCompile("[a-zA-Z]")
	match := re.FindAllString(ciphertext,-1)
	ciphertext = strings.Join(match,"") //Joining string array to become a single string.
	ciphertext = strings.Replace(ciphertext, " ", "", -1) //Removing whitespaces if any.
	//fmt.Println("The key length to be used for decryption is: ",keylength)
	//fmt.Println("The Ciphertext to be decrypted: ",ciphertext)

	//==================
	var seq string
	var arrayseq []string

		n := int(keylength)
		for m := 0; m < n; m++ {
			for p := m; p < len(ciphertext);{
				seq =  seq + string(ciphertext[p])
				p = p + n
			}
			arrayseq = append(arrayseq, seq)
			seq = ""
		}

	//fmt.Println("The sequences of the keys are: ",arrayseq)

	var keyx []int
	for r := 0; r < int(keylength); r++ {

		keyseq := arrayseq[r]
		keyseq = strings.ToUpper(keyseq)
		//fmt.Println("Here is the sequence for the given key to be decrypted ", keyseq)

		//Creating Ceasar Ciphers for every alphabet.
		var alphabet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		var ceasarcipher []string
		var ceasar []string
		for j := 0; j < 26; j++ {
			seqlength := len(keyseq)
			for k := 0; k < seqlength; k++ {
				ceasarcipher = append(ceasarcipher, string(alphabet[j]))
			}
			ceasarcipherx := strings.Join(ceasarcipher, "")
			ceasar = append(ceasar, ceasarcipherx)
			ceasarcipher = nil
		}
		//fmt.Println("Here are the ceasar ciphers to be used: ", ceasar)
		var seqarray []string
		for x := 0; x < 26; x++ {
			seqarray = append(seqarray, vdecrypt(ceasar[x], keyseq))

		}
		//fmt.Println("Here is the plain text array after decryption: ", seqarray)

		//Calculating chi-sq for all the sequences in the above array
		var chisqarray []float64
		for d := 0; d < 26; d++ {
			chisqarray = append(chisqarray, calculate_chisq(seqarray[d]))
		}

		//fmt.Println("Here are the values of Chi-Squared Statistic: ", chisqarray)

		//Taking the index of lowest value in a variable slice
		min := chisqarray[0]
		var index int
		for e := 0; e < 26; e++{
			if chisqarray[e] < min{
				min = chisqarray[e]
				index = e
			}
		}


		//Creating a map of deciphered text and chi-squared statistic values
		var chisqmap map[string]float64
		chisqmap = make(map[string]float64)

		for y := 0; y < 26; y++ {
			chisqmap[seqarray[y]] = chisqarray[y]
		}
		//fmt.Println("The Chi Sequared statistic values are: ", chisqmap)

		keyx = append(keyx, index)
	}

	var alphabetx string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var keyq []string
	for q := 0; q < len(keyx); q++{
		keyq = append(keyq, string(alphabetx[keyx[q]]))
	}
	keyp := strings.Join(keyq,"")
	fmt.Println("Recovered Decipherment Key: ",keyp)

}
