package main

import (
	"encoding/json"
	"github.com/hoisie/web"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Token string
}

type FigmaAnswer struct {
	Err    *string           `json:"err"`
	Images map[string]string `json:"images"`
	Status int               `json:"status"`
}

var Config Config

func GetFigma(ctx *web.Context, val string) {
	if (ctx.Params["file"] == "") || (ctx.Params["id"] == "") {
		ctx.Redirect(302, "http://999stories.com/wp-content/uploads/2017/09/Grustnyj-kot-Luhu-23.jpg")
		return
	}
	
	file := ctx.Params["file"]
	id := ctx.Params["id"]

	client := &http.Client{}
	geturl := "https://api.figma.com/v1/images/" + file + "?ids=" + id
	req, _ := http.NewRequest("GET", geturl, nil)
	req.Header.Set("X-FIGMA-TOKEN", Config.Token)
	res, _ := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}
	figma := new(FigmaAnswer)

	err2 := json.Unmarshal(body, figma)
	if err2 != nil {
		log.Fatal(err2)
	}

	if figma.Err != nil {
		ctx.Redirect(302, "http://999stories.com/wp-content/uploads/2017/09/Grustnyj-kot-Luhu-23.jpg")
	} else {
		ctx.Redirect(302, figma.Images[id])
	}
	return
}

func main() {
	tmp, filerr := os.Open("config.json")
	if filerr != nil {
		log.Fatal("No config file")
	}
	defer tmp.Close()
	decoder := json.NewDecoder(tmp)
	decoderr := decoder.Decode(&Config)
	if decoderr != nil {
		log.Fatal(decoderr)
	}

	web.Get("/figma(.*)", GetFigma)
	web.Run("0.0.0.0:5001")
}
