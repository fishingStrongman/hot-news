package bilibili

type Ret struct {
	Code    int
	Message string
	Ttl     int
	data    Data
}

type Data struct {
	Trending *Trend `json:"trending"`
}

type Trend struct {
	Title   string  `json:"title"`
	TrackId int     `json:"track_id"`
	List    []*List `json:"list"`
}

type List struct {
	Keyword  string `json:"keyword"`
	ShowName string `json:"show_name"`
	Icon     string `json:"icon"`
	Uri      string `json:"uri"`
	Goto     string `json:"goto"`
}
