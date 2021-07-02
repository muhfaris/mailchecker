package domain

import "testing"

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
			if err := smtp.GmailValidate(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("SMTPValidate.GmailValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
