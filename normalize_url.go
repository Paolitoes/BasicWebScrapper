package main


import (
	"net/url"
)

func normalizeURL(URL string) (string,error){
	Url, err := url.Parse(URL)

	if err != nil {
		return "", err
	}
	
	path := Url.Path
	normURL := ""

	if rune(path[len(path)-1])=='/'{
		normURL = Url.Host + path[:len(path)-1]
	}else{
		normURL = Url.Host + path
	}

	return normURL, nil
}
