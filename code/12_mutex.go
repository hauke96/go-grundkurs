var counter = 0
var mutex sync.Mutex

func root(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("%s\n\n", request)

	mutex.Lock()
	counter++
	count := counter
	mutex.Unlock()

	fmt.Fprintf(writer, "<html><body><pre>")
	fmt.Fprintln(writer, request.URL.Path)
	fmt.Fprintf(writer, "%d visits\n", count)
	fmt.Fprintf(writer, "</pre></body></html>")
}
