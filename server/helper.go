package server

import (
	"net/http"
	"net/url"

	"github.com/HelloLingC/moon-counter/common"
)

const JS = `
fetch('//%s/counter/text')
    .then(r => {
		return r.text();
    })
	.then(d => {
		document.getElementById("moon-counter").innerText = d;
	})
    .catch(e => {
        console.error(e);
    });
`

const JS_IMG = `
fetch('//%s/counter/img')
    .then(r => {
		return r.text();
    })
	.then(d => {
		document.getElementById("moon-counter").innerText = d;
	})
    .catch(e => {
        console.error(e);
    });
`

func checkOrigin(w http.ResponseWriter, origin string, hostnames []string) string {
	parsed, err := url.Parse(origin)
	if err != nil {
		http.Error(w, "Invaild orgin: not a url", http.StatusBadRequest)
		return ""
	}
	isAllowed := common.StrIsInSlice(parsed.Hostname(), hostnames)
	if !isAllowed {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return ""
	}
	rmOrigin, err := common.StrRemoveProtocol(origin)
	if err != nil {
		http.Error(w, "Invaild origin: missing protocal", http.StatusBadRequest)
		return ""
	}
	return rmOrigin
}
