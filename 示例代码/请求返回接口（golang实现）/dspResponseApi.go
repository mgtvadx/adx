package api

type DSPResponse struct {
	Version int    `json:"version"`
	Bid     string `json:"bid"`
	ErrCode int    `json:"err_code"`
	Ads     []ADS  `json:"ads"`
}

type ADS struct {
	SpaceId         string   `json:"space_id"`
	Price            int      `json:"price"`
	AdUrl            string   `json:"ad_url"`
	CreativeId       string   `json:"creative_id"` //素材id
	ClickThroughUrl  string   `json:"click_through_url"`
	LandingOpenMode  int      `json:"land_open_mode"` //落地页打开方式，1：webview，0：默认
	IUrl             []IURL   `json:"iurl"`
	CUrl             []string `json:"curl"`
	AdId             string   `json:"adid"`
	Title            string   `json:"title"`
	Duration         int      `json:"duration"`
	Ctype            int      `json:"ctype"`
	Width            int      `json:"width"`
	Height           int      `json:"height"`
	Dealid           string   `json:"dealid"`
	WinNotice        string   `json:"win_notice"` //win_notice url
	AdDesc           string   `json:"ad_desc"`
	CreativeStyle    int      `json:"style"`
	ClickAction      int      `json:"click_action"`
	Advertiser       string   `json:"advertiser"`
}

type IURL struct {
	Event int //0-开始, 1-四分之一, 2-二分之一, 3-四分之三, 4-结束
	Url   string
}
