package encryption

import (
	"reflect"
	"testing"
)

var (
	okKey = []byte("8dHWTNSAsGaaD7JbqVubF1aWVWGJYF3q")
	okIV  = []byte("M7Z4es7yWRcduU3m")
)

func TestEncryptAES(t *testing.T) {
	type args struct {
		key  []byte
		iv   []byte
		text string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "should pass if len(text) is less than aes.BlockSize",
			args:    args{key: okKey, iv: okIV, text: "Test"},
			want:    "6e99f097fb170326dbb136d2e518548c",
			wantErr: false,
		},
		{
			name:    "should pass if len(text) is greater than aes.BlockSize",
			args:    args{key: okKey, iv: okIV, text: "Larger message to test AES"},
			want:    "0bf017e3f6ae6628dec4ca256e1557300a76b6eefa7890315cbab0a0b59a492d",
			wantErr: false,
		},
		{
			name:    "should pass for text with newline symbol",
			args:    args{key: okKey, iv: okIV, text: "First \n Second \n"},
			want:    "5eebb94ad06e3c32aae78b4bf49fd3e3a6f724edbc4e973fad6f4ed6718dea13",
			wantErr: false,
		},
		{
			name:    "should encode emoji",
			args:    args{key: okKey, iv: okIV, text: "🤡 meh 🪲"},
			want:    "ff23a4cef8a4e15e1806f991594f1b80",
			wantErr: false,
		},
		{
			name:    "should encode empty string",
			args:    args{key: okKey, iv: okIV, text: ""},
			want:    "338ab1de2afc10c5e880b5db6fd6e2e2",
			wantErr: false,
		},
		{
			name:    "should return error for incorrect key",
			args:    args{key: []byte("small"), iv: okIV, text: ""},
			want:    "",
			wantErr: true,
		},
		{
			name:    "should return error for incorrect iv",
			args:    args{key: okKey, iv: []byte("small"), text: ""},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncryptAES(tt.args.key, tt.args.iv, tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptAES() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncryptAES() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecryptAES(t *testing.T) {
	type args struct {
		key        []byte
		iv         []byte
		cipherText string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "should decode small text",
			args:    args{key: okKey, iv: okIV, cipherText: "6e99f097fb170326dbb136d2e518548c"},
			want:    "Test",
			wantErr: false,
		},
		{
			name:    "should decode larger text with spaces",
			args:    args{key: okKey, iv: okIV, cipherText: "0bf017e3f6ae6628dec4ca256e1557300a76b6eefa7890315cbab0a0b59a492d"},
			want:    "Larger message to test AES",
			wantErr: false,
		},
		{
			name:    "should decode text with newline symbols",
			args:    args{key: okKey, iv: okIV, cipherText: "5eebb94ad06e3c32aae78b4bf49fd3e3a6f724edbc4e973fad6f4ed6718dea13"},
			want:    "First \n Second \n",
			wantErr: false,
		},
		{
			name:    "should decode emoji",
			args:    args{key: okKey, iv: okIV, cipherText: "ff23a4cef8a4e15e1806f991594f1b80"},
			want:    "🤡 meh 🪲",
			wantErr: false,
		},
		{
			name:    "should return error for incorrect key",
			args:    args{key: []byte("small"), iv: okIV, cipherText: ""},
			want:    "",
			wantErr: true,
		},
		{
			name:    "should return error for incorrect iv",
			args:    args{key: okKey, iv: []byte("small"), cipherText: ""},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecryptAES(tt.args.key, tt.args.iv, tt.args.cipherText)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptAES() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecryptAES() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPKCS5Padding(t *testing.T) {
	type args struct {
		src   []byte
		after int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "should add padding",
			args: args{src: []byte{115, 109, 97, 108, 108}, after: 5},
			want: []byte{115, 109, 97, 108, 108, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PKCS5Padding(tt.args.src, tt.args.after); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PKCS5Padding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPKCS5UnPadding(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "should remove padding",
			args: args{src: []byte{115, 109, 97, 108, 108, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11}},
			want: []byte{115, 109, 97, 108, 108},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PKCS5UnPadding(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PKCS5UnPadding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateRandomIV(t *testing.T) {
	type args struct {
		length uint
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "should generate random bytes",
			args:    args{length: 16},
			want:    16,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateRandomIV(tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateRandomIV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("GenerateRandomIV() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
