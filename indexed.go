// Copyright (c) 2016-2017 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package request

import (
	"fmt"
	"strconv"
	"strings"
)

// Indexed generates an indexed key based on formatting criteria.
//
// Examples:
//
//     Weather Underground
//     tempf  First outdoor temperature
//     temp2f Second outdoor temperature
//
//     Indexed{Format: "temp#f", Begin: 1, Zero: 2}
//
//     WeatherCloud
//     temp01 First outdoor temperature
//     temp02 Second outdoor temperature
//
//     Indexed{Format: "temp#", Begin: 1, Width: 2}
type Indexed struct {
	Format string // Template where the sharp symbol(#) will be substituted for the number
	Begin  int    // Begin index number
	Width  int    // Index number width
	Zero   int    // Any index number below the zero value not be printed (it's implied)
}

// Generate generates a key using the formatting criteria and finding
// the lowest a vailable index.
func (i Indexed) Generate(d Data) (key string) {
	for j := i.Begin; ; j++ {
		if j <= i.Zero {
			key = strings.Replace(i.Format, "#", "", 1)
		} else {
			key = strings.Replace(i.Format, "#", fmt.Sprintf("%0"+strconv.Itoa(i.Width)+"d", j), 1)
		}
		if !d.Exists(key) {
			return
		}
	}
}