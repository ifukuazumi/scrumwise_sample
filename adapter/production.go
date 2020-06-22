package adapter

import (
	_ "context"
	"encoding/json"
	"fmt"
	"github.com/cenkalti/backoff"
	"github.com/ifukuazumi/scrumwise_sample/model"
	"github.com/ifukuazumi/scrumwise_sample/usecase/repository"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net/http"
)

const backoffRequest = 4

var url = "https://api.scrumwise.com/service/api/v1/getData"

// ScrumwiseService is
type Scrumwise struct {
	CredentialUserName string
	CredentialPassword string
	ProjectID string
	TagName string
}

// NewScrumwise is
func NewScrumwise(credentialUserName, credentialPassword, projectID, tagName string) repository.Production {
	return &Scrumwise{
		CredentialUserName: credentialUserName,
		CredentialPassword: credentialPassword,
		ProjectID:  projectID,
		TagName: tagName,
	}
}

// GetTags Tag情報を取得する
func (s *Scrumwise) GetTagID() (string, error) {
	byteArray, err := s.request(http.MethodGet, url, "Project.tags", nil)
	if err != nil {
		fmt.Println("byteArrayでエラーが起きたよ")
		return "", errors.WithStack(err)
	}
	var data *model.Response
	if err := json.Unmarshal(byteArray, &data); err != nil {
		return "", errors.WithStack(err)
	}

	for _, backlogs := range data.Result.Projects[0].Tags {
		if backlogs.Name == s.TagName {
			return backlogs.ID, nil
		}
	}
	return "", nil
}

// GetAll 全ての情報を取得する
func (s *Scrumwise) GetAll() (*model.Project, error) {
	byteArray, err := s.request(http.MethodGet, url, "Project.sprints,Project.backlogItems,BacklogItem.tasks,Project.tags", nil)
	if err != nil {
		fmt.Println("byteArrayでエラーが起きたよ")
		return nil, errors.WithStack(err)
	}

	var data *model.Response
	if err := json.Unmarshal(byteArray, &data); err != nil {
		return nil, errors.WithStack(err)
	}
	
	// 1Project単位でしか見ないため、[0]番目を取得する
	return &data.Result.Projects[0], nil
}

func (s *Scrumwise) request(method, url, includeProperties string, body io.Reader) ([]byte, error) {
	var byteArray []byte
	operation := func() error {
		var err error
		req, err := http.NewRequest(method, url, body)
		if err != nil {
			return errors.WithStack(err)
		}

		params := req.URL.Query()
		params.Add("projectIDs",s.ProjectID)
		params.Add("includeProperties",includeProperties)
		req.URL.RawQuery = params.Encode()

		req.SetBasicAuth(s.CredentialUserName, s.CredentialPassword)
		client := new(http.Client)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("respでエラーが起きたよ")
			return errors.WithStack(err)
		}
		defer resp.Body.Close()
		

		byteArray, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("ioutil.ReadAll(resp.Body)でエラーが起きたよ")
			return errors.WithStack(err)
		}

		if resp.StatusCode >= http.StatusMultipleChoices {
			fmt.Println("resp.StatusCode >= http.StatusMultipleChoices でエラー起きたよ")
			return errors.New(string(byteArray))
		}
		return nil
	}
	b := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), backoffRequest)
	if err := backoff.Retry(operation, b); err != nil {
		return nil, err
	}
	return byteArray, nil
}