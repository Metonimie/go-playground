// This is a simple Go program that makes use of the LCS algorithm to find how much a file plagiarized another file.
// I don't have 'real' knowledge of the Go programming language. I implemented this in Python first and I saw that
// it ran very slow and then I decided to use Go for the first time. I haven't read any book on Go and some things
// may not be ok.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path"
	"path/filepath"
	"strings"
	"unicode"
)

var showLcs bool

// FilterString removes all characters that are not spaces nor letters
// from the given string and returns an array of words.
func FilterString(str *string) []string {
	var buffer bytes.Buffer

	for _, c := range *str {
		if unicode.IsLetter(c) || unicode.IsSpace(c) {
			buffer.WriteRune(c)
		}
	}

	return strings.Fields(buffer.String())
}

// This is a memoized version of the LCS. Will be used for large files.
// It works the same way as Lcs does but it only returns the length of the lcs.
// It is okayish compared to the Lcs which didn't even work on Set3 and ate 61GB of memory in 60s,
// it will run in O(m*n) time.
func LcsMemory(x *string, y *string) ([]string, int, int) {
	var x_words = FilterString(x)
	var y_words = FilterString(y)

	var m = len(x_words)
	var n = len(y_words)

	var A = make([]int, n+1)
	var B = make([]int, n+1)

	var placeholder []string

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if x_words[i] == y_words[j] {
				A[j] = 1 + B[j+1]
			} else {
				B[j] = int(math.Max(float64(B[j]), float64(A[j+1])))
			}
			B = A
		}
	}

	// The placeholder string is used to have the same signature as Lcs.
	return placeholder, A[0], len(x_words)
}

// Lcs calculates the longest common substring from input X to input Y and returns it.
// Lcs takes in consideration only letters.
// It also returns the len of the string
func Lcs(x *string, y *string) ([]string, int, int) {
	// Make an array of words and clean them
	var x_words = FilterString(x) // Has type of []string
	var y_words = FilterString(y)

	var m = len(x_words)
	var n = len(y_words)

	var c = make([][]int, m+1)
	for i := range c {
		c[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if x_words[i-1] == y_words[j-1] {
				c[i][j] = c[i-1][j-1] + 1
			} else {
				c[i][j] = int(math.Max(float64(c[i-1][j]), float64(c[i][j-1])))
			}
		}
	}

	var lcs_string []string
	var i = m
	var j = n
	for i > 0 && j > 0 {
		if x_words[i-1] == y_words[j-1] {
			lcs_string = append(lcs_string, string(x_words[i-1]))
			i -= 1
			j -= 1
		} else if c[i-1][j] > c[i][j-1] {
			i -= 1
		} else {
			j -= 1
		}
	}

	// Reverse
	var lcsStringLen = len(lcs_string)
	for i := 0; i < lcsStringLen/2; i++ {
		lcs_string[i], lcs_string[lcsStringLen-1-i] = lcs_string[lcsStringLen-1-i], lcs_string[i]
	}

	return lcs_string, c[m][n], len(lcs_string)
}

// LcsPair takes two file names and passes their data to the Lcs function
// It returns the lcs and the len of the lcs
func LcsPair(original string, test string) {
	data, err := ioutil.ReadFile(original)
	if err != nil {
		panic(err)
	}

	dataTwo, errTwo := ioutil.ReadFile(test)
	if errTwo != nil {
		panic(errTwo)
	}

	dataString := string(data)
	dataStringTwo := string(dataTwo)

	var (
		lcsString      []string
		length         int
		originalLength int
	)

	if !showLcs {
		lcsString, length, originalLength = LcsMemory(&dataString, &dataStringTwo)
	} else {
		lcsString, length, originalLength = Lcs(&dataString, &dataStringTwo)
	}

	GenerateReport(&lcsString, length, originalLength, original, test)
}

// Takes a directory path and computes LCS for every .txt
// file in the root.
func LcsTabular(myPath string) {
	files, err := ioutil.ReadDir(myPath)
	if err != nil {
		panic(err)
	}

	var validFiles []string

	// Validate all the files
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			validFiles = append(validFiles, path.Join(myPath, file.Name()))
		}
	}

	// Compute LCS for each file only once, don't do it vice-versa
	var validFilesLength = len(validFiles)
	for i := 0; i < validFilesLength; i++ {
		for j := i + 1; j < validFilesLength; j++ {
			// TODO: Add concurrency
			LcsPair(validFiles[i], validFiles[j])
		}
	}
}

// Pretty prints a report.
func GenerateReport(data *[]string, lcsLength int, originalLength int, fileNameOne string, fileNameTwo string) {
	fmt.Println("Report for", fileNameOne, "-", fileNameTwo)

	var plagiarized = (lcsLength * 100) / originalLength
	fmt.Printf("Plagiarized ammount: %d%%\n", plagiarized)
	fmt.Println("LCS length:", lcsLength, "words")

	if showLcs {
		fmt.Println(*data)
	}
	fmt.Println()
}

// Shows flag usage and exits gracefully.
func showUsage() {
	flag.Usage()
	os.Exit(0)
}

func main() {
	var pairFlag bool
	var tabularFlag bool

	// Hopefully nobody will enable this.
	flag.BoolVar(&showLcs, "showlcs", false, "Show the Lcs string. Warning: This option requires."+
		" a lot of memory! Also, it must come before the other arguments.")

	flag.BoolVar(&pairFlag, "pair", false, "Compare two files and show the plagiarized ammount.")
	flag.BoolVar(&tabularFlag, "tab", false, "Compare all the .txt files from the specified path.")
	flag.Parse()

	if pairFlag {
		if flag.Arg(0) == "" || flag.Arg(1) == "" {
			showUsage()
		}

		LcsPair(flag.Arg(0), flag.Arg(1))
	} else if tabularFlag {
		if flag.Arg(0) == "" {
			showUsage()
		}

		LcsTabular(flag.Arg(0))
	} else {
		showUsage()
	}
}
