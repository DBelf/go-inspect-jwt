package inspectjwt

import (
	"testing"
)

func Test_appEnv_run(t *testing.T) {
	type fields struct {
		jwt string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"empty jwt parse",
			fields{jwt: ""},
			true,
		},
		{"invalid jwt parse",
			fields{jwt: "a2c29f07da06ea3b56e8cc3af631fe57100c1010"},
			true,
		},
		{"invalid jwt parse",
			fields{jwt: "random words"},
			true,
		},
		{"example jwt parse",
			fields{jwt: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &appEnv{
				jwt: tt.fields.jwt,
			}
			if err := app.run(); (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
