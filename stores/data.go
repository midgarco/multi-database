package stores

import "time"

type Data struct {
	ConnId  int
	Results []map[string]interface{}
	Elapsed time.Duration
	Error   error

	// page data
	Start        int
	End          int
	CurrentToken string
	NextToken    string
}
