// Copyright (c) 2016 Eric Barkie. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package query

import (
	"fmt"
	"strconv"
	"strings"
)

// Indexed generates an indexed key based on formatting criteria.
//
// Examples:
//
//     WeatherCloud
//     temp   First outdoor temperature
//     temp02 Second outdoor temperature
//
//     Indexed{Format: "temp#", Begin: 1, Zero: 1, Width: 2}
//
//     Weather Underground
//     tempf  First outdoor temperature
//     temp2f Second outdoor temperature
//
//     Indexed{Format: "temp#f", Begin: 1, Zero: 1}
type Indexed struct {
	Format string // Template where the sharp symbol(#) will be substituted for the number
	Begin  int    // Begin index number
	Width  int    // Index number width
	Zero   int    // Any index number below the zero value not be printed (it's implied)
}

// Generate generates a key using the formatting criteria and finding
// the lowest a vailable index.
func (i Indexed) Generate(v Values) (key string) {
	for j := i.Begin; ; j++ {
		if j <= i.Zero {
			key = strings.Replace(i.Format, "#", "", 1)
		} else {
			key = strings.Replace(i.Format, "#", fmt.Sprintf("%0"+strconv.Itoa(i.Width)+"d", j), 1)
		}
		if !v.Exists(key) {
			return
		}
	}
}
