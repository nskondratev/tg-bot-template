package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.NoError(t, os.Setenv("TEST_STR_KEY", "test_val"))

	type args struct {
		key          string
		defaultValue string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Default value",
			args: args{
				key:          "NOTEXISTENT",
				defaultValue: "default_val",
			},
			want: "default_val",
		},
		{
			name: "Env value",
			args: args{
				key:          "TEST_STR_KEY",
				defaultValue: "test_val",
			},
			want: "test_val",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := String(tt.args.key, tt.args.defaultValue)

			assert.Equal(t, tt.want, got)
		})
	}
}
