package server

import (
	"fmt"
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

func cleanUrl(in *string) error {
	parsed, err := url.ParseRequestURI(*in)
	if err != nil {
		return err
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return fmt.Errorf("Not a url")
	}
	q := parsed.Query()
	if len(q) == 0 {
		return fmt.Errorf("No query params")
	}
	q.Del("ref")
	// q.Del("from")

	// Cloudflare
	// These params will be added respectively
	// after users finish IUAM/JS Challenge and Captcha
	q.Del("__cf_chl_jschl_tk__")
	q.Del("__cf_chl_captcha_tk__")
	parsed.RawQuery = q.Encode()
	*in = parsed.String()
	return nil
}
