package accumulate

const testVersion = 1

type operation func(string) string

//Accumulate performs a given operation on each element of a given slice of strings
func Accumulate(collection []string, op operation) []string {
	var result []string

	for _, item := range collection {
		result = append(result, op(item))
	}

	return result
}
