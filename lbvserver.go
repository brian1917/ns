package ns

import (
	"encoding/json"
	"fmt"
)

// LBServerResp is the response from the LBServer API
type LbServerResp struct {
	Errorcode      int             `json:"errorcode"`
	Message        string          `json:"message"`
	Severity       string          `json:"severity"`
	VirtualServers []VirtualServer `json:"lbvserver"`
}

// LBVserver is a virtual server in Netscaler
type VirtualServer struct {
	Name                       string `json:"name"`
	Insertvserveripport        string `json:"insertvserveripport"`
	Ipv46                      string `json:"ipv46"`
	Ippattern                  string `json:"ippattern"`
	Ipmask                     string `json:"ipmask"`
	Listenpolicy               string `json:"listenpolicy"`
	Ipmapping                  string `json:"ipmapping"`
	Port                       int    `json:"port"`
	Range                      string `json:"range"`
	Servicetype                string `json:"servicetype"`
	Type                       string `json:"type"`
	Curstate                   string `json:"curstate"`
	Effectivestate             string `json:"effectivestate"`
	Status                     int    `json:"status"`
	Lbrrreason                 int    `json:"lbrrreason"`
	Cachetype                  string `json:"cachetype"`
	Authentication             string `json:"authentication"`
	Authn401                   string `json:"authn401"`
	Dynamicweight              string `json:"dynamicweight"`
	Priority                   string `json:"priority"`
	Clttimeout                 string `json:"clttimeout"`
	Somethod                   string `json:"somethod"`
	Sopersistence              string `json:"sopersistence"`
	Sopersistencetimeout       string `json:"sopersistencetimeout"`
	Healththreshold            string `json:"healththreshold"`
	Lbmethod                   string `json:"lbmethod"`
	Dataoffset                 string `json:"dataoffset"`
	Health                     string `json:"health"`
	Datalength                 string `json:"datalength"`
	Ruletype                   string `json:"ruletype"`
	M                          string `json:"m"`
	Persistencetype            string `json:"persistencetype"`
	Timeout                    int    `json:"timeout"`
	Persistmask                string `json:"persistmask"`
	V6Persistmasklen           string `json:"v6persistmasklen"`
	Persistencebackup          string `json:"persistencebackup"`
	Backuppersistencetimeout   int    `json:"backuppersistencetimeout"`
	Cacheable                  string `json:"cacheable"`
	Rtspnat                    string `json:"rtspnat"`
	Sessionless                string `json:"sessionless"`
	Trofspersistence           string `json:"trofspersistence"`
	Map                        string `json:"map"`
	Connfailover               string `json:"connfailover"`
	Redirectportrewrite        string `json:"redirectportrewrite"`
	Downstateflush             string `json:"downstateflush"`
	Disableprimaryondown       string `json:"disableprimaryondown"`
	Gt2Gb                      string `json:"gt2gb"`
	Consolidatedlconn          string `json:"consolidatedlconn"`
	Consolidatedlconngbl       string `json:"consolidatedlconngbl"`
	Thresholdvalue             int    `json:"thresholdvalue"`
	Invoke                     bool   `json:"invoke"`
	Version                    int    `json:"version"`
	Totalservices              string `json:"totalservices"`
	Activeservices             string `json:"activeservices"`
	Statechangetimesec         string `json:"statechangetimesec"`
	Statechangetimeseconds     string `json:"statechangetimeseconds"`
	Statechangetimemsec        string `json:"statechangetimemsec"`
	Tickssincelaststatechange  string `json:"tickssincelaststatechange"`
	Hits                       string `json:"hits"`
	Pipolicyhits               string `json:"pipolicyhits"`
	Push                       string `json:"push"`
	Pushlabel                  string `json:"pushlabel"`
	Pushmulticlients           string `json:"pushmulticlients"`
	Httpprofilename            string `json:"httpprofilename,omitempty"`
	Policysubtype              string `json:"policysubtype"`
	L2Conn                     string `json:"l2conn"`
	Appflowlog                 string `json:"appflowlog"`
	Isgslb                     bool   `json:"isgslb"`
	Icmpvsrresponse            string `json:"icmpvsrresponse"`
	Rhistate                   string `json:"rhistate"`
	Newservicerequestunit      string `json:"newservicerequestunit"`
	Vsvrbindsvcip              string `json:"vsvrbindsvcip"`
	Vsvrbindsvcport            int    `json:"vsvrbindsvcport"`
	Skippersistency            string `json:"skippersistency"`
	Td                         string `json:"td"`
	Minautoscalemembers        string `json:"minautoscalemembers"`
	Maxautoscalemembers        string `json:"maxautoscalemembers"`
	Macmoderetainvlan          string `json:"macmoderetainvlan"`
	DNS64                      string `json:"dns64"`
	Bypassaaaa                 string `json:"bypassaaaa"`
	Processlocal               string `json:"processlocal"`
	Vsvrdynconnsothreshold     string `json:"vsvrdynconnsothreshold"`
	Retainconnectionsoncluster string `json:"retainconnectionsoncluster"`
	Nodefaultbindings          string `json:"nodefaultbindings"`
}

// GetVirtualServers returns all the Virtual Servers in a NetScaler
func (n *NetScaler) GetVirtualServers() (virtualServers []VirtualServer, api APIResp, err error) {

	// Get the Virtual Servers
	api, err = n.API("lbvserver", "GET", nil)
	if err != nil {
		return nil, api, fmt.Errorf("error getting virtual servers - %s", err)
	}
	if api.StatusCode != 200 {
		return nil, api, fmt.Errorf("expected 200. received %d - %s.", api.StatusCode, api.RespBody)
	}

	// Marshall the API response
	var resp LbServerResp
	json.Unmarshal([]byte(api.RespBody), &resp)

	// Check the API error code
	if resp.Errorcode != 0 {
		return resp.VirtualServers, api, fmt.Errorf("error code of %d", resp.Errorcode)
	}

	// Return
	return resp.VirtualServers, api, nil
}
