package server

import (
	"net/http"
	"net/url"

	"github.com/HelloLingC/moon-counter/common"
)

const JS = `fetch("//%s/counter/text").then(e=>e.text()).then(e=>{document.getElementById("moon-counter").innerText=e}).catch(e=>{console.error(e)});`

const JS_IMG = `fetch("//%s/counter/img").then(e=>e.text()).then(e=>{document.getElementById("moon-counter-img").src="data:image/svg+xml,"+e}).catch(e=>{console.error(e)});`

func checkOrigin(w http.ResponseWriter, origin string, hostnames []string) bool {
	parsed, err := url.Parse(origin)
	if err != nil {
		http.Error(w, "Invaild orgin: not a url", http.StatusBadRequest)
		return false
	}
	isAllowed := common.StrIsInSlice(parsed.Hostname(), hostnames)
	if !isAllowed {
		http.Error(w, "Invaild orgin: Forbidden", http.StatusForbidden)
	}
	return isAllowed
}
