package main

import (
	"github.com/LeakIX/l9format"
	"gopkg.in/ini.v1"
)

type GitConfigHttpPlugin struct {
	l9format.ServicePluginBase
}

func New() l9format.WebPluginInterface {
	return GitConfigHttpPlugin{}
}

func (GitConfigHttpPlugin) GetVersion() (int, int, int) {
	return 0, 0, 1
}

var getGitConfigRequest = l9format.WebPluginRequest{
		Method: "GET",
		Path: "/.git/config",
		Headers: map[string]string{},
		Body:[]byte(""),
}

func (GitConfigHttpPlugin) GetRequests() []l9format.WebPluginRequest {
	return []l9format.WebPluginRequest{getGitConfigRequest}
}

func (GitConfigHttpPlugin) GetName() string {
	return "GitConfigPlugin"
}

func (GitConfigHttpPlugin) GetStage() string {
	return "open"
}
func (plugin GitConfigHttpPlugin) Verify(request l9format.WebPluginRequest, response l9format.WebPluginResponse, event *l9format.L9Event, options map[string]string) (leak l9format.L9LeakEvent, hasLeak bool) {
	if !getGitConfigRequest.Equal(request) || response.Response.StatusCode != 200 {
		return leak, false
	}
	gitConfig, err := ini.ShadowLoad(response.Body)
	if err != nil {
		return leak, false
	}
	if section := gitConfig.Section(`remote "origin"`); section.HasKey("url") {
		leak.Data += string(response.Body)
		leak.Dataset.Files++
		leak.Dataset.Size = int64(len(response.Body))
		return leak, true
	}
	return leak, false
}