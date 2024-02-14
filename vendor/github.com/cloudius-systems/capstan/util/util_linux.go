/*
 * Copyright (C) 2014 Cloudius Systems, Ltd.
 *
 * This work is open source software, licensed under the terms of the
 * BSD license as described in the LICENSE file in the top-level directory.
 */

package util

import (
	"os"
	"syscall"
)

func IsDirectIOSupported(path string) bool {
	f, err := os.OpenFile(path, syscall.O_DIRECT, 0)
	defer f.Close()
	return err == nil
}
