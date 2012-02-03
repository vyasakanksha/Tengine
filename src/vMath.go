/****************************************************************************
* *
* This file is part of the Go-Asteria project. *
* Copyright (C) 2012 Samuel C. Payson *
* *
* Go-Asteria is free software: you can redistribute it and/or modify it under *
* the terms of the GNU General Public License as published by the Free *
* Software Foundation, either version 3 of the License, or (at your *
* option) any later version. *
* *
* Go-Asteria is distributed in the hope that it will be useful, but WITHOUT *
* ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or *
* FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License *
* for more details. *
* *
* You should have received a copy of the GNU General Public License along *
* with Go-Asteria. If not, see <http://www.gnu.org/licenses/>. *
* *
****************************************************************************/

package vmath

import (
"fmt"
"math"
"unsafe"
)

type Vec2 [2]float32

func (me Vec2) String() string {
   return fmt.Sprintf("❬%f, %f❭", me[0], me[1])
}

type Vec3 [3]float32

func (me Vec3) String() string {
   return fmt.Sprintf("❬%f, %f, %f❭", me[0], me[1], me[2])
}

type Vec4 [4]float32

func (me Vec4) String() string {
   return fmt.Sprintf("❬%f, %f, %f, %f❭", me[0], me[1], me[2], me[3])
}

func CrossV3(r, a, b *Vec3) *Vec3 {
   *r = Vec3{
      0: a[1]*b[2] - a[2]*b[1],
      1: a[2]*b[0] - a[0]*b[2],
      2: a[0]*b[1] - a[1]*b[0]}
      return r
}

func DotV2(a, b *Vec2) float32 {
   return a[0]*b[0] + a[1]*b[1]
}

func DotV3(a, b *Vec3) float32 {
   return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func DotV4(a, b *Vec4) float32 {
   return a[0]*b[0] + a[1]*b[1] + a[2]*b[2] + a[3]*b[3]
}

func ScaleV2(r, v *Vec2, s float32) *Vec2 {
   *r = Vec2{
      0: v[0] * s,
      1: v[1] * s}
      return r
}

func ScaleV3(r, v *Vec3, s float32) *Vec3 {
   *r = Vec3{
      0: v[0] * s,
      1: v[1] * s,
      2: v[2] * s}
      return r
}

func ScaleV4(r, v *Vec4, s float32) *Vec4 {
   *r = Vec4{
      0: v[0] * s,
      1: v[1] * s,
      2: v[2] * s,
      3: v[3] * s}
      return r
}

func NormalizeV2(r, v *Vec2) *Vec2 {
   ScaleV2(r, v, 1/float32(math.Sqrt(float64(DotV2(v, v)))))
   return r
}

func NormalizeV3(r, v *Vec3) *Vec3 {
   ScaleV3(r, v, 1/float32(math.Sqrt(float64(DotV3(v, v)))))
   return r
}

func NormalizeV4(r, v *Vec4) *Vec4 {
   ScaleV4(r, v, 1/float32(math.Sqrt(float64(DotV4(v, v)))))
   return r
}

func AddV2(r, a, b *Vec2) *Vec2 {
   *r = Vec2{
      0: a[0] + b[0],
      1: a[1] + b[1]}
      return r
}

func AddV3(r, a, b *Vec3) *Vec3 {
   *r = Vec3{
      0: a[0] + b[0],
      1: a[1] + b[1],
      2: a[2] + b[2]}
      return r
}

func AddV4(r, a, b *Vec4) *Vec4 {
   *r = Vec4{
      0: a[0] + b[0],
      1: a[1] + b[1],
      2: a[2] + b[2],
      3: a[3] + b[3]}
      return r
}

func SubtractV3(r, a, b *Vec3) *Vec3 {
   *r = Vec3{
      0: a[0] - b[0],
      1: a[1] - b[1],
      2: a[2] - b[2]}
      return r
}

// Performs linear interpolation of two 3D vectors.
func LerpV3(r, a, b *Vec3, t float32) *Vec3 {
   var x, y Vec3
   ScaleV3(&x, a, 1.0-t)
   ScaleV3(&y, b, t)
   return AddV3(r, &x, &y)
}

// Multiply two quaternions
func MulQt(r, a, b *Vec4) *Vec4 {
   *r = Vec4{
      0: a[0]*b[3] + a[3]*b[0] + a[1]*b[2] - a[2]*b[1],
      1: a[1]*b[3] + a[3]*b[1] + a[2]*b[0] - a[0]*b[2],
      2: a[2]*b[3] + a[3]*b[2] + a[0]*b[1] - a[1]*b[0],
      3: a[3]*b[3] - a[0]*b[0] - a[1]*b[1] - a[2]*b[2]}

      return r
}

func ConjugateQt(r, q *Vec4) *Vec4 {
   *r = Vec4{
      0: -q[0],
      1: -q[1],
      2: -q[2],
      3: q[3]}
      return r
}

func RotateQt(r, v *Vec3, q *Vec4) *Vec3 {
   // So that we don't have to copy data to use it as a Vec3
   qVec := (*Vec3)(unsafe.Pointer(q))

   dot := -DotV3(qVec, v)
   conj := ScaleV3(new(Vec3), qVec, -1.0)
   var intVec = new(Vec3)
   AddV3(intVec, ScaleV3(intVec, v, q[3]), CrossV3(new(Vec3), qVec, v))

   AddV3(r, ScaleV3(r, intVec, q[3]), ScaleV3(new(Vec3), conj, dot))

   return AddV3(r, r, CrossV3(intVec, intVec, conj))
}

func NlerpQt(r, a, b *Vec4, t float32) *Vec4 {
   var x, y, v Vec4
   var p *Vec4

   if DotV4(a, b) < 0.0 {
      p = ScaleV4(&v, a, -1.0)
   } else {
      p = a
   }

   ScaleV4(&x, p, 1.0-t)
   ScaleV4(&y, b, t)
   return NormalizeV4(r, AddV4(r, &x, &y))
}
