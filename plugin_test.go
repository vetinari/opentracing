package main

import (
	"testing"

	"github.com/zalando/skipper/tracing"
)

var pluginDir string = "./build"

func TestLoadPluginBasic(t *testing.T) {
	if _, err := tracing.LoadPlugin(pluginDir, []string{"tracing_basic"}); err != nil {
		t.Errorf("failed to load plugin `basic`: %s", err)
	}
}

func TestLoadPluginInstana(t *testing.T) {
	if _, err := tracing.LoadPlugin(pluginDir, []string{"tracing_instana"}); err != nil {
		t.Errorf("failed to load plugin `instana`: %s", err)
	}
}

func TestLoadPluginLightstep(t *testing.T) {
	if _, err := tracing.LoadPlugin(pluginDir, []string{"tracing_lightstep", "token=123456", "collector=some.host.name:12345"}); err != nil {
		t.Errorf("failed to load plugin `lightstep`: %s", err)
	}
}
