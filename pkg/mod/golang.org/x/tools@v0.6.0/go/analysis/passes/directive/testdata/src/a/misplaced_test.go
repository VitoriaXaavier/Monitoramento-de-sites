// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:debug panicnil=1

package p_test

// want +1 `//go:debug directive only valid before package declaration`
//go:debug panicnil=1
