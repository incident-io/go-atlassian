package confluence

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// FolderConnector represents the Confluence Cloud Folders.
// Use it to navigate the contents of a folder.
type FolderConnector interface {

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
	Descendants(ctx context.Context, folderID int, depth int, cursor string, limit int) (*models.FolderDescendantChunkScheme, *models.ResponseScheme, error)
}
