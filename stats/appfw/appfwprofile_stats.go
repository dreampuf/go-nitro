package appfw

type Appfwprofilestats struct {
	Appfirewallabortsperprofile                            int    `json:"appfirewallabortsperprofile,omitempty"`
	Appfirewallabortsperprofilerate                        int    `json:"appfirewallabortsperprofilerate,omitempty"`
	Appfirewalllogbufferoverflowperprofile                 int    `json:"appfirewalllogbufferoverflowperprofile,omitempty"`
	Appfirewalllogbufferoverflowperprofilerate             int    `json:"appfirewalllogbufferoverflowperprofilerate,omitempty"`
	Appfirewalllogcmdperprofile                            int    `json:"appfirewalllogcmdperprofile,omitempty"`
	Appfirewalllogcmdperprofilerate                        int    `json:"appfirewalllogcmdperprofilerate,omitempty"`
	Appfirewalllogcontenttypeperprofile                    int    `json:"appfirewalllogcontenttypeperprofile,omitempty"`
	Appfirewalllogcontenttypeperprofilerate                int    `json:"appfirewalllogcontenttypeperprofilerate,omitempty"`
	Appfirewalllogcookiehijackperprofile                   int    `json:"appfirewalllogcookiehijackperprofile,omitempty"`
	Appfirewalllogcookiehijackperprofilerate               int    `json:"appfirewalllogcookiehijackperprofilerate,omitempty"`
	Appfirewalllogcookieperprofile                         int    `json:"appfirewalllogcookieperprofile,omitempty"`
	Appfirewalllogcookieperprofilerate                     int    `json:"appfirewalllogcookieperprofilerate,omitempty"`
	Appfirewalllogcreditcardperprofile                     int    `json:"appfirewalllogcreditcardperprofile,omitempty"`
	Appfirewalllogcreditcardperprofilerate                 int    `json:"appfirewalllogcreditcardperprofilerate,omitempty"`
	Appfirewalllogcsrftagperprofile                        int    `json:"appfirewalllogcsrftagperprofile,omitempty"`
	Appfirewalllogcsrftagperprofilerate                    int    `json:"appfirewalllogcsrftagperprofilerate,omitempty"`
	Appfirewalllogdenyurlperprofile                        int    `json:"appfirewalllogdenyurlperprofile,omitempty"`
	Appfirewalllogdenyurlperprofilerate                    int    `json:"appfirewalllogdenyurlperprofilerate,omitempty"`
	Appfirewalllogfieldconsistencyperprofile               int    `json:"appfirewalllogfieldconsistencyperprofile,omitempty"`
	Appfirewalllogfieldconsistencyperprofilerate           int    `json:"appfirewalllogfieldconsistencyperprofilerate,omitempty"`
	Appfirewalllogfieldformatperprofile                    int    `json:"appfirewalllogfieldformatperprofile,omitempty"`
	Appfirewalllogfieldformatperprofilerate                int    `json:"appfirewalllogfieldformatperprofilerate,omitempty"`
	Appfirewalllogfileuploadtypesperprofile                int    `json:"appfirewalllogfileuploadtypesperprofile,omitempty"`
	Appfirewalllogfileuploadtypesperprofilerate            int    `json:"appfirewalllogfileuploadtypesperprofilerate,omitempty"`
	Appfirewallloginfercontenttypexmlpayloadperprofile     int    `json:"appfirewallloginfercontenttypexmlpayloadperprofile,omitempty"`
	Appfirewallloginfercontenttypexmlpayloadperprofilerate int    `json:"appfirewallloginfercontenttypexmlpayloadperprofilerate,omitempty"`
	Appfirewalllogperprofilerate                           int    `json:"appfirewalllogperprofilerate,omitempty"`
	Appfirewalllogrefererheaderperprofile                  int    `json:"appfirewalllogrefererheaderperprofile,omitempty"`
	Appfirewalllogrefererheaderperprofilerate              int    `json:"appfirewalllogrefererheaderperprofilerate,omitempty"`
	Appfirewalllogsafeobjectperprofile                     int    `json:"appfirewalllogsafeobjectperprofile,omitempty"`
	Appfirewalllogsafeobjectperprofilerate                 int    `json:"appfirewalllogsafeobjectperprofilerate,omitempty"`
	Appfirewalllogsjsoncmdperprofile                       int    `json:"appfirewalllogsjsoncmdperprofile,omitempty"`
	Appfirewalllogsjsoncmdperprofilerate                   int    `json:"appfirewalllogsjsoncmdperprofilerate,omitempty"`
	Appfirewalllogsjsondosperprofile                       int    `json:"appfirewalllogsjsondosperprofile,omitempty"`
	Appfirewalllogsjsondosperprofilerate                   int    `json:"appfirewalllogsjsondosperprofilerate,omitempty"`
	Appfirewalllogsjsonsqlgramperprofile                   int    `json:"appfirewalllogsjsonsqlgramperprofile,omitempty"`
	Appfirewalllogsjsonsqlgramperprofilerate               int    `json:"appfirewalllogsjsonsqlgramperprofilerate,omitempty"`
	Appfirewalllogsjsonsqlperprofile                       int    `json:"appfirewalllogsjsonsqlperprofile,omitempty"`
	Appfirewalllogsjsonsqlperprofilerate                   int    `json:"appfirewalllogsjsonsqlperprofilerate,omitempty"`
	Appfirewalllogsjsonxssperprofile                       int    `json:"appfirewalllogsjsonxssperprofile,omitempty"`
	Appfirewalllogsjsonxssperprofilerate                   int    `json:"appfirewalllogsjsonxssperprofilerate,omitempty"`
	Appfirewalllogsqlgramperprofile                        int    `json:"appfirewalllogsqlgramperprofile,omitempty"`
	Appfirewalllogsqlgramperprofilerate                    int    `json:"appfirewalllogsqlgramperprofilerate,omitempty"`
	Appfirewalllogsqlperprofile                            int    `json:"appfirewalllogsqlperprofile,omitempty"`
	Appfirewalllogsqlperprofilerate                        int    `json:"appfirewalllogsqlperprofilerate,omitempty"`
	Appfirewalllogstarturlperprofile                       int    `json:"appfirewalllogstarturlperprofile,omitempty"`
	Appfirewalllogstarturlperprofilerate                   int    `json:"appfirewalllogstarturlperprofilerate,omitempty"`
	Appfirewalllogxformsqlperprofile                       int    `json:"appfirewalllogxformsqlperprofile,omitempty"`
	Appfirewalllogxformsqlperprofilerate                   int    `json:"appfirewalllogxformsqlperprofilerate,omitempty"`
	Appfirewalllogxformxssperprofile                       int    `json:"appfirewalllogxformxssperprofile,omitempty"`
	Appfirewalllogxformxssperprofilerate                   int    `json:"appfirewalllogxformxssperprofilerate,omitempty"`
	Appfirewalllogxssperprofile                            int    `json:"appfirewalllogxssperprofile,omitempty"`
	Appfirewalllogxssperprofilerate                        int    `json:"appfirewalllogxssperprofilerate,omitempty"`
	Appfirewalllongavgresptimeperprofile                   int    `json:"appfirewalllongavgresptimeperprofile,omitempty"`
	Appfirewallmsgvallogsperprofile                        int    `json:"appfirewallmsgvallogsperprofile,omitempty"`
	Appfirewallmsgvallogsperprofilerate                    int    `json:"appfirewallmsgvallogsperprofilerate,omitempty"`
	Appfirewallpostbodylimitlogsperprofile                 int    `json:"appfirewallpostbodylimitlogsperprofile,omitempty"`
	Appfirewallpostbodylimitlogsperprofilerate             int    `json:"appfirewallpostbodylimitlogsperprofilerate,omitempty"`
	Appfirewallpostbodylimitviolationsperprofile           int    `json:"appfirewallpostbodylimitviolationsperprofile,omitempty"`
	Appfirewallpostbodylimitviolationsperprofilerate       int    `json:"appfirewallpostbodylimitviolationsperprofilerate,omitempty"`
	Appfirewallredirectsperprofile                         int    `json:"appfirewallredirectsperprofile,omitempty"`
	Appfirewallredirectsperprofilerate                     int    `json:"appfirewallredirectsperprofilerate,omitempty"`
	Appfirewallreqbytesperprofile                          int    `json:"appfirewallreqbytesperprofile,omitempty"`
	Appfirewallreqbytesperprofilerate                      int    `json:"appfirewallreqbytesperprofilerate,omitempty"`
	Appfirewallrequestsperprofile                          int    `json:"appfirewallrequestsperprofile,omitempty"`
	Appfirewallrequestsperprofilerate                      int    `json:"appfirewallrequestsperprofilerate,omitempty"`
	Appfirewallresbytesperprofile                          int    `json:"appfirewallresbytesperprofile,omitempty"`
	Appfirewallresbytesperprofilerate                      int    `json:"appfirewallresbytesperprofilerate,omitempty"`
	Appfirewallresponsesperprofile                         int    `json:"appfirewallresponsesperprofile,omitempty"`
	Appfirewallresponsesperprofilerate                     int    `json:"appfirewallresponsesperprofilerate,omitempty"`
	Appfirewallret4xxperprofile                            int    `json:"appfirewallret4xxperprofile,omitempty"`
	Appfirewallret4xxperprofilerate                        int    `json:"appfirewallret4xxperprofilerate,omitempty"`
	Appfirewallret5xxperprofile                            int    `json:"appfirewallret5xxperprofile,omitempty"`
	Appfirewallret5xxperprofilerate                        int    `json:"appfirewallret5xxperprofilerate,omitempty"`
	Appfirewallshortavgresptimeperprofile                  int    `json:"appfirewallshortavgresptimeperprofile,omitempty"`
	Appfirewalltotallogperprofile                          int    `json:"appfirewalltotallogperprofile,omitempty"`
	Appfirewalltotalviolperprofile                         int    `json:"appfirewalltotalviolperprofile,omitempty"`
	Appfirewallviolbufferoverflowperprofile                int    `json:"appfirewallviolbufferoverflowperprofile,omitempty"`
	Appfirewallviolbufferoverflowperprofilerate            int    `json:"appfirewallviolbufferoverflowperprofilerate,omitempty"`
	Appfirewallviolcmdperprofile                           int    `json:"appfirewallviolcmdperprofile,omitempty"`
	Appfirewallviolcmdperprofilerate                       int    `json:"appfirewallviolcmdperprofilerate,omitempty"`
	Appfirewallviolcontenttypeperprofile                   int    `json:"appfirewallviolcontenttypeperprofile,omitempty"`
	Appfirewallviolcontenttypeperprofilerate               int    `json:"appfirewallviolcontenttypeperprofilerate,omitempty"`
	Appfirewallviolcookiehijackperprofile                  int    `json:"appfirewallviolcookiehijackperprofile,omitempty"`
	Appfirewallviolcookiehijackperprofilerate              int    `json:"appfirewallviolcookiehijackperprofilerate,omitempty"`
	Appfirewallviolcookieperprofile                        int    `json:"appfirewallviolcookieperprofile,omitempty"`
	Appfirewallviolcookieperprofilerate                    int    `json:"appfirewallviolcookieperprofilerate,omitempty"`
	Appfirewallviolcreditcardperprofile                    int    `json:"appfirewallviolcreditcardperprofile,omitempty"`
	Appfirewallviolcreditcardperprofilerate                int    `json:"appfirewallviolcreditcardperprofilerate,omitempty"`
	Appfirewallviolcsrftagperprofile                       int    `json:"appfirewallviolcsrftagperprofile,omitempty"`
	Appfirewallviolcsrftagperprofilerate                   int    `json:"appfirewallviolcsrftagperprofilerate,omitempty"`
	Appfirewallvioldenyurlperprofile                       int    `json:"appfirewallvioldenyurlperprofile,omitempty"`
	Appfirewallvioldenyurlperprofilerate                   int    `json:"appfirewallvioldenyurlperprofilerate,omitempty"`
	Appfirewallviolfieldconsistencyperprofile              int    `json:"appfirewallviolfieldconsistencyperprofile,omitempty"`
	Appfirewallviolfieldconsistencyperprofilerate          int    `json:"appfirewallviolfieldconsistencyperprofilerate,omitempty"`
	Appfirewallviolfieldformatperprofile                   int    `json:"appfirewallviolfieldformatperprofile,omitempty"`
	Appfirewallviolfieldformatperprofilerate               int    `json:"appfirewallviolfieldformatperprofilerate,omitempty"`
	Appfirewallviolfileuploadtypesperprofile               int    `json:"appfirewallviolfileuploadtypesperprofile,omitempty"`
	Appfirewallviolfileuploadtypesperprofilerate           int    `json:"appfirewallviolfileuploadtypesperprofilerate,omitempty"`
	Appfirewallvioljsoncmdperprofile                       int    `json:"appfirewallvioljsoncmdperprofile,omitempty"`
	Appfirewallvioljsoncmdperprofilerate                   int    `json:"appfirewallvioljsoncmdperprofilerate,omitempty"`
	Appfirewallvioljsondosperprofile                       int    `json:"appfirewallvioljsondosperprofile,omitempty"`
	Appfirewallvioljsondosperprofilerate                   int    `json:"appfirewallvioljsondosperprofilerate,omitempty"`
	Appfirewallvioljsonsqlgramperprofile                   int    `json:"appfirewallvioljsonsqlgramperprofile,omitempty"`
	Appfirewallvioljsonsqlgramperprofilerate               int    `json:"appfirewallvioljsonsqlgramperprofilerate,omitempty"`
	Appfirewallvioljsonsqlperprofile                       int    `json:"appfirewallvioljsonsqlperprofile,omitempty"`
	Appfirewallvioljsonsqlperprofilerate                   int    `json:"appfirewallvioljsonsqlperprofilerate,omitempty"`
	Appfirewallvioljsonxssperprofile                       int    `json:"appfirewallvioljsonxssperprofile,omitempty"`
	Appfirewallvioljsonxssperprofilerate                   int    `json:"appfirewallvioljsonxssperprofilerate,omitempty"`
	Appfirewallviolmsgvalviolationsperprofile              int    `json:"appfirewallviolmsgvalviolationsperprofile,omitempty"`
	Appfirewallviolmsgvalviolationsperprofilerate          int    `json:"appfirewallviolmsgvalviolationsperprofilerate,omitempty"`
	Appfirewallviolperprofilerate                          int    `json:"appfirewallviolperprofilerate,omitempty"`
	Appfirewallviolrefererheaderperprofile                 int    `json:"appfirewallviolrefererheaderperprofile,omitempty"`
	Appfirewallviolrefererheaderperprofilerate             int    `json:"appfirewallviolrefererheaderperprofilerate,omitempty"`
	Appfirewallviolsafeobjectperprofile                    int    `json:"appfirewallviolsafeobjectperprofile,omitempty"`
	Appfirewallviolsafeobjectperprofilerate                int    `json:"appfirewallviolsafeobjectperprofilerate,omitempty"`
	Appfirewallviolsignatureperprofile                     int    `json:"appfirewallviolsignatureperprofile,omitempty"`
	Appfirewallviolsignatureperprofilerate                 int    `json:"appfirewallviolsignatureperprofilerate,omitempty"`
	Appfirewallviolsqlgramperprofile                       int    `json:"appfirewallviolsqlgramperprofile,omitempty"`
	Appfirewallviolsqlgramperprofilerate                   int    `json:"appfirewallviolsqlgramperprofilerate,omitempty"`
	Appfirewallviolsqlperprofile                           int    `json:"appfirewallviolsqlperprofile,omitempty"`
	Appfirewallviolsqlperprofilerate                       int    `json:"appfirewallviolsqlperprofilerate,omitempty"`
	Appfirewallviolstarturlperprofile                      int    `json:"appfirewallviolstarturlperprofile,omitempty"`
	Appfirewallviolstarturlperprofilerate                  int    `json:"appfirewallviolstarturlperprofilerate,omitempty"`
	Appfirewallviolwellformednessviolationsperprofile      int    `json:"appfirewallviolwellformednessviolationsperprofile,omitempty"`
	Appfirewallviolwellformednessviolationsperprofilerate  int    `json:"appfirewallviolwellformednessviolationsperprofilerate,omitempty"`
	Appfirewallviolwsiviolationsperprofile                 int    `json:"appfirewallviolwsiviolationsperprofile,omitempty"`
	Appfirewallviolwsiviolationsperprofilerate             int    `json:"appfirewallviolwsiviolationsperprofilerate,omitempty"`
	Appfirewallviolxdosviolationsperprofile                int    `json:"appfirewallviolxdosviolationsperprofile,omitempty"`
	Appfirewallviolxdosviolationsperprofilerate            int    `json:"appfirewallviolxdosviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlattachmentviolationsperprofile       int    `json:"appfirewallviolxmlattachmentviolationsperprofile,omitempty"`
	Appfirewallviolxmlattachmentviolationsperprofilerate   int    `json:"appfirewallviolxmlattachmentviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlgenericviolationsperprofile          int    `json:"appfirewallviolxmlgenericviolationsperprofile,omitempty"`
	Appfirewallviolxmlgenericviolationsperprofilerate      int    `json:"appfirewallviolxmlgenericviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlsoapfaultviolationsperprofile        int    `json:"appfirewallviolxmlsoapfaultviolationsperprofile,omitempty"`
	Appfirewallviolxmlsoapfaultviolationsperprofilerate    int    `json:"appfirewallviolxmlsoapfaultviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlsqlviolationsperprofile              int    `json:"appfirewallviolxmlsqlviolationsperprofile,omitempty"`
	Appfirewallviolxmlsqlviolationsperprofilerate          int    `json:"appfirewallviolxmlsqlviolationsperprofilerate,omitempty"`
	Appfirewallviolxmlxssviolationsperprofile              int    `json:"appfirewallviolxmlxssviolationsperprofile,omitempty"`
	Appfirewallviolxmlxssviolationsperprofilerate          int    `json:"appfirewallviolxmlxssviolationsperprofilerate,omitempty"`
	Appfirewallviolxssperprofile                           int    `json:"appfirewallviolxssperprofile,omitempty"`
	Appfirewallviolxssperprofilerate                       int    `json:"appfirewallviolxssperprofilerate,omitempty"`
	Appfirewallwellformednesslogsperprofile                int    `json:"appfirewallwellformednesslogsperprofile,omitempty"`
	Appfirewallwellformednesslogsperprofilerate            int    `json:"appfirewallwellformednesslogsperprofilerate,omitempty"`
	Appfirewallwsilogsperprofile                           int    `json:"appfirewallwsilogsperprofile,omitempty"`
	Appfirewallwsilogsperprofilerate                       int    `json:"appfirewallwsilogsperprofilerate,omitempty"`
	Appfirewallxdoslogsperprofile                          int    `json:"appfirewallxdoslogsperprofile,omitempty"`
	Appfirewallxdoslogsperprofilerate                      int    `json:"appfirewallxdoslogsperprofilerate,omitempty"`
	Appfirewallxformlogcreditcardperprofile                int    `json:"appfirewallxformlogcreditcardperprofile,omitempty"`
	Appfirewallxformlogcreditcardperprofilerate            int    `json:"appfirewallxformlogcreditcardperprofilerate,omitempty"`
	Appfirewallxmlattachmentlogsperprofile                 int    `json:"appfirewallxmlattachmentlogsperprofile,omitempty"`
	Appfirewallxmlattachmentlogsperprofilerate             int    `json:"appfirewallxmlattachmentlogsperprofilerate,omitempty"`
	Appfirewallxmlgenericlogsperprofile                    int    `json:"appfirewallxmlgenericlogsperprofile,omitempty"`
	Appfirewallxmlgenericlogsperprofilerate                int    `json:"appfirewallxmlgenericlogsperprofilerate,omitempty"`
	Appfirewallxmlpayloadcontenttypemismatchperprofile     int    `json:"appfirewallxmlpayloadcontenttypemismatchperprofile,omitempty"`
	Appfirewallxmlpayloadcontenttypemismatchperprofilerate int    `json:"appfirewallxmlpayloadcontenttypemismatchperprofilerate,omitempty"`
	Appfirewallxmlsoapfaultlogsperprofile                  int    `json:"appfirewallxmlsoapfaultlogsperprofile,omitempty"`
	Appfirewallxmlsoapfaultlogsperprofilerate              int    `json:"appfirewallxmlsoapfaultlogsperprofilerate,omitempty"`
	Appfirewallxmlsqllogsperprofile                        int    `json:"appfirewallxmlsqllogsperprofile,omitempty"`
	Appfirewallxmlsqllogsperprofilerate                    int    `json:"appfirewallxmlsqllogsperprofilerate,omitempty"`
	Appfirewallxmlxsslogsperprofile                        int    `json:"appfirewallxmlxsslogsperprofile,omitempty"`
	Appfirewallxmlxsslogsperprofilerate                    int    `json:"appfirewallxmlxsslogsperprofilerate,omitempty"`
	Clearstats                                             string `json:"clearstats,omitempty"`
	Name                                                   string `json:"name,omitempty"`
}