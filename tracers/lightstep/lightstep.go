package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"

	lightstep "github.com/lightstep/lightstep-tracer-go"
	opentracing "github.com/opentracing/opentracing-go"
)

func InitTracer(opts []string) (opentracing.Tracer, error) {
	var token string
	useThrift := false
	collector := lightstep.Endpoint{}
	for _, o := range opts {
		if strings.HasPrefix(o, "token=") {
			token = o[6:]
		}
		if o == "use_thrift" {
			useThrift = true
		}
		if strings.HasPrefix(o, "collector=") {
			host, portStr, err := net.SplitHostPort(o[10:])
			if err != nil {
				return nil, fmt.Errorf("failed to split host/port: %s", err)
			}
			port, err := strconv.ParseInt(portStr, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse %s as int: %s", portStr, err)
			}
			collector = lightstep.Endpoint{
				Host: host,
				Port: int(port),
			}
		}
	}
	if token == "" {
		return nil, errors.New("missing token= option")
	}

	lsOpts := lightstep.Options{
		AccessToken: token,
		UseThrift:   useThrift,
		Tags: map[string]interface{}{
			lightstep.ComponentNameKey: "skipper",
		},
	}
	if collector.Host != "" {
		lsOpts.Collector = collector
	}
	return lightstep.NewTracer(lsOpts), nil
}
