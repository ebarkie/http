// Copyright (c) 2016 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlain(t *testing.T) {
	d := Data{}
	d.SetFloat("f1", 1.0)
	d.SetFloat("f2", 1.1)
	d.SetInt("i", 2)
	d.SetString("s", "v")

	assert.Equal(t, map[string]string{
		"f1": "1",
		"f2": "1.1",
		"i":  "2",
		"s":  "v",
	}, d.Values())
}
