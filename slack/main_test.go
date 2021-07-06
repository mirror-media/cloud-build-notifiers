package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/slack-go/slack"
	cbpb "google.golang.org/genproto/googleapis/devtools/cloudbuild/v1"
)

func TestWriteMessage(t *testing.T) {
	n := new(slackNotifier)
	b := &cbpb.Build{
		Name: 	   "my-project-name",
		ProjectId:        "project-id",
		Status:    cbpb.Build_SUCCESS,
        Images:    []string{"built-image"},
        Tags:      []string{"built-tag"},
		LogUrl:    "https://some.example.com/log/url?foo=bar",
        BuildTriggerId: "triger-Id",
		Source:		Source{
					RepoSource: RepoSource {
						RepoName: "reponame",
						BranchName: "BranchName",
						TagName: "TagName",
					}},
	}

	got, err := n.writeMessage(b)
	if err != nil {
		t.Fatalf("writeMessage failed: %v", err)
	}

	want := &slack.WebhookMessage{
		Attachments: []slack.Attachment{{
			Text:  "SUCCESS: project-id - built-image:built-tag",
			Color: "good",
			Actions: []slack.AttachmentAction{{
				Text: "Build Logs",
				Type: "button",
				URL:  "https://some.example.com/log/url?foo=bar&utm_campaign=google-cloud-build-notifiers&utm_medium=chat&utm_source=google-cloud-build",
			}},
		}},
	}

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("writeMessage got unexpected diff: %s", diff)
	}
}
