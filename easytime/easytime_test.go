package easytime

import (
    "testing"
    "time"
)

func TestBaseFormatTime(t *testing.T) {
    res := BaseFormatTime()
    t.Logf("formatTime: %s", res)
}

func TestFormatTime(t *testing.T) {
    res := FormatTime("y-m-d h:i:s", time.Now())
    t.Logf("formatTime: %v", res)
}
