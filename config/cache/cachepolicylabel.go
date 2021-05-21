package cache

type Cachepolicylabel struct {
	Builtin                interface{} `json:"builtin,omitempty"`
	Evaluates              string      `json:"evaluates,omitempty"`
	Feature                string      `json:"feature,omitempty"`
	Flowtype               int         `json:"flowtype,omitempty"`
	Gotopriorityexpression string      `json:"gotopriorityexpression,omitempty"`
	Hits                   int         `json:"hits,omitempty"`
	Invokelabelname        string      `json:"invoke_labelname,omitempty"`
	Labelname              string      `json:"labelname,omitempty"`
	Labeltype              string      `json:"labeltype,omitempty"`
	Newname                string      `json:"newname,omitempty"`
	Numpol                 int         `json:"numpol,omitempty"`
	Priority               int         `json:"priority,omitempty"`
}
