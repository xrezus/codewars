package ip_validation

import (
	"regexp"
)

func Is_valid_ip(ip string) bool {
	matched, _ := regexp.MatchString(`^((25[0-5]|(2[0-4]|1\d|[1-9]|)\d)\.?\b){4}$`, ip)

	return matched
}
