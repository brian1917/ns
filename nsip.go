package ns

import (
	"encoding/json"
	"fmt"
)

// NSIPResp is the response from the nsip API
type NSIPResp struct {
	Errorcode int    `json:"errorcode"`
	Message   string `json:"message"`
	Severity  string `json:"severity"`
	NSIPs     []NSIP `json:"nsip"`
}

// NSIP is a NSIP in Netscaler
type NSIP struct {
	Ipaddress                   string   `json:"ipaddress"`
	Td                          string   `json:"td"`
	Type                        string   `json:"type"`
	Netmask                     string   `json:"netmask"`
	Flags                       string   `json:"flags"`
	Arp                         string   `json:"arp"`
	Icmp                        string   `json:"icmp"`
	Vserver                     string   `json:"vserver"`
	Telnet                      string   `json:"telnet"`
	SSH                         string   `json:"ssh"`
	Gui                         string   `json:"gui"`
	Snmp                        string   `json:"snmp"`
	Ftp                         string   `json:"ftp"`
	Mgmtaccess                  string   `json:"mgmtaccess"`
	Restrictaccess              string   `json:"restrictaccess"`
	Decrementttl                string   `json:"decrementttl"`
	Dynamicrouting              string   `json:"dynamicrouting"`
	Hostroute                   string   `json:"hostroute"`
	Advertiseondefaultpartition string   `json:"advertiseondefaultpartition"`
	Networkroute                string   `json:"networkroute"`
	Tag                         string   `json:"tag"`
	Hostrtgwact                 string   `json:"hostrtgwact"`
	Metric                      int      `json:"metric"`
	Ospfareaval                 string   `json:"ospfareaval"`
	Vserverrhilevel             string   `json:"vserverrhilevel"`
	Viprtadv2Bsd                bool     `json:"viprtadv2bsd"`
	Vipvsercount                string   `json:"vipvsercount"`
	Vipvserdowncount            string   `json:"vipvserdowncount"`
	Vipvsrvrrhiactivecount      string   `json:"vipvsrvrrhiactivecount"`
	Vipvsrvrrhiactiveupcount    string   `json:"vipvsrvrrhiactiveupcount"`
	Ospflsatype                 string   `json:"ospflsatype"`
	State                       string   `json:"state"`
	Freeports                   string   `json:"freeports"`
	Iptype                      []string `json:"iptype"`
	Icmpresponse                string   `json:"icmpresponse"`
	Ownernode                   string   `json:"ownernode"`
	Arpresponse                 string   `json:"arpresponse"`
	Ownerdownresponse           string   `json:"ownerdownresponse"`
	Arpowner                    string   `json:"arpowner"`
	Operationalarpowner         string   `json:"operationalarpowner"`
	Mptcpadvertise              string   `json:"mptcpadvertise"`
}

// GetNSIPs returns all the Virtual Servers in a NetScaler
func (n *NetScaler) GetNSIP() (NSIPs []NSIP, api APIResp, err error) {

	// Get the Virtual Servers
	api, err = n.API("nsip", "GET", nil)
	if err != nil {
		return nil, api, fmt.Errorf("error getting nsips - %s", err)
	}
	if api.StatusCode != 200 {
		return nil, api, fmt.Errorf("expected 200. received %d - %s.", api.StatusCode, api.RespBody)
	}

	// Marshall the API response
	var resp NSIPResp
	json.Unmarshal([]byte(api.RespBody), &resp)

	// Check the API error code
	if resp.Errorcode != 0 {
		return resp.NSIPs, api, fmt.Errorf("error code of %d", resp.Errorcode)
	}

	// Return
	return resp.NSIPs, api, nil
}
