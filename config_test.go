package main

import (
	"github.com/stretchr/testify/assert"

	"io/ioutil"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	tempFile, _ := ioutil.TempFile(os.TempDir(), "test-")
	defer os.RemoveAll(tempFile.Name())

	JSONConfig := []byte(`{"server": {"address": "127.0.0.1", "port": 32400, "token": "webtoken"}}`)

	tempFile.Write(JSONConfig)

	config := Load(tempFile.Name())

	assert.Equal(t, config.Server.Address, "127.0.0.1")
	assert.Equal(t, config.Server.Port, 32400)
	assert.Equal(t, config.Server.Token, "webtoken")
}
