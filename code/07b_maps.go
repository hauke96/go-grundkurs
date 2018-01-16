func main() {
	birth := map[string]int{
		"C":    1972,
		"Java": 1994,
		"go":  2007,
	}
	for language, year := range birth {
		fmt.Printf("%s was released in %d\n", language, year)
	}
}
