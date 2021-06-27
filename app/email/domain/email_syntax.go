package domain

import (
	"fmt"
	"regexp"
	"strings"
)

const emailRXFormat string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"

var (
	userRegexp    = regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$")
	userDotRegexp = regexp.MustCompile("(^[.]{1})|([.]{1}$)|([.]{2,})")
	rxEmail       = regexp.MustCompile(emailRXFormat)
	hostRegexp    = regexp.MustCompile("^[^\\s]+\\.[^\\s]+$")
)

// EmailSyntaxValidate is wrap data email syntax
type EmailSyntaxValidate struct {
	User   string
	Domain string
}

// EmailSyntaxValidator is validate email syntax
func EmailSyntaxValidator(email string) (EmailSyntaxValidate, error) {
	if email == "" {
		return EmailSyntaxValidate{}, fmt.Errorf(ErrorMissingParam, "email")
	}

	if len(email) < 6 || len(email) > 254 {
		return EmailSyntaxValidate{}, fmt.Errorf(ErrorInvalidParam, "email")
	}

	at := strings.LastIndex(email, "@")
	if at <= 0 || at > len(email)-3 {
		return EmailSyntaxValidate{}, fmt.Errorf(ErrorInvalidParam, "email")
	}

	user := email[:at]
	host := email[at+1:]
	if len(user) > 64 {
		return EmailSyntaxValidate{}, fmt.Errorf(ErrorInvalidParam, "email")
	}

	if userDotRegexp.MatchString(user) || !userRegexp.MatchString(user) || !hostRegexp.MatchString(host) {
		return EmailSyntaxValidate{}, fmt.Errorf(ErrorInvalidParam, "email")
	}

	return EmailSyntaxValidate{
		User:   user,
		Domain: host,
	}, nil
}
