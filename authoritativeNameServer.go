package goWHOIS

import (
	"strings"
)

var TLDArray = [...]string{"ac", "ad", "ae", "aero", "af", "ag", "ai", "al", "am",
	"an", "ao", "aq", "ar", "arpa", "as", "asia", "at", "au", "aw", "ax", "az",
	"ba", "bb", "bd", "be", "bf", "bg", "bh", "bi", "biz", "bj", "bm", "bn", "bo",
	"br", "bs", "bt", "bv", "bw", "by", "bz", "ca", "cat", "cc", "cd", "cf", "cg",
	"ch", "ci", "ck", "cl", "cm", "cn", "co", "com", "coop", "cr", "cu", "cv",
	"cx", "cy", "cz", "de", "dj", "dk", "dm", "do", "dz", "ec", "edu", "ee", "eg",
	"er", "es", "et", "eu", "fi", "fj", "fk", "fm", "fo", "fr", "ga", "gb", "gd",
	"ge", "gf", "gg", "gh", "gi", "gl", "gm", "gn", "gov", "gp", "gq", "gr", "gs",
	"gt", "gu", "gw", "gy", "hk", "hm", "hn", "hr", "ht", "hu", "id", "ie", "il",
	"im", "in", "info", "int", "io", "iq", "ir", "is", "it", "je", "jm", "jo",
	"jobs", "jp", "ke", "kg", "kh", "ki", "km", "kn", "kp", "kr", "kw", "ky",
	"kz", "la", "lb", "lc", "li", "lk", "lr", "ls", "lt", "lu", "lv", "ly",
	"ma", "mc", "md", "me", "mg", "mh", "mil", "mk", "ml", "mm", "mn", "mo",
	"mobi", "mp", "mq", "mr", "ms", "mt", "mu", "mv", "mw", "mx", "my", "mz",
	"na", "name", "nc", "ne", "net", "nf", "ng", "ni", "nl", "no", "np", "nr",
	"nu", "nz", "om", "org", "pa", "pe", "pf", "pg", "ph", "pk", "pl", "pm",
	"pn", "pr", "pro", "ps", "pt", "pw", "py", "qa", "re", "ro", "rs", "ru",
	"rw", "sa", "sb", "sc", "sd", "se", "sg", "sh", "si", "sj", "sk", "sl", "sm",
	"sn", "so", "sr", "st", "su", "sv", "sy", "sz", "tc", "td", "tel", "tf", "tg",
	"th", "tj", "tk", "tl", "tm", "tn", "to", "tp", "tr", "tt", "tv", "tw", "tz",
	"ua", "ug", "uk", "us", "uy", "uz", "va", "vc", "ve", "vg", "vi", "vn", "vu",
	"wf", "ws", "xn", "ye", "yt", "za", "zm", "zw"}

func (req *Req) getAuthoritativeNameServer() (string, error) {
	tokenizedObject := strings.Split(req.Object, ".")
	for _, value := range TLDArray {
		if value == tokenizedObject[len(tokenizedObject)-1] {
			return parseWhoisServersNet(req.whoisRequest(value+".whois-servers.net", "="))
		}
	}
	return "whois.arin.net", nil
}

func parseWhoisServersNet(recursiveResponse string, err error) (string, error) {
	if err != nil {
		return recursiveResponse, err
	}
	foundDomain := false
	for _, value := range strings.Split(recursiveResponse, "\n") {
		if strings.Contains(value, "Domain Name: ") {
			foundDomain = true
		} else if !foundDomain {
			continue
		}

		if strings.Contains(value, "Whois Server: ") {
			return strings.Split(value, "Whois Server: ")[1], nil
		}
	}
	return "", nil

}
