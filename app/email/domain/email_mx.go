package domain

import (
	"fmt"
	"net"
)

// MXValidate is wrap mx record
type MXValidate struct {
	Records       []string `json:"records"`
	IsMXRecod     bool     `json:"is_mx_recod"`
	IsCNAMERecord bool     `json:"is_cname_record"`
	IsARecord     bool     `json:"is_a_record"`
	IsValid       bool     `json:"is_valid"`
	Message       string   `json:"message"`
}

// EmailMXValidator is validate mx domain
// reference to RFC7505
func EmailMXValidator(domain string) (*MXValidate, error) {
	var mx = &MXValidate{}
	err := mx.getMXRecords(domain)
	if err != nil {
		return mx, err
	}

	mx.IsValid = true
	return mx, nil
}

func (mx *MXValidate) getMXRecords(domain string) error {
	mxs, err := net.LookupMX(domain)
	if err != nil {
		return err
	}

	var mxss []string
	for _, mx := range mxs {
		mxss = append(mxss, mx.Host)
	}

	mx.Records = mxss
	mx.IsMXRecod = true
	return nil
}

// ChangeMessage is is change message mx record
func (mx *MXValidate) ChangeMessage(message string) error {
	if message == "" {
		return fmt.Errorf(ErrorMissingParam, "message mx record")
	}

	mx.Message = message
	return nil
}

// ChangeMessageF is is change message mx record
func (mx *MXValidate) ChangeMessageF(message string, err error) error {
	if message == "" {
		return fmt.Errorf(ErrorMissingParam, "message mx record")
	}

	mx.Message = fmt.Sprintf("%s, %v", message, err)
	return nil
}
