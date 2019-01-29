package enderlog

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type DefaultSlackRepository struct {
	slackUrl  string
	message   string
	debugable bool
}

type DefaultSlackPayload struct {
	Text string `json:"text"`
}

func NewDefaultSlackRepository(url string) LoggerRepository {
	return &DefaultSlackRepository{
		slackUrl:  url,
		debugable: false,
	}
}

func (s *DefaultSlackRepository) SetDebugable(debugable bool) {
	s.debugable = debugable
}

func (s *DefaultSlackRepository) Debugable() bool {
	return s.debugable
}

func (s *DefaultSlackRepository) Info(msg string) {
	s.message = msg
	s.send()
}

func (s *DefaultSlackRepository) Warn(msg string) {
	s.message = msg
	s.send()
}

func (s *DefaultSlackRepository) Error(msg string) {
	s.message = msg
	s.send()
}

func (s *DefaultSlackRepository) Fatal(msg string) {
	s.message = msg
	s.send()
}

func (s *DefaultSlackRepository) send() {
	p := &DefaultSlackPayload{}
	p.Text = s.message
	payload, err := json.Marshal(p)
	if err != nil {
		log.Println("ender-log error : create request payload failed")
		log.Println("ender-log error : " + err.Error())
	}

	req, err := http.NewRequest(http.MethodPost, s.slackUrl, bytes.NewReader(payload))
	if err != nil {
		log.Println("ender-log error : create request to slack failed")
		log.Println("ender-log error : " + err.Error())
	}

	client := http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("ender-log error : send request to slack failed")
		log.Println("ender-log error : " + err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println("ender-log error : send request to slack failed")
		log.Printf("ender-log error : status code %d", resp.StatusCode)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("ender-log error : " + err.Error())
		}
		log.Println("ender-log error : " + string(body))
		return
	}
}
