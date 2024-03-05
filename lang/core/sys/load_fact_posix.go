// Mgmt
// Copyright (C) 2013-2024+ James Shubin and the project contributors
// Written by James Shubin <james@shubin.ca> and the project contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//
// Additional permission under GNU GPL version 3 section 7
//
// If you modify this program, or any covered work, by linking or combining it
// with embedded mcl code and modules (and that the embedded mcl code and
// modules which link with this program, contain a copy of their source code in
// the authoritative form) containing parts covered by the terms of any other
// license, the licensors of this program grant you additional permission to
// convey the resulting work. Furthermore, the licensors of this program grant
// the original author, James Shubin, additional permission to update this
// additional permission if he deems it necessary to achieve the goals of this
// additional permission.

//go:build !darwin

package coresys

import (
	"syscall"
)

const (
	// LoadScale factor scales the output from sysinfo to the correct float
	// value.
	LoadScale = 65536 // XXX: is this correct or should it be 65535?
)

// load returns the system load averages for the last minute, five minutes and
// fifteen minutes. Calling this more often than once every five seconds seems
// to be unnecessary, since the kernel only updates these values that often.
// TODO: is the kernel update interval configurable?
func load() (one, five, fifteen float64, err error) {
	var sysinfo syscall.Sysinfo_t
	if err = syscall.Sysinfo(&sysinfo); err != nil {
		return
	}
	one = float64(sysinfo.Loads[0]) / LoadScale
	five = float64(sysinfo.Loads[1]) / LoadScale
	fifteen = float64(sysinfo.Loads[2]) / LoadScale
	return
}
