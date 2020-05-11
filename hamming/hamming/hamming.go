// Package hamming provides a utility to calculate the hamming distance between two DNA strands
package hamming

// Hamming calculates the Hamming distance between two strands of DNA
func Hamming(dna1 string, dna2 string) int {
	count := 0
	for i := 0; i < len(dna1); i++ {
		if dna1[i] != dna2[i] {
			count++
		}
	}

	return count
}
