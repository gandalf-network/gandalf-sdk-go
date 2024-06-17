package connect

import (
	"testing"
)

func TestGenerateURL(t *testing.T) {
	tests := []struct {
		name           string
		config         Config
		expectedErr    error
		validPublicKey bool
	}{
		{
			name: "Valid parameters",
			config: Config{
				PublicKey:   "0x036518f1c7a10fc77f835becc0aca9916c54505f771c82d87dd5943bb01ba5ca08",
				RedirectURL: "https://example.com/redirect",
				Data: InputData{
					"uber": Service{
						Traits:     []string{"rating"},
						Activities: []string{"trip"},
					},
				},
				Platform: PlatformTypeIOS,
			},
			expectedErr:    nil,
			validPublicKey: true,
		},
		{
			name: "Invalid public key",
			config: Config{
				PublicKey:   "invalid-public-key",
				RedirectURL: "https://example.com/redirect",
				Data: InputData{
					"uber": Service{
						Traits:     []string{"rating"},
						Activities: []string{"trip"},
					},
				},
				Platform: PlatformTypeIOS,
			},
			expectedErr:    &GandalfError{Message: "Invalid public key", Code: InvalidPublicKey},
			validPublicKey: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn, err := NewConnect(tt.config)
			if err != nil {
				t.Fatalf("NewConnect() error = %v", err)
			}

			url, err := conn.GenerateURL()
			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Fatalf("GenerateURL() error = %v, expectedErr = %v", err, tt.expectedErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("GenerateURL() error = %v, expectedErr = %v", err, tt.expectedErr)
			}
			if url == "" {
				t.Fatal("GenerateURL() returned an empty URL")
			}
		})
	}
}

func TestGenerateQRCode(t *testing.T) {
	tests := []struct {
		name           string
		config         Config
		expectedErr    error
		validPublicKey bool
	}{
		{
			name: "Valid parameters",
			config: Config{
				PublicKey:   "0x036518f1c7a10fc77f835becc0aca9916c54505f771c82d87dd5943bb01ba5ca08",
				RedirectURL: "https://example.com/redirect",
				Data: InputData{
					"uber": Service{
						Traits:     []string{"rating"},
						Activities: []string{"trip"},
					},
				},
				Platform: PlatformTypeIOS,
			},
			expectedErr:    nil,
			validPublicKey: true,
		},
		{
			name: "Invalid public key",
			config: Config{
				PublicKey:   "invalid-public-key",
				RedirectURL: "https://example.com/redirect",
				Data: InputData{
					"uber": Service{
						Traits:     []string{"rating"},
						Activities: []string{"trip"},
					},
				},
				Platform: PlatformTypeIOS,
			},
			expectedErr:    &GandalfError{Message: "Invalid public key", Code: InvalidPublicKey},
			validPublicKey: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			conn, err := NewConnect(tt.config)
			if err != nil {
				t.Fatalf("NewConnect() error = %v", err)
			}

			qrCode, err := conn.GenerateQRCode()
			if tt.expectedErr != nil {
				if err == nil || err.Error() != tt.expectedErr.Error() {
					t.Fatalf("GenerateQRCode() error = %v, expectedErr = %v", err, tt.expectedErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("GenerateQRCode() error = %v, expectedErr = %v", err, tt.expectedErr)
			}
			if qrCode == "" {
				t.Fatal("GenerateQRCode() returned an empty QR code URL")
			}
		})
	}
}
