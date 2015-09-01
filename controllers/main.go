package controllers

import (
	"fmt"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"github.com/mpgarate/JustRead/models"
	"net/http"
	"strings"
)

func HomeController(r render.Render, session sessions.Session) {
	r.HTML(200, "index", nil)
}

// build the beautified readable response
func ReadController(req *http.Request, r render.Render, session sessions.Session) {
	urlString := req.URL.Query().Get("url")

	if urlString == "" {
		r.Redirect("/", 300)
		return
	}

	article := readByUrl(urlString, req, r, session)
	if article != nil {
		r.HTML(200, "read", article)
	}
}

func readByUrl(
	urlString string,
	req *http.Request,
	r render.Render,
	session sessions.Session) *models.Article {

	article := models.Article{}
	article.Url = formatUrl(urlString)

	err := models.SetReadableContent(&article)
	if err != nil {
		fmt.Println(err)
		r.HTML(400, "error", "Could not retrieve article. ")
		return nil
	}

	return &article
}

func formatUrl(url string) string {
	if strings.Index(url, "http://") == 0 {
		return url
	} else if strings.Index(url, "https://") == 0 {
		return url
	} else {
		return "http://" + url
	}
}
