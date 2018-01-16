func main() {
	xkcd, err := FetchCurrentXkcdComic()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n\n%s\n", xkcd.Title, xkcd.Hover)
	}
}
