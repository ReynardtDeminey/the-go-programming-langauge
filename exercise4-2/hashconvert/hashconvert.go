// Package hashconvert provides a utility to return the 256/384/512 checksum of given data
package hashconvert

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

//HashConvert takes in a string as data and returns the checksum of that data
func HashConvert() string {
	hashFlag := flag.Int("h", 256, "Flag to select the type of hash to display of the input")
	wordFlag := flag.String("s", "Go", "Enter the word for which you want the hash value")
	flag.Parse()
	userHash := *hashFlag
	userWord := *wordFlag

	if userHash == 256 {
		hashedString := sha256.Sum256([]byte(userWord))
		return fmt.Sprintf("The %v-checksum of %v is %x\n", userHash, userWord, hashedString)
	} else if userHash == 384 {
		hashedString := sha512.Sum384([]byte(userWord))
		return fmt.Sprintf("The %v-checksum of %v is %x\n", userHash, userWord, hashedString)
	} else if userHash == 512 {
		hashedString := sha512.Sum512([]byte(userWord))
		return fmt.Sprintf("The %v-checksum of %v is %x\n", userHash, userWord, hashedString)
	}

	return "Cannot perform operation, not a recognized hash type"

}
