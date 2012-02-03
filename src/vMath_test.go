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

import "testing"

// Allowed error in floating point comparisons
const ε = 1e-4

func similarScalar(a, b float32) bool {
   return a > b-ε && a < b+ε
}

// Returns true if the float32s that make up a and b are within ε of
// eachother.
func similarV2(a, b *Vec2) bool {
   rv := a[0] > b[0]-ε && a[0] < b[0]+ε &&
   a[1] > b[1]-ε && a[1] < b[1]+ε
   return rv
}

// Returns true if the float32s that make up a and b are within ε of
// eachother.
func similarV3(a, b *Vec3) bool {
   rv := a[0] > b[0]-ε && a[0] < b[0]+ε &&
   a[1] > b[1]-ε && a[1] < b[1]+ε &&
   a[2] > b[2]-ε && a[2] < b[2]+ε
   return rv
}

// Returns true if the float32s that make up a and b are within ε of
// eachother.
func similarV4(a, b *Vec4) bool {
   rv := a[0] > b[0]-ε && a[0] < b[0]+ε &&
   a[1] > b[1]-ε && a[1] < b[1]+ε &&
   a[2] > b[2]-ε && a[2] < b[2]+ε &&
   a[3] > b[3]-ε && a[3] < b[3]+ε
   return rv
}

