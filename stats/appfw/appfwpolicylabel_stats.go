package appfw

type Appfwpolicylabelstats struct {
	Appfwpolicy           interface{} `json:"appfwpolicy,omitempty"`
	Clearstats            string      `json:"clearstats,omitempty"`
	Labelname             string      `json:"labelname,omitempty"`
	Pipolicylabelhits     int         `json:"pipolicylabelhits,omitempty"`
	Pipolicylabelhitsrate int         `json:"pipolicylabelhitsrate,omitempty"`
}
