
func FetchInfo(site string, infos chan *Info, errs chan error) {
	url := "https://api.stackexchange.com/2.2/info?site=" + site
	response, err := http.Get(url)
	if err != nil {
		errs <- err
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		errs <- errors.New(response.Status)
		return
	}
	var info Info
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&info); err != nil {
		errs <- err
		return
	}
	info.Site = site
	infos <- &info
}