// Test vector functions that operate on 2D vectors.
func TestVec2(t *testing.T) {
   var a, b *Vec2

   a, b = &Vec2{1.0, 2.0}, &Vec2{3.0, 4.0}

   // Test dot product with distinct operands
   res := DotV2(a, b)
   if !similarScalar(res, 11.0) {
      t.Errorf("DotV2(a, b):\n"+
      " yielded: %f\n"+
      " expected: %f\n",
      res, 11.0)
   }

   // Test dot product with a repeated operand
   res = DotV2(a, a)
   if !similarScalar(res, 5.0) {
      t.Errorf("DotV2(a, a):\n"+
      " yielded: %f\n"+
      " expected: %f\n",
      res, 5.0)
   }

   // Test multiplying a vector by a scalar
   resV2 := ScaleV2(&Vec2{}, b, 0.5)
   if !similarV2(resV2, &Vec2{1.5, 2.0}) {
      t.Errorf("ScaleV2:\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      resV2, Vec2{1.5, 2.0})
   }

   // Test multiplying a vector by a scalar, overwriting the Vec2 operand
   ScaleV2(a, a, 0.5)
   if !similarV2(a, &Vec2{0.5, 1.0}) {
      t.Errorf("ScaleV2 with operand overwrite:\n"+
      " yielded: %v:\n"+
      " expected: %v\n",
      a, Vec2{0.5, 1.0})
   }

   // Test normalizing a vector
   a = &Vec2{1.0, 2.0}
   if !similarV2(NormalizeV2(resV2, a), &Vec2{0.447213, 0.894427}) {
      t.Errorf("NormalizeV2:\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      resV2, Vec2{0.447213, 0.894427})
   }

   // Test normalizing a vector, overwriting the Vec2 operand
   if !similarV2(NormalizeV2(a, a), &Vec2{0.447213, 0.894427}) {
      t.Errorf("NormalizeV2 with operand overwrite:\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      a, Vec2{0.447213, 0.894427})
   }

   // Test vector addition
   a = &Vec2{1.0, 2.0}
   if !similarV2(AddV2(resV2, a, b), &Vec2{4.0, 6.0}) {
      t.Errorf("AddV2(r, a, b):\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      resV2, Vec2{4.0, 6.0})
   }

   // Test vector addition with a reused operand
   if !similarV2(AddV2(resV2, a, a), &Vec2{2.0, 4.0}) {
      t.Errorf("AddV2(r, a, a):\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      resV2, &Vec2{2.0, 4.0})
   }

   // Test vector addition with result stored in operand
   if !similarV2(AddV2(a, a, a), &Vec2{2.0, 4.0}) {
      t.Errorf("AddV2(a, a, a):\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      a, &Vec2{2.0, 4.0})
   }
}

// Test vector functions that operate on 2D vectors.
func TestVec3(t *testing.T) {
   var a, b *Vec3
   var resV3 = new(Vec3)

   a, b = &Vec3{1.0, 2.0, 3.0}, &Vec3{4.0, 5.0, 6.0}

   // Test cross product
   if !similarV3(CrossV3(resV3, a, b), &Vec3{-3, 6, -3}) {
      t.Errorf("CrossV3(r, a, b):\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      resV3, &Vec3{-3, 6, -3})
   }

   // Test cross product with reused operand
   if !similarV3(CrossV3(resV3, a, a), &Vec3{0, 0, 0}) {
      t.Errorf("CrossV3(r, a, a):\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      resV3, &Vec3{0, 0, 0})
   }

   // Test cross product with operand overwrite
   if !similarV3(CrossV3(a, a, b), &Vec3{-3, 6, -3}) {
      t.Errorf("CrossV3(a, a, b):\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      a, &Vec3{-3, 6, -3})
   }

   a = &Vec3{1.0, 2.0, 3.0}
   // Test dot product with distinct operands
   res := DotV3(a, b)
   if !similarScalar(res, 32.0) {
      t.Errorf("DotV3(a, b):\n"+
      " yielded: %f\n"+
      " expected: %f\n",
      res, 32.0)
   }

   // Test dot product with a repeated operand
   res = DotV3(a, a)
   if !similarScalar(res, 14.0) {
      t.Errorf("DotV3(a, a):\n"+
      " yielded: %f\n"+
      " expected: %f\n",
      res, 14.0)
   }

   // Test multiplying a vector by a scalar
   if !similarV3(ScaleV3(resV3, b, 0.5), &Vec3{2.0, 2.5, 3.0}) {
      t.Errorf("ScaleV2:\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      resV3, Vec3{2.0, 2.5, 3.0})
   }

   // Test multiplying a vector by a scalar, overwriting the Vec2 operand
   ScaleV3(a, a, 0.5)
   if !similarV3(a, &Vec3{0.5, 1.0, 1.5}) {
      t.Errorf("ScaleV2 with operand overwrite:\n"+
      " yielded: %v:\n"+
      " expected: %v\n",
      a, Vec3{0.5, 1.0, 1.5})
   }

   // Test normalizing a vector
   a = &Vec3{1.0, 2.0, 3.0}
   if !similarV3(NormalizeV3(resV3, a), &Vec3{0.267261, 0.534522, 0.801783}) {
      t.Errorf("NormalizeV3:\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      resV3, Vec3{0.267261, 0.534522, 0.801783})
   }

   // Test normalizing a vector, overwriting the Vec2 operand
   if !similarV3(NormalizeV3(a, a), &Vec3{0.267261, 0.534522, 0.801783}) {
      t.Errorf("NormalizeV3 with operand overwrite:\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      a, Vec3{0.267261, 0.534522, 0.801783})
   }

   // Test vector addition
   a = &Vec3{1.0, 2.0, 3.0}
   if !similarV3(AddV3(resV3, a, b), &Vec3{5.0, 7.0, 9.0}) {
      t.Errorf("AddV2(r, a, b):\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      resV3, Vec3{5.0, 7.0, 9.0})
   }

   // Test vector addition with a reused operand
   if !similarV3(AddV3(resV3, a, a), &Vec3{2.0, 4.0, 6.0}) {
      t.Errorf("AddV3(r, a, a):\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      resV3, &Vec3{2.0, 4.0, 6.0})
   }

   // Test vector addition with result stored in operand
   if !similarV3(AddV3(a, a, a), &Vec3{2.0, 4.0, 6.0}) {
      t.Errorf("AddV3(a, a, a):\n"+
      " yielded: %v\n"+
      " expected: %v\n",
      a, &Vec3{2.0, 4.0, 6.})
   }
}
