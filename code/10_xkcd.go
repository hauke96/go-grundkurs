type Xkcd struct { // full spec at https://xkcd.com/json.html
	Title string
	Hover string `json:"alt"`
}

func FetchCurrentXkcdComic() (*Xkcd, error) {
	response, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}
	var xkcd Xkcd
	decoder := json.NewDecoder(response.Body)
	if err := decoder.Decode(&xkcd); err != nil {
		return nil, err
	}
	return &xkcd, nil
}
