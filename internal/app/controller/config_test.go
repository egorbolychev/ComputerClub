package controller_test

import (
	"testing"

	"github.com/egorbolychev/internal/app/controller"
)

func TestConfig_Configure(t *testing.T) {
	conf := controller.NewConfig()

	testCases := []struct {
		name    string
		s       []string
		isValid bool
	}{
		{
			name: "valid_test",
			s: []string{
				"3",
				"14:00 17:00",
				"12",
			},
			isValid: true,
		},
		{
			name: "wrong_maxTables",
			s: []string{
				"wrong",
				"14:00 17:00",
				"12",
			},
			isValid: false,
		},
		{
			name: "wrong_cost",
			s: []string{
				"3",
				"14:00 17:00",
				"wrong",
			},
			isValid: false,
		},
		{
			name: "wrong time",
			s: []string{
				"3",
				"time time",
				"20",
			},
			isValid: false,
		},
		{
			name: "not_enough_time_bounds",
			s: []string{
				"3",
				"14:00",
				"20",
			},
			isValid: false,
		},
		{
			name: "incorrect_sequence",
			s: []string{
				"3",
				"14:00 11:00",
				"20",
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := conf.ConfigureConfig(tc.s)
			if tc.isValid {
				if err != nil {
					t.Log(err)
					t.Fail()
				}
			} else {
				if err == nil {
					t.Log("must be wrong")
					t.Fail()
				}
			}
		})
	}
}
