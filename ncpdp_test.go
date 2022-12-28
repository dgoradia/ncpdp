package ncpdp

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestNewDecoder(t *testing.T) {
	file, err := os.Open("testdata/sample-newrx.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	tests := []struct {
		name    string
		msg     io.Reader
		wantErr bool
		err     error
	}{
		{
			name:    "valid message",
			msg:     file,
			wantErr: false,
			err:     nil,
		},
		{
			name:    "invalid xml",
			msg:     strings.NewReader("aGVsbG8="),
			wantErr: true,
			err:     io.EOF,
		},
		{
			name:    "empty message",
			msg:     strings.NewReader(""),
			wantErr: true,
			err:     io.EOF,
		},
		{
			name:    "date parse error",
			msg:     strings.NewReader(`<Message><Body><NewRx><Patient><HumanPatient><DateOfBirth><Date>20219-01-01</Date></DateOfBirth></HumanPatient></Patient></NewRx></Body></Message>`),
			wantErr: true,
		},
		{
			name:    "date decode error",
			msg:     strings.NewReader(`<Message><Body><NewRx><Patient><HumanPatient><DateOfBirth><Date><>20219-01-01</Date></DateOfBirth></HumanPatient></Patient></NewRx></Body></Message>`),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewDecoder(tt.msg)
			_, err := msg.Decode()
			if tt.err == nil {
				if (err == nil) == tt.wantErr {
					t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if !errors.Is(err, tt.err) {
					t.Errorf("Decode() error = %v, wantErr %v", err, tt.err)
				}
			}

			_, err = msg.ToJson()
			if (err == nil) == tt.wantErr {
				t.Errorf("ToJson() error = %v, err %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoadTerminology(t *testing.T) {
	file, err := os.Open(baseModulePath(t) + "/NCPDPTerminology.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	tests := []struct {
		name    string
		file    io.Reader
		wantLen int
		wantErr bool
	}{
		{
			name:    "custom file",
			file:    file,
			wantLen: 544,
			wantErr: false,
		},
		{
			name:    "use embedded file",
			file:    nil,
			wantLen: 544,
			wantErr: false,
		},
		{
			name:    "invalid file data",
			file:    strings.NewReader("aGVsbG8="),
			wantLen: 0,
			wantErr: true,
		},
		{
			name:    "empty file data",
			file:    strings.NewReader(""),
			wantLen: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadTerminology(tt.file)
			if (err == nil) == tt.wantErr {
				t.Errorf("LoadTerminologyReader() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got.Len() != tt.wantLen {
				t.Errorf("LoadTerminologyReader() got = %v, want %v", got.Len(), tt.wantLen)
			}
		})
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		name  string
		terms *Terminologies
		want  int
	}{
		{
			name: "loaded terms",
			terms: &Terminologies{
				&Terminology{},
				&Terminology{},
				&Terminology{},
				&Terminology{},
				&Terminology{},
			},
			want: 5,
		},
		{
			name:  "empty terms",
			terms: &Terminologies{},
			want:  0,
		},
		{
			name: "nil terms",
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.terms.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindTermByQuantityUnitOfMeasureCode(t *testing.T) {
	terms, err := LoadTerminology(nil)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		terms   *Terminologies
		code    string
		want    string
		wantErr bool
	}{
		{
			name:    "valid code",
			terms:   terms,
			code:    "C48672",
			want:    "Schedule I Substance",
			wantErr: false,
		},
		{
			name:    "invalid code",
			terms:   terms,
			code:    "C486232",
			want:    "",
			wantErr: false,
		},
		{
			name:    "empty terms",
			terms:   &Terminologies{},
			code:    "2",
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.terms.FindTermByQuantityUnitOfMeasureCode(tt.code)
			if (err == nil) == tt.wantErr {
				t.Errorf("FindTermByQuantityUnitOfMeasureCode() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("FindTermByQuantityUnitOfMeasureCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadLoinc(t *testing.T) {
	file, err := os.Open(baseModulePath(t) + "/Loinc.csv")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	tests := []struct {
		name    string
		file    io.Reader
		wantLen int
		wantErr bool
	}{
		{
			name:    "custom file",
			file:    file,
			wantLen: 96635,
			wantErr: false,
		},
		{
			name:    "use embedded file",
			file:    nil,
			wantLen: 96635,
			wantErr: false,
		},
		{
			name:    "invalid file data",
			file:    strings.NewReader("aGVsbG8="),
			wantLen: 0,
			wantErr: false,
		},
		{
			name:    "empty file data",
			file:    strings.NewReader(""),
			wantLen: 0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadLoinc(tt.file)
			if (err == nil) == tt.wantErr {
				t.Errorf("LoadLoinc() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got.Len() != tt.wantLen {
				t.Errorf("LoadLoinc() got = %v, want %v", got.Len(), tt.wantLen)
			}
		})
	}
}

func baseModulePath(t *testing.T) string {
	t.Helper()
	_, b, _, _ := runtime.Caller(0)
	return filepath.Dir(b)
}
