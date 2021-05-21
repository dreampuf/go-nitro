package authentication

type Authenticationoauthidpprofile struct {
	Attributes                 string `json:"attributes,omitempty"`
	Audience                   string `json:"audience,omitempty"`
	Clientid                   string `json:"clientid,omitempty"`
	Clientsecret               string `json:"clientsecret,omitempty"`
	Configservice              string `json:"configservice,omitempty"`
	Defaultauthenticationgroup string `json:"defaultauthenticationgroup,omitempty"`
	Encrypttoken               string `json:"encrypttoken,omitempty"`
	Issuer                     string `json:"issuer,omitempty"`
	Name                       string `json:"name,omitempty"`
	Oauthstatus                string `json:"oauthstatus,omitempty"`
	Redirecturl                string `json:"redirecturl,omitempty"`
	Refreshinterval            int    `json:"refreshinterval,omitempty"`
	Relyingpartymetadataurl    string `json:"relyingpartymetadataurl,omitempty"`
	Sendpassword               string `json:"sendpassword,omitempty"`
	Signaturealg               string `json:"signaturealg,omitempty"`
	Signatureservice           string `json:"signatureservice,omitempty"`
	Skewtime                   int    `json:"skewtime,omitempty"`
}
