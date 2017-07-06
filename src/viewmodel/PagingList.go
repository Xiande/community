package viewmodel

type PagingList struct {
	TotalRows   int         `json:"count"`
	ResultList  interface{} `json:"result"`
	CurLanguage string      `json:"language"`
	CurUserName string      `json:"userName"`
	IsAdmin     string      `json:"IsAdmin"`
}
