package main

import (
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/elazarl/go-bindata-assetfs"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"strings"
	"os"
	"html/template"
)

var config *Config
//go:generate go-bindata -debug templates static
func main() {
	var err error
	config, err = ParseConfig("launcher.toml")
	if err != nil {
		fmt.Printf("Error while parsing launcher.toml config: %v\n", err.Error())
		os.Exit(1)
	}
	mux := pat.New()
	fs := http.FileServer(&assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, Prefix: "static"})
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	go h.run()
	http.Handle("/ws", websocket.Handler(wsHandler))

	http.Handle("/", mux)
	mux.Get("/", http.HandlerFunc(Home))
	mux.Get("/scripts/:name", http.HandlerFunc(ScriptHandler))
	log.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}

func GetAsset(path string) []byte {
	data, err := Asset(path)
	if err != nil {
		panic(err)
	}
	return data
}

func Home(w http.ResponseWriter, r *http.Request) {
	content := GetAsset("templates/homepage.html.tmpl")
	fmt.Fprint(w, string(content))
}

func ScriptHandler(w http.ResponseWriter, r *http.Request) {
	content := GetAsset("templates/script.html.tmpl")
	tmpl, err := template.New("script").Parse(string(content))
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	params := r.URL.Query()
	name := params.Get(":name")
	if config.Scripts[name] == nil {
		fmt.Fprint(w, "script " + name + " missing")
		return
	}
	templateParams := make(map[string]interface{})
	templateParams["name"] = name
	templateParams["script"] = config.Scripts[name]
	tmpl.Execute(w, templateParams)
}

func ExecuteScript(name string, send chan string) {
	var script *Script
	var host *Host
	for scriptName, configScript := range config.Scripts {
		if scriptName == name {
			script = configScript
		}
	}
	if script != nil {
		for hostName, configHost := range config.Hosts {
			if hostName == script.Host {
				host = configHost
			}
		}
	}
	if script != nil && host != nil {
		err := script.Execute(host, &ChannelWriter{send})
		if err != nil {
			panic(err)
		}
	}
}

type ChannelWriter struct {
	c chan string
}

func (self *ChannelWriter) Write(p []byte) (n int, err error) {
	message := string(p)
	message = strings.Replace(message, "\n", "<br/>", -1)
	self.c <- message
	return len(p), nil
}
