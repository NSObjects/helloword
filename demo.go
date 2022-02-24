// Package plugindemo a demo plugin.
package helloworddemo

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Config the plugin configuration.
type Config struct {
	Log bool `json:"log,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

// Demo a Demo plugin.
type Demo struct {
	next http.Handler
	name string
	log  bool
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {

	return &Demo{
		next: next,
		name: name,
		log:  config.Log,
	}, nil
}

func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	start := time.Now()
	a.next.ServeHTTP(rw, req)
	cost := time.Since(start)
	if a.log {
		fmt.Println("请求花费时间：", cost)
	}
	a.next.ServeHTTP(rw, req)
}
