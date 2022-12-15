package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"io"
	"net/http"
	"time"
)

var (
	httpClient = &http.Client{Timeout: 3 * time.Second}
	host       = "https://google.com"
)

type (
	Response      string
	ErrorResponse string
)

type Ping struct {
	response <-chan string
}

func NewPing() Ping {
	return Ping{response: watcher()}
}

func (p Ping) Cmd() tea.Msg {
	response, ok := <-p.response
	if !ok {
		return ErrorResponse("error, cannot make any more requests")
	}
	return Response(response)
}

// --- this is just an example of pushing to the command instead of pulling from the command ---

func watcher() <-chan string {
	out := make(chan string)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			resp, err := request()
			if err != nil {
				close(out)
				return
			}
			out <- resp
		}
	}()
	return out
}

func request() (string, error) {
	resp, err := httpClient.Get(host)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)
	return fmt.Sprintf("%s %s %s", time.Now().Format(time.StampMilli), host, resp.Status), nil
}
