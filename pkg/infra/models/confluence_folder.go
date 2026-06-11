package models

// FolderDescendantChunkScheme represents a chunk of folder descendants in Confluence.
type FolderDescendantChunkScheme struct {
	Results []*FolderDescendantScheme         `json:"results,omitempty"` // The descendants in the chunk.
	Links   *FolderDescendantChunkLinksScheme `json:"_links,omitempty"`  // The links of the chunk.
}

// FolderDescendantChunkLinksScheme represents the links of a chunk of folder descendants in Confluence.
type FolderDescendantChunkLinksScheme struct {
	Next string `json:"next,omitempty"` // The link to the next chunk of descendants.
}

// FolderDescendantScheme represents a descendant of a folder in Confluence.
//
// Descendants are mixed content types (page, folder, whiteboard, database, embed).
type FolderDescendantScheme struct {
	ID            string `json:"id,omitempty"`            // The ID of the descendant.
	Type          string `json:"type,omitempty"`          // The content type (page, folder, whiteboard, database, embed).
	Status        string `json:"status,omitempty"`        // The status of the descendant.
	Title         string `json:"title,omitempty"`         // The title of the descendant.
	ParentID      string `json:"parentId,omitempty"`      // The ID of the parent of the descendant.
	ParentType    string `json:"parentType,omitempty"`    // The content type of the parent.
	Depth         int    `json:"depth,omitempty"`         // The depth of the descendant relative to the folder.
	ChildPosition int    `json:"childPosition,omitempty"` // The position of the descendant amongst its siblings.
}
