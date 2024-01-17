package models

type ResponseItems struct {
	NextPage *int64   `json:"nextPage,omitempty"`
	PrevPage *int64   `json:"prevPage,omitempty"`
	PageInfo PageInfo `json:"pageInfo"`
	Items    []any    `json:"items"`
}

type PageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"limit"`
}

func MontarResponse(items []any, nextPage int64, prevPage int64, totalResults int64, resultsPerPage int64) *ResponseItems {
	return &ResponseItems{
		NextPage: &nextPage,
		PrevPage: &prevPage,
		Items:    items,
		PageInfo: PageInfo{
			TotalResults:   totalResults,
			ResultsPerPage: resultsPerPage,
		},
	}
}
