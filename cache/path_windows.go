/*
	Copyright (C) 2014  Oscar Campos <oscar.campos@member.fsf.org>

	This program is free software; you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation; either version 2 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License along
    with this program; if not, write to the Free Software Foundation, Inc.,
    51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

	See LICENSE file for more details.
*/

package cache

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

// Return CacheDirectory on Windows, %APPDATA%\\VenGO
func CacheDirectory() string {
	APPDATA := os.Getenv("APPDATA")
	if APPDATA == "" {
		APPDATA = ExpandUser("~")
	}
	return path.Join(APPDATA, "VenGO")
}

// return back the binary string version for downloads in Microsoft Windows
func GetBinaryVersion(version string) string {
	return fmt.Sprintf("%s.windows-%s", version, runtime.GOARCH)
}
