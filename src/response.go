package src

/**
返回结果
*/
type result struct {
	Ok     int         `json:"ok"`
	ErrNo  int         `json:"err_no"`
	Failed interface{} `json:"failed"`
	Data   string      `json:"data"`
}

type status struct {
	Status int    `json:"status"`
	Desc   string `json:"desc"`
}
