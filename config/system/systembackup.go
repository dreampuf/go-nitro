package system

type Systembackup struct {
	Comment          string `json:"comment,omitempty"`
	Createdby        string `json:"createdby,omitempty"`
	Creationtime     string `json:"creationtime,omitempty"`
	Filename         string `json:"filename,omitempty"`
	Includekernel    string `json:"includekernel,omitempty"`
	Ipaddress        string `json:"ipaddress,omitempty"`
	Level            string `json:"level,omitempty"`
	Size             int    `json:"size,omitempty"`
	Skipbackup       bool   `json:"skipbackup,omitempty"`
	Uselocaltimezone bool   `json:"uselocaltimezone,omitempty"`
	Version          string `json:"version,omitempty"`
}
