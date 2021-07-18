package jira

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type LabelService struct{ client *Client }

type IssueLabelsScheme struct {
	MaxResults int      `json:"maxResults"`
	StartAt    int      `json:"startAt"`
	Total      int      `json:"total"`
	IsLast     bool     `json:"isLast"`
	Values     []string `json:"values"`
}

// Gets returns a paginated list of labels.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/labels#get-all-labels
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-labels/#api-rest-api-3-label-get
func (l *LabelService) Gets(ctx context.Context, startAt, maxResults int) (result *IssueLabelsScheme,
	response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/3/label?%v", params.Encode())

	request, err := l.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = l.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
