func main() {
	sites := [...]string{
		"stackoverflow",
		"serverfault",
		"superuser"}

	infos := make(chan *Info)
	errs := make(chan error)

	for _, site := range sites {
		go FetchInfo(site, infos, errs)
	}

	for range sites {
		select {
		case info := <-infos:
			fmt.Printf("%#v\n\n", info)
		case err := <-errs:
			fmt.Printf("%s\n\n", err)
		}
	}
}
