package domain

import (
	"net"
)

// EmailMXValidator is validate mx domain
func EmailMXValidator(domain string) ([]*net.MX, bool) {
	mxs, err := net.LookupMX(domain)
	if err != nil {
		return nil, false
	}

	return mxs, true
}
