package tool

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestDo(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockmyInterface(ctrl)

	m.
		EXPECT().
		Int().
		Return(101)

	if Do(m) != 101 {
		t.Fail()
	}
}
