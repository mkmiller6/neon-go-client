package neon

type SearchContainer interface {
	GetSearchMeta() *SearchMeta
}

// SearchMeta is the structure that contains the common properties of the search iterators
type SearchMeta struct {
	HasMore  bool   `json:"has_more"`
	NextPage string `json:"next_page"`
	URL      string `json:"url"`
	// TotalCount is the total number of objects in the search result (beyond just
	// on the current page).
	TotalCount uint32 `json:"total_count"`
}

func (l *SearchMeta) GetSearchMeta() *SearchMeta {
	return l
}
