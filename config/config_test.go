package config

import (
	"github.com/stretchr/testify/assert"

	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {

	tempFile, _ := ioutil.TempFile(os.TempDir(), "test-")
	defer os.RemoveAll(tempFile.Name())

	JSONConfig := []byte(`{"exporter": {"port": 9594}, "server": {"address": "127.0.0.1", "port": 32400, "token": "webtoken"}}`)

	tempFile.Write(JSONConfig)

	config, err := Load(tempFile.Name())

	assert.Nil(t, err)
	assert.Equal(t, config.Exporter.Port, "9594")
	assert.Equal(t, config.Server.Address, "127.0.0.1")
	assert.Equal(t, config.Server.Port, 32400)
	assert.Equal(t, config.Server.Token, "webtoken")
}

func TestLoadConfigFailsConfigFileDoesNotExists(t *testing.T) {

	file := "/tmp/config.json"
	_, err := Load(file)

	assert.NotNil(t, err)
	assert.Equal(t, err, fmt.Errorf("Config file %s does not exist", file))
}

func TestLoadConfigBadJsonFormat(t *testing.T) {

	tempFile, _ := ioutil.TempFile(os.TempDir(), "test-")
	defer os.RemoveAll(tempFile.Name())

	JSONConfig := []byte(`blabla`)

	tempFile.Write(JSONConfig)

	_, err := Load(tempFile.Name())

	assert.NotNil(t, err)
	assert.Equal(t, err, fmt.Errorf("Config file %s is not a valid json", tempFile.Name()))
}
