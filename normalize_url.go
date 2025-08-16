package main


import (
	"net/url"
)
/*
normalize url function
 - parse url
	- does the parsing from net/url package (which we still have to import) handle the
	erroring for non url inputs?
 - check if crucial parts of the url match
I don't put the error handling here because I put it in the unit test? I think thats mostly right.
maybe one error for non url input look at unit test if it can handle the normalize url function
having two outputs, one for the normalized url one for the error
*/
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
