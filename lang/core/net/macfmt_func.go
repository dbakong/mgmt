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

package corenet

import (
	"fmt"
	"net"
	"strings"

	"github.com/purpleidea/mgmt/lang/funcs/simple"
	"github.com/purpleidea/mgmt/lang/types"
)

func init() {
	simple.ModuleRegister(ModuleName, "macfmt", &types.FuncValue{
		T: types.NewType("func(a str) str"),
		V: MacFmt,
	})
	simple.ModuleRegister(ModuleName, "oldmacfmt", &types.FuncValue{
		T: types.NewType("func(a str) str"),
		V: OldMacFmt,
	})
}

// MacFmt takes a MAC address with hyphens and converts it to a format with
// colons.
func MacFmt(input []types.Value) (types.Value, error) {
	mac := input[0].Str()

	// Check if the MAC address is valid.
	if len(mac) != len("00:00:00:00:00:00") {
		return nil, fmt.Errorf("invalid MAC address length: %s", mac)
	}
	_, err := net.ParseMAC(mac)
	if err != nil {
		return nil, err
	}

	return &types.StrValue{
		V: strings.Replace(mac, "-", ":", -1),
	}, nil
}

// OldMacFmt takes a MAC address with colons and converts it to a format with
// hyphens. This is the old deprecated style that nobody likes.
func OldMacFmt(input []types.Value) (types.Value, error) {
	mac := input[0].Str()

	// Check if the MAC address is valid.
	if len(mac) != len("00:00:00:00:00:00") {
		return nil, fmt.Errorf("invalid MAC address length: %s", mac)
	}
	_, err := net.ParseMAC(mac)
	if err != nil {
		return nil, err
	}

	return &types.StrValue{
		V: strings.Replace(mac, ":", "-", -1),
	}, nil
}
