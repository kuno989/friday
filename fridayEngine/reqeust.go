package fridayEngine

import (
	"bytes"
	"encoding/json"
	"github.com/kuno989/friday/backend/schema"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func (s *Server) vmRequest(key, sha256, action string) {
	client := &http.Client{}
	client.Timeout = time.Second * 20
	res := schema.RequestMalware{
		ObjectKey: key,
		Sha256:    sha256,
	}
	buff, err := json.Marshal(res)
	if err != nil {
		logrus.Errorf("Failed to json marshall object: %v ", err)
	}
	uri := s.Config.AgentURI + s.Config.AgentPort + "/api/" + action
	body := bytes.NewBuffer(buff)
	req, err := http.NewRequest(http.MethodPost, uri, body)
	if err != nil {
		logrus.Fatalf("failed to request %v", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatalf("failed to request %v", err)
	}
	defer resp.Body.Close()
}

func (s *Server) updateDocument(sha256 string, buff []byte) {
	client := &http.Client{}
	client.Timeout = time.Second * 20
	uri := s.Config.URI + s.Config.WebserverPort + "/api/file/" + sha256
	body := bytes.NewBuffer(buff)
	req, err := http.NewRequest(http.MethodPut, uri, body)
	if err != nil {
		logrus.Fatalf("failed to request %v", err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatalf("failed to request %v", err)
	}
	defer resp.Body.Close()
	d, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("failed to read response %v", err)
	}
	logrus.Infof("status code: %d, response: %s", resp.StatusCode, string(d))
}
