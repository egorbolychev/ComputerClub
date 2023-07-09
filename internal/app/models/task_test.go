package models_test

import (
	"testing"

	"github.com/egorbolychev/internal/app/models"
)

func TestConfig_Configure(t *testing.T) {

	testCases := []struct {
		name    string
		s       string
		isValid bool
	}{
		{
			name:    "valid_test",
			s:       "19:40 1 clien1",
			isValid: true,
		},
		{
			name:    "valid_test_4",
			s:       "19:40 1 clien1 2",
			isValid: true,
		},
		{
			name:    "too_many_args",
			s:       "19:40 1 client1 a a",
			isValid: false,
		},
		{
			name:    "wrong_time",
			s:       "wrong 1 client1",
			isValid: false,
		},
		{
			name:    "wrong_id",
			s:       "19:40 wrong client1",
			isValid: false,
		},
		{
			name:    "wrong_username",
			s:       "19:40 1 qwe?[]",
			isValid: false,
		},
		{
			name:    "wrong_tablenum",
			s:       "19:40 1 client1 qweqw",
			isValid: false,
		},
		{
			name:    "wrong_task_id",
			s:       "19:40 15 client1",
			isValid: false,
		},
		{
			name:    "wrong_task_with_id=2",
			s:       "19:40 2 client1",
			isValid: false,
		},
		{
			name:    "wrong_table_num_is_bigger_than_tables_amount",
			s:       "19:40 2 client1 4",
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			task := models.NewTask()
			err := task.ParseTask(tc.s, 3)
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
