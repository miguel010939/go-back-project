package handlers

import (
	"errors"
	"regexp"
	"strconv"
)

// ParseIntPathParam identifies a segment of the URL and returns the integer path param immediately after
func ParseIntPathParam(URL string, idName string) (int, error) {
	// ... foo/   {id}    /bar?x=3
	//    A         B         C
	var pathParamString string
	var re = regexp.MustCompile(idName)
	partAindeces := re.FindStringIndex(URL)
	if partAindeces == nil {
		return -1, errors.New("not present")
	}
	partBC := URL[partAindeces[1]:]
	partCindex := regexp.MustCompile("[/?]").FindStringIndex(partBC)
	if partCindex == nil {
		pathParamString = partBC
	} else {
		pathParamString = partBC[:partCindex[0]]
	}
	pathInt, err := strconv.Atoi(pathParamString)
	if err != nil {
		return -1, errors.New("path param format error: expected int")
	}
	return pathInt, nil
	// TODO i have to think if this can be done more efficiently...
}
