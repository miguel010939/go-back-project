package repositories

import (
	"errors"
	"fmt"
)

// Something that can be considered to make sense or not. For example, -1 number of shoes is not null, but doesn't make
// sense... In the case of the number of shoes doesItMakeSense() could be nShoes>0 && nShoes%2==0
type optional interface {
	doesItMakeSense() bool
}

// A value that can "make sense", or not, paired with a query that will be used only if the associated value makes sense
type valueWithQuery struct {
	value optional
	query Query
}

// I wrap the array in a struct so i can define methods on it
type ArrayOfValuesAndQueries struct {
	vq []valueWithQuery
}

func NewArrayOfValuesAndQueries(valuesAndQueries ...valueWithQuery) *[]valueWithQuery {
	var array []valueWithQuery
	for _, valquery := range valuesAndQueries {
		array = append(array, valquery)
	}
	return &array
}

// Queries related to nonsensical values will be replaced with blank spaces
func (vqs *ArrayOfValuesAndQueries) filterNonSense() {
	for i, vq := range vqs.vq {
		if !vq.value.doesItMakeSense() {
			vqs.vq[i].query = " %v "
			vqs.vq[i].value = Query(" ") // by construction, the value must be a type optional
		}
	}
}

// inserts the values in the query, instead of using traditional placeholders
func insertValues(query string, values ...any) (string, error) {
	if len(values) == 2 {
		return fmt.Sprintf(query, values[0], values[1]), nil
	}
	if len(values) == 3 {
		return fmt.Sprintf(query, values[0], values[1], values[2]), nil
	}
	if len(values) == 4 {
		return fmt.Sprintf(query, values[0], values[1], values[2], values[3]), nil
	}
	return "", errors.New("number of args not supported")
}

// Go doesnt let me implement an interface for int, maybe because its a built-in type
type Number int

func (n Number) doesItMakeSense() bool {
	if n >= 0 {
		return true
	}
	return false
}

// Same as with Number: i can't implement an interface on a non-local type
type Query string

func (q Query) doesItMakeSense() bool {
	return q == ""
}
