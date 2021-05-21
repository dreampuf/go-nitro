package vpn

type Vpnvserver struct {
	Advancedepa              string `json:"advancedepa,omitempty"`
	Appflowlog               string `json:"appflowlog,omitempty"`
	Authentication           string `json:"authentication,omitempty"`
	Authnprofile             string `json:"authnprofile,omitempty"`
	Backupvserver            string `json:"backupvserver,omitempty"`
	Bindpoint                string `json:"bindpoint,omitempty"`
	Cachetype                string `json:"cachetype,omitempty"`
	Cachevserver             string `json:"cachevserver,omitempty"`
	Certkeynames             string `json:"certkeynames,omitempty"`
	Cginfrahomepageredirect  string `json:"cginfrahomepageredirect,omitempty"`
	Clttimeout               int    `json:"clttimeout,omitempty"`
	Comment                  string `json:"comment,omitempty"`
	Csvserver                string `json:"csvserver,omitempty"`
	Curaaausers              int    `json:"curaaausers,omitempty"`
	Curstate                 string `json:"curstate,omitempty"`
	Curtotalusers            int    `json:"curtotalusers,omitempty"`
	Deploymenttype           string `json:"deploymenttype,omitempty"`
	Devicecert               string `json:"devicecert,omitempty"`
	Disableprimaryondown     string `json:"disableprimaryondown,omitempty"`
	Domain                   string `json:"domain,omitempty"`
	Doublehop                string `json:"doublehop,omitempty"`
	Downstateflush           string `json:"downstateflush,omitempty"`
	Dtls                     string `json:"dtls,omitempty"`
	Epaprofileoptional       bool   `json:"epaprofileoptional,omitempty"`
	Failedlogintimeout       int    `json:"failedlogintimeout,omitempty"`
	Groupextraction          bool   `json:"groupextraction,omitempty"`
	Httpprofilename          string `json:"httpprofilename,omitempty"`
	Icaonly                  string `json:"icaonly,omitempty"`
	Icaproxysessionmigration string `json:"icaproxysessionmigration,omitempty"`
	Icmpvsrresponse          string `json:"icmpvsrresponse,omitempty"`
	Ip                       string `json:"ip,omitempty"`
	Ipset                    string `json:"ipset,omitempty"`
	Ipv46                    string `json:"ipv46,omitempty"`
	L2conn                   string `json:"l2conn,omitempty"`
	Linuxepapluginupgrade    string `json:"linuxepapluginupgrade,omitempty"`
	Listenpolicy             string `json:"listenpolicy,omitempty"`
	Listenpriority           int    `json:"listenpriority,omitempty"`
	Loginonce                string `json:"loginonce,omitempty"`
	Logoutonsmartcardremoval string `json:"logoutonsmartcardremoval,omitempty"`
	Macepapluginupgrade      string `json:"macepapluginupgrade,omitempty"`
	Map                      string `json:"map,omitempty"`
	Maxaaausers              int    `json:"maxaaausers,omitempty"`
	Maxloginattempts         int    `json:"maxloginattempts,omitempty"`
	Name                     string `json:"name,omitempty"`
	Netprofile               string `json:"netprofile,omitempty"`
	Newname                  string `json:"newname,omitempty"`
	Ngname                   string `json:"ngname,omitempty"`
	Nodefaultbindings        string `json:"nodefaultbindings,omitempty"`
	Pcoipvserverprofilename  string `json:"pcoipvserverprofilename,omitempty"`
	Port                     int    `json:"port,omitempty"`
	Precedence               string `json:"precedence,omitempty"`
	Range                    int    `json:"range,omitempty"`
	Rdpserverprofilename     string `json:"rdpserverprofilename,omitempty"`
	Redirect                 string `json:"redirect,omitempty"`
	Redirecturl              string `json:"redirecturl,omitempty"`
	Response                 string `json:"response,omitempty"`
	Rhistate                 string `json:"rhistate,omitempty"`
	Rule                     string `json:"rule,omitempty"`
	Samesite                 string `json:"samesite,omitempty"`
	Secondary                bool   `json:"secondary,omitempty"`
	Servicename              string `json:"servicename,omitempty"`
	Servicetype              string `json:"servicetype,omitempty"`
	Somethod                 string `json:"somethod,omitempty"`
	Sopersistence            string `json:"sopersistence,omitempty"`
	Sopersistencetimeout     int    `json:"sopersistencetimeout,omitempty"`
	Sothreshold              int    `json:"sothreshold,omitempty"`
	State                    string `json:"state,omitempty"`
	Status                   int    `json:"status,omitempty"`
	Tcpprofilename           string `json:"tcpprofilename,omitempty"`
	Type                     string `json:"type,omitempty"`
	Usemip                   string `json:"usemip,omitempty"`
	Userdomains              string `json:"userdomains,omitempty"`
	Value                    string `json:"value,omitempty"`
	Vserverfqdn              string `json:"vserverfqdn,omitempty"`
	Weight                   int    `json:"weight,omitempty"`
	Windowsepapluginupgrade  string `json:"windowsepapluginupgrade,omitempty"`
}
