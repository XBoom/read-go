// errorcheck

// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package p

var x any // ERROR "undefined: any|undefined type .*any.*|cannot use any outside constraint position"
