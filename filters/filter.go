// filter
package filters

import (
	"fmt"
	"net/http"
)

func AccessFilter(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, 100)

	r.Body.Read(body)
	fmt.Println(body)

}
