package system

type Systemcmdpolicy struct {
	Action     string      `json:"action,omitempty"`
	Builtin    interface{} `json:"builtin,omitempty"`
	Cmdspec    string      `json:"cmdspec,omitempty"`
	Feature    string      `json:"feature,omitempty"`
	Policyname string      `json:"policyname,omitempty"`
}
