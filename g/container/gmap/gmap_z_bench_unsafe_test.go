// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/gogf/gf.

// go test *.go -bench=".*" -benchmem

package gmap

import (
    "testing"
    "strconv"
)


var ibmUnsafe   = NewIntBoolMap(true)
var iimUnsafe   = NewIntIntMap(true)
var iifmUnsafe  = NewIntInterfaceMap(true)
var ismUnsafe   = NewIntStringMap(true)
var ififmUnsafe = NewMap(true)
var sbmUnsafe   = NewStringBoolMap(true)
var simUnsafe   = NewStringIntMap(true)
var sifmUnsafe  = NewStringInterfaceMap(true)
var ssmUnsafe   = NewStringStringMap(true)

// 写入性能测试

func Benchmark_Unsafe_IntBoolMap_Set(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ibmUnsafe.Set(i, true)
    }
}

func Benchmark_Unsafe_IntIntMap_Set(b *testing.B) {
    for i := 0; i < b.N; i++ {
        iimUnsafe.Set(i, i)
    }
}

func Benchmark_Unsafe_IntInterfaceMap_Set(b *testing.B) {
    for i := 0; i < b.N; i++ {
        iifmUnsafe.Set(i, i)
    }
}

func Benchmark_Unsafe_IntStringMap_Set(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ismUnsafe.Set(i, strconv.Itoa(i))
    }
}

func Benchmark_Unsafe_InterfaceInterfaceMap_Set(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ififmUnsafe.Set(i, i)
    }
}

func Benchmark_Unsafe_StringBoolMap_Set(b *testing.B) {
    for i := 0; i < b.N; i++ {
        sbmUnsafe.Set(strconv.Itoa(i), true)
    }
}

func Benchmark_Unsafe_StringIntMap_Set(b *testing.B) {
    for i := 0; i < b.N; i++ {
        simUnsafe.Set(strconv.Itoa(i), i)
    }
}

func Benchmark_Unsafe_StringInterfaceMap_Set(b *testing.B) {
    for i := 0; i < b.N; i++ {
        sifmUnsafe.Set(strconv.Itoa(i), i)
    }
}

func Benchmark_Unsafe_StringStringMap_Set(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ssmUnsafe.Set(strconv.Itoa(i), strconv.Itoa(i))
    }
}


// 读取性能测试

func Benchmark_Unsafe_IntBoolMap_Get(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ibmUnsafe.Get(i)
    }
}

func Benchmark_Unsafe_IntIntMap_Get(b *testing.B) {
    for i := 0; i < b.N; i++ {
        iimUnsafe.Get(i)
    }
}

func Benchmark_Unsafe_IntInterfaceMap_Get(b *testing.B) {
    for i := 0; i < b.N; i++ {
        iifmUnsafe.Get(i)
    }
}

func Benchmark_Unsafe_IntStringMap_Get(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ismUnsafe.Get(i)
    }
}

func Benchmark_Unsafe_InterfaceInterfaceMap_Get(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ififmUnsafe.Get(i)
    }
}

func Benchmark_Unsafe_StringBoolMap_Get(b *testing.B) {
    for i := 0; i < b.N; i++ {
        sbmUnsafe.Get(strconv.Itoa(i))
    }
}

func Benchmark_Unsafe_StringIntMap_Get(b *testing.B) {
    for i := 0; i < b.N; i++ {
        simUnsafe.Get(strconv.Itoa(i))
    }
}

func Benchmark_Unsafe_StringInterfaceMap_Get(b *testing.B) {
    for i := 0; i < b.N; i++ {
        sifmUnsafe.Get(strconv.Itoa(i))
    }
}

func Benchmark_Unsafe_StringStringMap_Get(b *testing.B) {
    for i := 0; i < b.N; i++ {
        ssmUnsafe.Get(strconv.Itoa(i))
    }
}

