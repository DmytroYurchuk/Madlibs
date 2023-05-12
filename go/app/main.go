package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/valyala/fasthttp"
	"net/http"
	"sync"
)

var (
	httpClient = &fasthttp.Client{}
	baseUrl    = "https://reminiscent-steady-albertosaurus.glitch.me/"
	fetchWord  = fetchWordFunc
)

func fetchWordFunc(part string) (string, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	url := baseUrl + part
	req.SetRequestURI(url)

	err := httpClient.Do(req, resp)
	if err != nil {
		return "", err
	}

	var word string
	err = json.Unmarshal(resp.Body(), &word)
	if err != nil {
		return "", err
	}

	return word, nil
}

func madlibHandler(c *gin.Context) {
	var wg sync.WaitGroup
	wg.Add(3)

	var adjective, verb, noun string
	var errAdjective, errVerb, errNoun error

	go func() {
		adjective, errAdjective = fetchWord("adjective")
		wg.Done()
	}()

	go func() {
		verb, errVerb = fetchWord("verb")
		wg.Done()
	}()

	go func() {
		noun, errNoun = fetchWord("noun")
		wg.Done()
	}()

	wg.Wait()

	if errAdjective != nil || errVerb != nil || errNoun != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch word(s)",
		})
		return
	}

	c.JSON(http.StatusOK, "It was a "+adjective+" day. I went downstairs to see if I could "+verb+" dinner. I asked, 'Does the stew need fresh "+noun+"?'")
}

func main() {
	router := gin.Default()
	router.GET("/madlib", madlibHandler)
	router.Run(":8080")
}
