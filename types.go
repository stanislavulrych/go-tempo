package tempo

type TempoResponse struct {
	Metadata TempoMetadata `json:"metadata"`
	// Results  []interface{} `json:"results"`
	Self string `json:"self"`
}

type TempoMetadata struct {
	Count    int    `json:"count"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
}

type TempoWorklogAttributeValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TempoWorklogAttributes struct {
	Self   string                       `json:"self"`
	Values []TempoWorklogAttributeValue `json:"values"`
}

type TempoWorklogAuthor struct {
	AccountId string `json:"accountId"`
	Self      string `json:"self"`
}

type TempoWorklogIssue struct {
	Id   int    `json:"id"`
	Self string `json:"self"`
}

type TempoWorklog struct {
	Attributes       TempoWorklogAttributes `json:"attributes"`
	Author           TempoWorklogAuthor     `json:"author"`
	BillableSeconds  int                    `json:"billableSeconds"`
	CreateAt         string                 `json:"createdAt"`
	Description      string                 `json:"description"`
	Issue            TempoWorklogIssue      `json:"issue"`
	Self             string                 `json:"self"`
	StartDate        string                 `json:"startDate"`
	StartTime        string                 `json:"startTime"`
	TempoWorklogId   int                    `json:"tempoWorklogId"`
	TimeSpentSeconds int                    `json:"timeSpentSeconds"`
	UpdateAt         string                 `json:"updatedAt"`
}

type TempoWorklogs = []TempoWorklog

type TempoResponseWorklogs struct {
	Metadata TempoMetadata `json:"metadata"`
	Results  TempoWorklogs `json:"results"`
	Self     string        `json:"self"`
}

type TempoAccount struct {
	Id  int    `json:"id"`
	Key string `json:"key"`
}

type TempoAccounts = []TempoAccount

type TempoResponseAccounts struct {
	Metadata TempoMetadata `json:"metadata"`
	Results  TempoAccounts `json:"results"`
	Self     string        `json:"self"`
}

type TempoHolidayScheme struct {
	DefaultScheme bool   `json:"defaultScheme"`
	Description   string `json:"description"`
	Id            int    `json:"id"`
	MemberCount   int    `json:"memberCount"`
	Name          string `json:"name"`
	Self          string `json:"self"`
}

type TempoHolidaySchemes = []TempoHolidayScheme

type TempoResponseHolidaySchemes struct {
	Metadata TempoMetadata       `json:"metadata"`
	Results  TempoHolidaySchemes `json:"results"`
	Self     string              `json:"self"`
}

////
