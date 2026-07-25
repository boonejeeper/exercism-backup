package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA []rune

// InvalidNucleotideError is an error type for invalid nucleotides.
type InvalidNucleotideError struct {
	Nucleotide rune
}

// Error implements the error interface for InvalidNucleotideError.
func (e InvalidNucleotideError) Error() string {
	return fmt.Sprintf("invalid nucleotide: %c", e.Nucleotide)
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
// /
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nucleotide := range d {
		if nucleotide != 'A' && nucleotide != 'C' && nucleotide != 'G' && nucleotide != 'T' {
			return nil, InvalidNucleotideError{nucleotide}
		}
		h[nucleotide]++
	}
	return h, nil
}
