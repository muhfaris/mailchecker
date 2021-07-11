package domain

import (
	"reflect"
	"testing"
)

func TestSMTPValidate_GmailValidate(t *testing.T) {
	type fields struct {
		IsValid bool
	}
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "check valid mail",
			fields: fields{},
			args: args{
				email: "devmuhfaris@gmail",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smtp := &SMTPValidate{
				IsValid: tt.fields.IsValid,
			}
			if err := smtp.GmailValidate(); (err != nil) != tt.wantErr {
				t.Errorf("SMTPValidate.GmailValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEmailSMTPValidator(t *testing.T) {
	type args struct {
		mxs []string
		e   *EmailVerifier
	}
	tests := []struct {
		name    string
		args    args
		want    *SMTPValidate
		wantErr bool
	}{
		{
			name: "check email from mail.com",
			args: args{
				mxs: []string{"mx01.mail.com", "mx00.mail.com"},
				e: &EmailVerifier{
					Email: "muhfaris@mail.com",
				},
			},
			want: &SMTPValidate{IsValid: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EmailSMTPValidator(tt.args.mxs, tt.args.e)
			if (err != nil) != tt.wantErr {
				t.Errorf("EmailSMTPValidator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EmailSMTPValidator() = %v, want %v", got, tt.want)
			}
		})
	}
}
