package qos

type Qosstats struct {
	Clearstats                               string `json:"clearstats,omitempty"`
	Name                                     string `json:"name,omitempty"`
	Snmpqosqosactionscreated                 int    `json:"snmpqosqos_actions_created,omitempty"`
	Snmpqosqosactionscreatedrate             int    `json:"snmpqosqos_actions_createdrate,omitempty"`
	Snmpqosqosbytesrx                        int    `json:"snmpqosqos_bytes_rx,omitempty"`
	Snmpqosqosbytesrxrate                    int    `json:"snmpqosqos_bytes_rxrate,omitempty"`
	Snmpqosqosbytestx                        int    `json:"snmpqosqos_bytes_tx,omitempty"`
	Snmpqosqosbytestxrate                    int    `json:"snmpqosqos_bytes_txrate,omitempty"`
	Snmpqosqoscfytcpunknown                  int    `json:"snmpqosqos_cfy_tcp_unknown,omitempty"`
	Snmpqosqoscfytcpunknownrate              int    `json:"snmpqosqos_cfy_tcp_unknownrate,omitempty"`
	Snmpqosqoscfyudpunknown                  int    `json:"snmpqosqos_cfy_udp_unknown,omitempty"`
	Snmpqosqoscfyudpunknownrate              int    `json:"snmpqosqos_cfy_udp_unknownrate,omitempty"`
	Snmpqosqoserrorapisesaddinsession        int    `json:"snmpqosqos_error_api_ses_add_insession,omitempty"`
	Snmpqosqoserrorapisesaddinsessionrate    int    `json:"snmpqosqos_error_api_ses_add_insessionrate,omitempty"`
	Snmpqosqoserrorapisesaddother            int    `json:"snmpqosqos_error_api_ses_add_other,omitempty"`
	Snmpqosqoserrorapisesaddotherrate        int    `json:"snmpqosqos_error_api_ses_add_otherrate,omitempty"`
	Snmpqosqoserrorapisesdel                 int    `json:"snmpqosqos_error_api_ses_del,omitempty"`
	Snmpqosqoserrorapisesdelrate             int    `json:"snmpqosqos_error_api_ses_delrate,omitempty"`
	Snmpqosqoserrorapisesinvalidpcb          int    `json:"snmpqosqos_error_api_ses_invalidpcb,omitempty"`
	Snmpqosqoserrorapisesinvalidpcbrate      int    `json:"snmpqosqos_error_api_ses_invalidpcbrate,omitempty"`
	Snmpqosqoserrorapisesnotready            int    `json:"snmpqosqos_error_api_ses_notready,omitempty"`
	Snmpqosqoserrorapisesnotreadyrate        int    `json:"snmpqosqos_error_api_ses_notreadyrate,omitempty"`
	Snmpqosqoserrorapisesremnotinsession     int    `json:"snmpqosqos_error_api_ses_rem_notinsession,omitempty"`
	Snmpqosqoserrorapisesremnotinsessionrate int    `json:"snmpqosqos_error_api_ses_rem_notinsessionrate,omitempty"`
	Snmpqosqoserrorapisesremother            int    `json:"snmpqosqos_error_api_ses_rem_other,omitempty"`
	Snmpqosqoserrorapisesremotherrate        int    `json:"snmpqosqos_error_api_ses_rem_otherrate,omitempty"`
	Snmpqosqoserrorcliunknown                int    `json:"snmpqosqos_error_cli_unknown,omitempty"`
	Snmpqosqoserrorcliunknownrate            int    `json:"snmpqosqos_error_cli_unknownrate,omitempty"`
	Snmpqosqoserrorcreateactionfailed        int    `json:"snmpqosqos_error_create_action_failed,omitempty"`
	Snmpqosqoserrorcreateactionfailedrate    int    `json:"snmpqosqos_error_create_action_failedrate,omitempty"`
	Snmpqosqoserrorcreatepolicyfailed        int    `json:"snmpqosqos_error_create_policy_failed,omitempty"`
	Snmpqosqoserrorcreatepolicyfailedrate    int    `json:"snmpqosqos_error_create_policy_failedrate,omitempty"`
	Snmpqosqoserroripc                       int    `json:"snmpqosqos_error_ipc,omitempty"`
	Snmpqosqoserroripcrate                   int    `json:"snmpqosqos_error_ipcrate,omitempty"`
	Snmpqosqoserrorlibqosapifailures         int    `json:"snmpqosqos_error_libqos_api_failures,omitempty"`
	Snmpqosqoserrorlibqosapifailuresrate     int    `json:"snmpqosqos_error_libqos_api_failuresrate,omitempty"`
	Snmpqosqoserrormodifyactionfailed        int    `json:"snmpqosqos_error_modify_action_failed,omitempty"`
	Snmpqosqoserrormodifyactionfailedrate    int    `json:"snmpqosqos_error_modify_action_failedrate,omitempty"`
	Snmpqosqoserrornoflows                   int    `json:"snmpqosqos_error_no_flows,omitempty"`
	Snmpqosqoserrornoflowsrate               int    `json:"snmpqosqos_error_no_flowsrate,omitempty"`
	Snmpqosqoserrornosessions                int    `json:"snmpqosqos_error_no_sessions,omitempty"`
	Snmpqosqoserrornosessionsrate            int    `json:"snmpqosqos_error_no_sessionsrate,omitempty"`
	Snmpqosqoserrorremoveactionfailed        int    `json:"snmpqosqos_error_remove_action_failed,omitempty"`
	Snmpqosqoserrorremoveactionfailedrate    int    `json:"snmpqosqos_error_remove_action_failedrate,omitempty"`
	Snmpqosqoserrorremovepolicyfailed        int    `json:"snmpqosqos_error_remove_policy_failed,omitempty"`
	Snmpqosqoserrorremovepolicyfailedrate    int    `json:"snmpqosqos_error_remove_policy_failedrate,omitempty"`
	Snmpqosqoserrorrenamenotimplemented      int    `json:"snmpqosqos_error_rename_not_implemented,omitempty"`
	Snmpqosqoserrorrenamenotimplementedrate  int    `json:"snmpqosqos_error_rename_not_implementedrate,omitempty"`
	Snmpqosqosflowmem                        int    `json:"snmpqosqos_flow_mem,omitempty"`
	Snmpqosqosflowmemrate                    int    `json:"snmpqosqos_flow_memrate,omitempty"`
	Snmpqosqosflowrecycles                   int    `json:"snmpqosqos_flow_recycles,omitempty"`
	Snmpqosqosflowrecyclesrate               int    `json:"snmpqosqos_flow_recyclesrate,omitempty"`
	Snmpqosqosflows                          int    `json:"snmpqosqos_flows,omitempty"`
	Snmpqosqosflowsavailable                 int    `json:"snmpqosqos_flows_available,omitempty"`
	Snmpqosqosflowsavailablerate             int    `json:"snmpqosqos_flows_availablerate,omitempty"`
	Snmpqosqosflowsrate                      int    `json:"snmpqosqos_flowsrate,omitempty"`
	Snmpqosqoslazybytes                      int    `json:"snmpqosqos_lazy_bytes,omitempty"`
	Snmpqosqoslazybytesrate                  int    `json:"snmpqosqos_lazy_bytesrate,omitempty"`
	Snmpqosqospacketsbypassed                int    `json:"snmpqosqos_packets_bypassed,omitempty"`
	Snmpqosqospacketsbypassedrate            int    `json:"snmpqosqos_packets_bypassedrate,omitempty"`
	Snmpqosqospacketsclassified              int    `json:"snmpqosqos_packets_classified,omitempty"`
	Snmpqosqospacketsclassifiedrate          int    `json:"snmpqosqos_packets_classifiedrate,omitempty"`
	Snmpqosqospacketsdropped                 int    `json:"snmpqosqos_packets_dropped,omitempty"`
	Snmpqosqospacketsdroppedrate             int    `json:"snmpqosqos_packets_droppedrate,omitempty"`
	Snmpqosqospacketsfiltered                int    `json:"snmpqosqos_packets_filtered,omitempty"`
	Snmpqosqospacketsfilteredrate            int    `json:"snmpqosqos_packets_filteredrate,omitempty"`
	Snmpqosqospacketsreceived                int    `json:"snmpqosqos_packets_received,omitempty"`
	Snmpqosqospacketsreceivedrate            int    `json:"snmpqosqos_packets_receivedrate,omitempty"`
	Snmpqosqospacketssent                    int    `json:"snmpqosqos_packets_sent,omitempty"`
	Snmpqosqospacketssentrate                int    `json:"snmpqosqos_packets_sentrate,omitempty"`
	Snmpqosqospolicyreeval                   int    `json:"snmpqosqos_policy_reeval,omitempty"`
	Snmpqosqospolicyreevalrate               int    `json:"snmpqosqos_policy_reevalrate,omitempty"`
	Snmpqosqosrealbytes                      int    `json:"snmpqosqos_real_bytes,omitempty"`
	Snmpqosqosrealbytesrate                  int    `json:"snmpqosqos_real_bytesrate,omitempty"`
	Snmpqosqosrecyclefailedbacklog           int    `json:"snmpqosqos_recycle_failed_backlog,omitempty"`
	Snmpqosqosrecyclefailedbacklograte       int    `json:"snmpqosqos_recycle_failed_backlograte,omitempty"`
	Snmpqosqosrecyclefailedsession           int    `json:"snmpqosqos_recycle_failed_session,omitempty"`
	Snmpqosqosrecyclefailedsessionrate       int    `json:"snmpqosqos_recycle_failed_sessionrate,omitempty"`
	Snmpqosqosschleafrecyclefailures         int    `json:"snmpqosqos_sch_leaf_recycle_failures,omitempty"`
	Snmpqosqosschleafrecyclefailuresrate     int    `json:"snmpqosqos_sch_leaf_recycle_failuresrate,omitempty"`
	Snmpqosqosschleafs                       int    `json:"snmpqosqos_sch_leafs,omitempty"`
	Snmpqosqosschleafsrate                   int    `json:"snmpqosqos_sch_leafsrate,omitempty"`
	Snmpqosqosschlinkscreated                int    `json:"snmpqosqos_sch_links_created,omitempty"`
	Snmpqosqosschlinkscreatedrate            int    `json:"snmpqosqos_sch_links_createdrate,omitempty"`
	Snmpqosqosschlinksdeleted                int    `json:"snmpqosqos_sch_links_deleted,omitempty"`
	Snmpqosqosschlinksdeletedrate            int    `json:"snmpqosqos_sch_links_deletedrate,omitempty"`
	Snmpqosqosschlinksupdated                int    `json:"snmpqosqos_sch_links_updated,omitempty"`
	Snmpqosqosschlinksupdatedrate            int    `json:"snmpqosqos_sch_links_updatedrate,omitempty"`
	Snmpqosqosschnoderegulatedcount          int    `json:"snmpqosqos_sch_node_regulated_count,omitempty"`
	Snmpqosqosschnoderegulatedcountrate      int    `json:"snmpqosqos_sch_node_regulated_countrate,omitempty"`
	Snmpqosqosschpeermsgs                    int    `json:"snmpqosqos_sch_peer_msgs,omitempty"`
	Snmpqosqosschpeermsgsrate                int    `json:"snmpqosqos_sch_peer_msgsrate,omitempty"`
	Snmpqosqosschpollcount                   int    `json:"snmpqosqos_sch_poll_count,omitempty"`
	Snmpqosqosschpollcountrate               int    `json:"snmpqosqos_sch_poll_countrate,omitempty"`
	Snmpqosqosschregulatedcount              int    `json:"snmpqosqos_sch_regulated_count,omitempty"`
	Snmpqosqosschregulatedcountrate          int    `json:"snmpqosqos_sch_regulated_countrate,omitempty"`
	Snmpqosqosschsdrrnodes                   int    `json:"snmpqosqos_sch_sdrr_nodes,omitempty"`
	Snmpqosqosschsdrrnodesrate               int    `json:"snmpqosqos_sch_sdrr_nodesrate,omitempty"`
	Snmpqosqosschsessionconns                int    `json:"snmpqosqos_sch_session_conns,omitempty"`
	Snmpqosqosschsessionconnsrate            int    `json:"snmpqosqos_sch_session_connsrate,omitempty"`
	Snmpqosqosschsessionconnsremoved         int    `json:"snmpqosqos_sch_session_conns_removed,omitempty"`
	Snmpqosqosschsessionconnsremovedrate     int    `json:"snmpqosqos_sch_session_conns_removedrate,omitempty"`
	Snmpqosqosschsessionsbytecount           int    `json:"snmpqosqos_sch_sessions_byte_count,omitempty"`
	Snmpqosqosschsessionsbytecountrate       int    `json:"snmpqosqos_sch_sessions_byte_countrate,omitempty"`
	Snmpqosqosschsessionscreated             int    `json:"snmpqosqos_sch_sessions_created,omitempty"`
	Snmpqosqosschsessionscreatedrate         int    `json:"snmpqosqos_sch_sessions_createdrate,omitempty"`
	Snmpqosqosschsessionsdeleted             int    `json:"snmpqosqos_sch_sessions_deleted,omitempty"`
	Snmpqosqosschsessionsdeletedrate         int    `json:"snmpqosqos_sch_sessions_deletedrate,omitempty"`
	Snmpqosqosschsessionsregulatedcount      int    `json:"snmpqosqos_sch_sessions_regulated_count,omitempty"`
	Snmpqosqosschsessionsregulatedcountrate  int    `json:"snmpqosqos_sch_sessions_regulated_countrate,omitempty"`
	Snmpqosqosschvirtualbytesaccepted        int    `json:"snmpqosqos_sch_virtual_bytes_accepted,omitempty"`
	Snmpqosqosschvirtualbytesacceptedrate    int    `json:"snmpqosqos_sch_virtual_bytes_acceptedrate,omitempty"`
	Snmpqosqosschvirtualpackets              int    `json:"snmpqosqos_sch_virtual_packets,omitempty"`
	Snmpqosqosschvirtualpacketsrate          int    `json:"snmpqosqos_sch_virtual_packetsrate,omitempty"`
	Snmpqosqossessionmem                     int    `json:"snmpqosqos_session_mem,omitempty"`
	Snmpqosqossessionmemrate                 int    `json:"snmpqosqos_session_memrate,omitempty"`
	Snmpqosqossessionrecyclefailure          int    `json:"snmpqosqos_session_recycle_failure,omitempty"`
	Snmpqosqossessionrecyclefailurerate      int    `json:"snmpqosqos_session_recycle_failurerate,omitempty"`
	Snmpqosqossessionsconsumed               int    `json:"snmpqosqos_sessions_consumed,omitempty"`
	Snmpqosqossessionsconsumedrate           int    `json:"snmpqosqos_sessions_consumedrate,omitempty"`
	Snmpqosqossessionsignored                int    `json:"snmpqosqos_sessions_ignored,omitempty"`
	Snmpqosqossessionsignoredrate            int    `json:"snmpqosqos_sessions_ignoredrate,omitempty"`
}
