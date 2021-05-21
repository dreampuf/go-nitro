package tm

type Tmsessionparameter struct {
	Defaultauthorizationaction string `json:"defaultauthorizationaction,omitempty"`
	Homepage                   string `json:"homepage,omitempty"`
	Httponlycookie             string `json:"httponlycookie,omitempty"`
	Kcdaccount                 string `json:"kcdaccount,omitempty"`
	Name                       string `json:"name,omitempty"`
	Persistentcookie           string `json:"persistentcookie,omitempty"`
	Persistentcookievalidity   int    `json:"persistentcookievalidity,omitempty"`
	Sesstimeout                int    `json:"sesstimeout,omitempty"`
	Sso                        string `json:"sso,omitempty"`
	Ssocredential              string `json:"ssocredential,omitempty"`
	Ssodomain                  string `json:"ssodomain,omitempty"`
	Tmsessionpolicybindtype    string `json:"tmsessionpolicybindtype,omitempty"`
	Tmsessionpolicycount       int    `json:"tmsessionpolicycount,omitempty"`
}
