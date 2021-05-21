package contentinspection

type Contentinspectionpolicylabel struct {
	Comment                string `json:"comment,omitempty"`
	Flowtype               int    `json:"flowtype,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Hits                   int    `json:"hits,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
	Isdefault              bool   `json:"isdefault,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Newname                string `json:"newname,omitempty"`
	Numpol                 int    `json:"numpol,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Type                   string `json:"type,omitempty"`
}
