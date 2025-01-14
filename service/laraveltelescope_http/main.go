package main

import (
	"github.com/LeakIX/l9format"
)

type LaravelTelescopeHttpPlugin struct {
	l9format.ServicePluginBase
}

func New() l9format.WebPluginInterface {
	return LaravelTelescopeHttpPlugin{}
}

func (LaravelTelescopeHttpPlugin) GetVersion() (int, int, int) {
	return 0, 0, 1
}

var getGitConfigRequest = l9format.WebPluginRequest{
		Method: "GET",
		Path: "/telescope/requests",
		Headers: map[string]string{},
		Body:[]byte(""),
}

func (LaravelTelescopeHttpPlugin) GetRequests() []l9format.WebPluginRequest {
	return []l9format.WebPluginRequest{getGitConfigRequest}
}

func (LaravelTelescopeHttpPlugin) GetName() string {
	return "LaravelTelescopeHttpPlugin"
}

func (LaravelTelescopeHttpPlugin) GetStage() string {
	return "open"
}
func (plugin LaravelTelescopeHttpPlugin) Verify(request l9format.WebPluginRequest, response l9format.WebPluginResponse, event *l9format.L9Event, options map[string]string) ( hasLeak bool) {
	if !getGitConfigRequest.Equal(request) || response.Response.StatusCode != 200 || response.Document == nil {
		return  false
	}
	if response.Document.Find("title").Text() == "Telescope" {
		event.Summary = "Laravel Telescope enabled at " + event.Url()
		return true
	}
	return  false
}
