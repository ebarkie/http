package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexedWC(t *testing.T) {
	d := Data{}
	d.SetIntf(Indexed{Format: "hum#", Begin: 1, Zero: 1, Width: 2}, 1)
	d.SetIntf(Indexed{Format: "hum#", Begin: 1, Zero: 1, Width: 2}, 2)
	assert.Equal(t, map[string]string{
		"hum":   "1",
		"hum02": "2",
	}, d.Values())
}

func TestIndexedWU(t *testing.T) {
	d := Data{}
	d.SetFloatf(Indexed{Format: "temp#f", Begin: 1, Zero: 1}, 1.0)
	d.SetFloatf(Indexed{Format: "temp#f", Begin: 1, Zero: 1}, 2.0)
	assert.Equal(t, map[string]string{
		"tempf":  "1",
		"temp2f": "2",
	}, d.Values())
}
