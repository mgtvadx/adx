package api

type DSPRequest struct {
	Version     int       `json:"version"`
	Bid         string    `json:"bid"`
	RequestType int       `json:"request_type"`
	Device      DSPDevice `json:"device"`
	Imp         []DSPIMP  `json:"imp"`
	Video       Video     `json:"video"`
}

type DSPDevice struct {
	Len               string `json:"len,omitempty"` //地理位置-经度
	Lon               string `json:"lon,omitempty"`
	Imei              string `json:"imei,omitempty"`
	Idfa              string `json:"idfa,omitempty"`
	Anid              string `json:"anid,omitempty"`
	Mac               string `json:"mac,omitempty"`
	Os                string `json:"os,omitempty"`
	Brand             string `json:"brand,omitempty"`
	Model             string `json:"model,omitempty"`
	Sw                int    `json:"sw,omitempty"` //屏幕分辨率宽度，单位为像素
	Sh                int    `json:"sh,omitempty"` // 屏幕分辨率高度，单位为像素
	Ip                string `json:"ip,omitempty"`
	CityCode          int    `json:"city_code,omitempty"`
	Ua                string `json:"ua,omitempty"`
	ConnectionType    int    `json:"connectiontype,omitempty"` //网络环境，0-wifi，1-mobile，2-no network
	Carrier           int    `json:"carrier"`        //运营商，0-未知，1-中国移动，2-中国联通，3-中国电信，4-互联网电视
	Type              int    `json:"type,omitempty"`           
	Version           string `json:"version,omitempty"`
	ScreenOrientation int    `json:"screen_orientation,omitempty"`
	Openudid          string `json:"openudid,omitempty"`
	Odin              string `json:"odin,omitempty"`
	Udid              string `json:"udid,omitempty"`
	Duid              string `json:"duid,omitempty"`
}

type DSPIMP struct {
	SpaceId     string      `json:"space_id"`
	Width       int         `json:"width"`
	Height      int         `json:"height"`
	MinCpmPrice int         `json:"min_cpm_price"`
	PlayerId    int         `json:"player_id"`
	Ctype       []int       `json:"ctype"`
	PlayTime    int         `json:"playtime"`
	MinPlayTime int         `json:"min_playtime"`
	MaxPlayTime int         `json:"max_playtime"`
	Location    int         `json:"location"`
	Order       int         `json:"order"`
	Pmp         []DspImpPmp `json:"pmp"`
}

type DspImpPmp struct {
	Id       string  `json:"id"`
	BidStyle int     `json:"at"`
	Price    float64 `json:"price"`
}

type Video struct {
	VideoId        int    `json:"video_id"`
	VideoName      string `json:"video_name,omitempty"`
	VideoUrl       string `json:"video_url,omitempty"`
	CollectionId   int    `json:"collection_id"`
	CollectionName string `json:"collection_name,omitempty"`
	ItemIds        string `json:"item_ids"`
	ItemNames      string `json:"item_names,omitempty"`
	AreaId         int    `json:"area_id"`
	AreaName       string `json:"area_name,omitempty"`
	ChannelId      int    `json:"channel_id"`
	Grades         string `json:"grades,omitempty"`
	Year           int    `json:"year"`
	Duration       int    `json:"duration"`
	Type           int    `json:"type"`
}
