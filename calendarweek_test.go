package main

import (
	"testing"
	"time"
)

func TestReturnCalendarWeekIfChanged(t *testing.T) {
	var returnValue = 1
	it := calendarWeekIteratorWithCustomProvider(10*time.Millisecond, func() int {
		return returnValue
	})

	cw := <-it.ChangedCh

	if cw != returnValue {
		t.Errorf("Return value was incorrect, got: %d, want: %d.", cw, returnValue)
	}

	timeout := time.NewTimer(100 * time.Millisecond)
	select {
	case <-it.ChangedCh:
		t.Errorf("Should not be called if not changed")
	case <-timeout.C:
		// we want the timeout to be hit, so we can validate that the function does not return an unchanged value
		break
	}

	returnValue = 2

	cw2 := <-it.ChangedCh

	if cw2 != returnValue {
		t.Errorf("Return value was incorrect, got: %d, want: %d.", cw, returnValue)
	}

}
