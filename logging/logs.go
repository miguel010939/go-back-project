package logging

import (
	"fmt"
	"net/http"
	"time"
)

func Log(r *http.Request, statusCode int) {
	preciseTime := fmt.Sprintf("%v/%v/%v %v:%v ", time.Now().Day(), time.Now().Month(), time.Now().Year(), time.Now().Hour(), time.Now().Minute())
	fmt.Printf("[%s] %s -> \"%s\"  status: %d\n", preciseTime, r.Method, r.URL.Path, statusCode)
}
