func main() {
	birth := map[string]int{
		"C":    1972,
		"Java": 1994,
	}
	birth["Go"] = 2007
	examine(birth, "Go")
	delete(birth, "Go")
	examine(birth, "Go")
}

func examine(birth map[string]int, language string) {
	if year, present := birth[language]; present {
		fmt.Printf("%s was released in %d\n", language, year)
	} else {
		fmt.Printf("Never heard of %s...\n", language)
	}
}
