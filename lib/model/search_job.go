package model

type searchJobService interface {
	NewSearch(spl string) (*SearchJob, error)

}
// SearchJob specifies the fields returned for a /search/jobs/ entry for a specific job
type SearchJob struct {
	Sid           string                 `json:"sid"`
	Content       map[string]interface{} `json:"content"`
	SearchService searchJobService
}

// NewSearchJob creates a SearchJob
func NewSearchJob(service searchJobService) SearchJob {
	return SearchJob{SearchService: service}
}