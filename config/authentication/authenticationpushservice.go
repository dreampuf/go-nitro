package authentication

type Authenticationpushservice struct {
	Certendpoint          string `json:"certendpoint,omitempty"`
	Clientid              string `json:"clientid,omitempty"`
	Clientsecret          string `json:"clientsecret,omitempty"`
	Customerid            string `json:"customerid,omitempty"`
	Hubname               string `json:"hubname,omitempty"`
	Name                  string `json:"name,omitempty"`
	Namespace             string `json:"Namespace,omitempty"`
	Pushcloudserverstatus string `json:"pushcloudserverstatus,omitempty"`
	Pushservicestatus     string `json:"pushservicestatus,omitempty"`
	Refreshinterval       int    `json:"refreshinterval,omitempty"`
	Servicekey            string `json:"servicekey,omitempty"`
	Servicekeyname        string `json:"servicekeyname,omitempty"`
	Signingkey            string `json:"signingkey,omitempty"`
	Signingkeyname        string `json:"signingkeyname,omitempty"`
	Trustservice          string `json:"trustservice,omitempty"`
}
