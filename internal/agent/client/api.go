package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/frauniki/ping-mesh/pkg/config"
	"github.com/frauniki/ping-mesh/pkg/domain"
	log "github.com/sirupsen/logrus"
)

func Post(req *domain.PushData) error {
	r, err := json.Marshal(req)
	if err != nil {
		return err
	}

	log.Infof("Push to %s", config.Config.ServerURL)
	resp, err := http.Post(config.Config.ServerURL, "application/json", bytes.NewBuffer(r))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Failed to POST %s (%d): %s", config.Config.ServerURL, resp.StatusCode, string(bodyBytes))
	}

	return nil
}
