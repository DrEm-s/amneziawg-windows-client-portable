/* SPDX-License-Identifier: MIT
 *
 * Portable mode additions.
 */

package main

import (
	"os"
	"path/filepath"

	"github.com/amnezia-vpn/amneziawg-windows/v3/conf"
)

func init() {
	executable, err := os.Executable()
	if err != nil {
		panic(err)
	}

	dataDir := filepath.Join(filepath.Dir(executable), "Data")
	if err := os.MkdirAll(dataDir, 0700); err != nil {
		panic(err)
	}

	conf.PresetRootDirectory(dataDir)
}
