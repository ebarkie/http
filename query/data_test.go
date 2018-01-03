// Copyright (c) 2016 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testValues(t *testing.T, d *Data) {
	a := assert.New(t)

	a.Nil(d.Values())

	d.SetString("k", "v")
	a.Equal(map[string]string{"k": "v"}, d.Values())
}

func TestClear(t *testing.T) {
	d := Data{}
	testValues(t, &d)

	d.Clear()
	assert.Equal(t, map[string]string{}, d.Values())
}

func TestExists(t *testing.T) {
	a := assert.New(t)

	d := Data{}

	a.False(d.Exists("k"))

	d.SetString("k", "v")
	a.True(d.Exists("k"))
}

func TestValues(t *testing.T) {
	d := Data{}
	testValues(t, &d)
}
