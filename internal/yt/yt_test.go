package yt

import "testing"

func Test_YTDLPSetup(t *testing.T) {
	_, err := NewYtClient("firefox")
	if err != nil {
		t.Error(err)
	}
}
