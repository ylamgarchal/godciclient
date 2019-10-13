package tests

import (
	"testing"

	"github.com/ylamgarchal/godciclient/dci"
)

func TestGetTopicByName(t *testing.T) {
	var dciAPI = dci.GetClient(
		"http://127.0.0.1:5000/api/v1",
		"admin",
		"admin")
	mytopic, err := dciAPI.GetTopicByName("RHEL-8.1")
	if err != nil {
		t.Error(err)
	}

	mytopic, err = dciAPI.GetTopicByName("non_existing_topic")
	if mytopic != nil && err == nil {
		t.Error(err)
	}
}
