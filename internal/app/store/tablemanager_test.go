package store_test

import (
	"testing"
	"time"

	"github.com/egorbolychev/internal/app/store"
)

func TestTableManager_AddTable(t *testing.T) {
	st := store.New()
	if st.Table().Len() != 0 {
		t.Fail()
	}
	st.Table().AddTable(1)
	if st.Table().Len() != 1 {
		t.Fail()
	}
}

func TestTableManager_IsBusy(t *testing.T) {
	st := store.New()
	st.Table().AddTable(1)
	if st.Table().IsBusy(1) {
		t.Fail()
	}
	time := time.Now()
	st.Table().NewUser(1, "user", time)
	if !st.Table().IsBusy(1) {
		t.Fail()
	}
}

func TestTableManager_ClearTable(t *testing.T) {
	hourCost := 10
	st := store.New()
	st.Table().AddTable(1)
	startTime1, _ := time.Parse("15:04", "11:00")
	st.Table().NewUser(1, "user", startTime1)
	time1, _ := time.Parse("15:04", "13:00")
	st.Table().ClearTable(1, hourCost, time1)
	money, timeAll := st.Table().GetSum(1)
	if money != 20 {
		t.Log("wrong money sum")
		t.Fail()
	}
	resTime, _ := time.Parse("15:04", "02:00")
	if timeAll.Format("15:04") != resTime.Format("15:04") {
		t.Log("wrong allTime")
		t.Fail()
	}
	startTime2, _ := time.Parse("15:04", "17:04")
	time2, _ := time.Parse("15:04", "20:05")
	st.Table().NewUser(1, "user", startTime2)
	st.Table().ClearTable(1, hourCost, time2)
	money, timeAll = st.Table().GetSum(1)
	if money != 60 {
		t.Log("wrong money sum")
		t.Fail()
	}
	resTime, _ = time.Parse("15:04", "05:01")
	if timeAll.Format("15:04") != resTime.Format("15:04") {
		t.Log("wrong allTime")
		t.Fail()
	}
}
