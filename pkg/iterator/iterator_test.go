package iterator

import (
	"io"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {

	type args struct {
		r io.Reader
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Reader with proper json map",
			args{strings.NewReader(`{"1" : {"a":"b"}}`)},
			false,
		},
		{
			"Reader with corrupted json",
			args{strings.NewReader("]ds")},
			true,
		},
		{
			"Reader without map",
			args{strings.NewReader("[{'a':'b'}]")},
			true,
		},
	}

	for _, tt := range testCases {

		t.Run(tt.name, func(t *testing.T) {

			got, err := New(tt.args.r)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewT() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			if got == nil {
				t.Errorf("New() got nil")
			}
		})
	}
}

func TestJsonMap_NextObject(t *testing.T) {

	tk, err := New(strings.NewReader(`{"1":{"a":"b"},"2":{"c":"decoder"},"3":{"a":{"b":"c"},"decoder":1}}`))
	if err != nil {
		t.Errorf("NextObject() error = %v", err)
	}

	testCases := []struct {
		expectedVal string
		wantErr     bool
	}{
		{`{"a":"b"}`, false},
		{`{"c":"decoder"}`, false},
		{`{"a":{"b":"c"},"decoder":1}`, false},
		{`###`, true},
	}

	for _, test := range testCases {

		b, err := tk.NextObject()

		if !test.wantErr && err != nil {
			t.Errorf("NextObject() returned error %v", err)
			return
		}

		if test.wantErr && err == nil {
			t.Errorf("NextObject() did not return an error %v", err)
			return
		}

		if !test.wantErr && string(b) != test.expectedVal {
			t.Errorf("NextObject() returned bad value. Got %s, expected %v", string(b), test.expectedVal)
			return
		}
	}

	if tk.More() {
		t.Error("NextObject() returned unexpected elements")
	}
}
