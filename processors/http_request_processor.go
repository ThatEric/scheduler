package processors

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"scheduler/types"

	"github.com/fatih/color"
)

// HTTPProcessor contains the protocol type and if the connection requires SSL
type HTTPProcessor struct {
	Name  string
	IsSSL bool
}

// Processing does a GET request on the specified URL and writes the response to the console
func (hp HTTPProcessor) Processing(sche types.Schedule) {
	colorRed := color.New(color.FgRed)
	colorGreen := color.New(color.FgGreen)

	fmt.Printf("%s Executing schedule %v....for %s....", time.Now().Format(time.StampMilli), sche, hp.Name)

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	url := sche.URL
	if hp.IsSSL {
		url = strings.Replace(url, "http://", "https://", 1)
	}

	if response, err := netClient.Get(url); err != nil {
		colorRed.Println("unknown response")
	} else {
		colorGreen.Println(response.StatusCode)
	}
}
