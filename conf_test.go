package goom

import (
	"testing"

	"github.com/spf13/viper"
)

func TestDefaultEnvConf(t *testing.T) {
	goomPath := viper.Get("goomPath")
	if goomPath == nil {
		t.Fatalf(`goomPath should be defined by default`)
	}

	configPath := viper.Get("configPath")
	if configPath == nil {
		t.Fatalf(`configPath should be defined by default`)
	}

	dataPath := viper.Get("dataPath")
	if dataPath == nil {
		t.Fatalf(`dataPath should be defined by default`)
	}

	boltdbPath := viper.Get("boltdbPath")
	if boltdbPath == nil {
		t.Fatalf(`boltdbPath should be defined by default`)
	}

}
