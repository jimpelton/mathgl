// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl64

import "golang.org/x/image/math/f64"

type Vec2 f64.Vec2
type Vec3 f64.Vec3
type Vec4 f64.Vec4

func (v Vec2) Vec3(z float64) Vec3 {
	return Vec3{v[0], v[1], z}
}

func (v Vec2) Vec4(z, w float64) Vec4 {
	return Vec4{v[0], v[1], z, w}
}

func (v Vec3) Vec2() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec3) Vec4(w float64) Vec4 {
	return Vec4{v[0], v[1], v[2], w}
}

func (v Vec4) Vec2() Vec2 {
	return Vec2{v[0], v[1]}
}

func (v Vec4) Vec3() Vec3 {
	return Vec3{v[0], v[1], v[2]}
}

// extracts the elements of the vector for direct value assignment
func (v Vec2) Elem() (x, y float64) {
	return v[0], v[1]
}

// extracts the elements of the vector for direct value assignment
func (v Vec3) Elem() (x, y, z float64) {
	return v[0], v[1], v[2]
}

// extracts the elements of the vector for direct value assignment
func (v Vec4) Elem() (x, y, z, w float64) {
	return v[0], v[1], v[2], v[3]
}
