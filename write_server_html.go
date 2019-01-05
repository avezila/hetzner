package main

import (
	"html/template"
	"os"
)

const WWWPath = "/usr/share/nginx/html/web/"
const WWWServerPath = WWWPath + "s/"

var ServerTemplate *template.Template

func init() {
	var err error
	ServerTemplate, err = template.ParseFiles("./server.go.html")
	if err != nil {
		panic(err)
	}
}

func (a AServer) WriteServerHTML() error {
	if err := os.MkdirAll(WWWServerPath, 0700); err != nil {
		return err
	}
	serverHTMLFile, err := os.OpenFile(WWWServerPath+a.Hash+".html", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer serverHTMLFile.Close()
	return ServerTemplate.Execute(serverHTMLFile, a)
}
