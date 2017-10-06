package processors

import (
	"fmt"
	"net/http"
	"scheduler/types"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

// HTTPProcessor contains the protocol type and if the connection requires SSL
type HTTPProcessor struct {
	Name       string
	IsSSL      bool
	RunningLog *logrus.Logger
}

// Processing does a GET request on the specified URL and writes the response to the console
func (hp HTTPProcessor) Processing(sche types.Schedule) {
	colorRed := color.New(color.FgRed)
	colorYellow := color.New(color.FgYellow)
	colorGreen := color.New(color.FgGreen)

	fmt.Printf("%s Executing schedule %v....for %s....", time.Now().Format(time.StampMilli), sche, hp.Name)

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	url := sche.URL
	if hp.IsSSL {
		url = strings.Replace(url, "http://", "https://", 1)
	}

	responseText := "unknown response"
	if response, err := netClient.Get(url); err != nil {
		colorYellow.Println(responseText)
	} else if responseText = strconv.Itoa(response.StatusCode); response.StatusCode == 200 {
		colorGreen.Println(responseText)
	} else {
		colorRed.Println(responseText)
	}

	hp.RunningLog.WithFields(logrus.Fields{
		"Name":     hp.Name,
		"IsSSL":    hp.IsSSL,
		"Schedule": sche,
		"Response": responseText,
	}).Info("Executing Schedule")
}
