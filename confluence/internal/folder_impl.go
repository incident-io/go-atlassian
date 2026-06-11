package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/confluence"
	"net/http"
	"net/url"
	"strconv"
)

// NewFolderService returns a new Confluence V2 Folder service
func NewFolderService(client service.Connector) *FolderService {
	return &FolderService{internalClient: &internalFolderImpl{c: client}}
}

type FolderService struct {
	internalClient confluence.FolderConnector
}

// Descendants returns all descendants of a folder, at any depth.
//
// Descendants are mixed content types (page, folder, whiteboard, database, embed).
//
// The number of results is limited by the limit parameter and additional results (if available)
//
// will be available through the next cursor. A depth of 0 returns all levels.
//
// GET /wiki/api/v2/folders/{id}/descendants
//
// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-descendants/#api-folders-id-descendants-get
func (f *FolderService) Descendants(ctx context.Context, folderID int, depth int, cursor string, limit int) (*model.FolderDescendantChunkScheme, *model.ResponseScheme, error) {
	return f.internalClient.Descendants(ctx, folderID, depth, cursor, limit)
}

type internalFolderImpl struct {
	c service.Connector
}

func (i *internalFolderImpl) Descendants(ctx context.Context, folderID int, depth int, cursor string, limit int) (*model.FolderDescendantChunkScheme, *model.ResponseScheme, error) {

	if folderID == 0 {
		return nil, nil, model.ErrNoFolderIDError
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	if depth != 0 {
		query.Add("depth", strconv.Itoa(depth))
	}

	endpoint := fmt.Sprintf("wiki/api/v2/folders/%v/descendants?%v", folderID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	chunk := new(model.FolderDescendantChunkScheme)
	response, err := i.c.Call(request, chunk)
	if err != nil {
		return nil, response, err
	}

	return chunk, response, nil
}
