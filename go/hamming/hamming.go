package hamming

import "errors"

const testVersion = 6

//Distance calculates the Hamming distance of two equal length biological sequences
func Distance(sequence1, sequence2 string) (int, error) {
	sequence1Length := len(sequence1)

	if sequence1Length != len(sequence2) {
		return 0, errors.New("sequences must be of equal length")
	}

	difference := 0

	for i := 0; i < sequence1Length; i++ {
		if sequence1[i] != sequence2[i] {
			difference++
		}
	}

	return difference, nil
}
