package system

type Systemsession struct {
	All                   bool   `json:"all,omitempty"`
	Clienttype            string `json:"clienttype,omitempty"`
	Currentconn           bool   `json:"currentconn,omitempty"`
	Expirytime            int    `json:"expirytime,omitempty"`
	Lastactivitytime      string `json:"lastactivitytime,omitempty"`
	Lastactivitytimelocal string `json:"lastactivitytimelocal,omitempty"`
	Logintime             string `json:"logintime,omitempty"`
	Logintimelocal        string `json:"logintimelocal,omitempty"`
	Numofconnections      int    `json:"numofconnections,omitempty"`
	Partitionname         string `json:"partitionname,omitempty"`
	Sid                   int    `json:"sid,omitempty"`
	Username              string `json:"username,omitempty"`
}
