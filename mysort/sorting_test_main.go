package main

import (
	"fmt"

	"github.com/user/mysort"
)

func main() {
	var (
		boolSlice       = []bool{true, false, false, false, true, true, false, true, false}
		intSlice        = []int{-2, 3, -7, -8, 4, 7, 8, -5, 10, 0, 8, 9, -1, 5, 2, -3, -10, 1, -6, 6, -9, -4}
		int8Slice       = []int8{-2, 3, -7, -8, 4, 7, 8, -5, 10, 0, 8, 9, -1, 5, 2, -3, -10, 1, -6, 6, -9, -4}
		int16Slice      = []int16{-2, 3, -7, -8, 4, 7, 8, -5, 10, 0, 8, 9, -1, 5, 2, -3, -10, 1, -6, 6, -9, -4}
		int32Slice      = []int32{-2, 3, -7, -8, 4, 7, 8, -5, 10, 0, 8, 9, -1, 5, 2, -3, -10, 1, -6, 6, -9, -4}
		int64Slice      = []int64{-2, 3, -7, -8, 4, 7, 8, -5, 10, 0, 8, 9, -1, 5, 2, -3, -10, 1, -6, 6, -9, -4}
		uintSlice       = []uint{3, 4, 7, 8, 10, 0, 8, 9, 5, 2, 1, 6}
		uint8Slice      = []uint8{3, 4, 7, 8, 10, 0, 8, 9, 5, 2, 1, 6}
		uint16Slice     = []uint16{3, 4, 7, 8, 10, 0, 8, 9, 5, 2, 1, 6}
		uint32Slice     = []uint32{3, 4, 7, 8, 10, 0, 8, 9, 5, 2, 1, 6}
		uint64Slice     = []uint64{3, 4, 7, 8, 10, 0, 8, 9, 5, 2, 1, 6}
		uintptrSlice    = []uintptr{3, 4, 7, 8, 10, 0, 8, 9, 5, 2, 1, 6}
		byteSlice       = []byte{'.', 'c', 'a', 'b', '!'}
		runeSlice       = []rune{'.', 'c', 'a', 'b', '!'}
		float32Slice    = []float32{3.0, 4.0, 7.0, 8.0, 10.0, 0.0, 8.0, 9.0, 5.0, 2.0, 1.0, 6.0, 3.5, 4.5, 7.5, 8.5, 10.5, 0.5, 8.5, 9.5, 5.5, 2.5, 1.5, 6.5}
		float64Slice    = []float64{3.0, 4.0, 7.0, 8.0, 10.0, 0.0, 8.0, 9.0, 5.0, 2.0, 1.0, 6.0, 3.5, 4.5, 7.5, 8.5, 10.5, 0.5, 8.5, 9.5, 5.5, 2.5, 1.5, 6.5}
		complex64Slice  = []complex64{1 + 2i, 1 + 1i, 2 + 3i, 3 + 3i, 2 + 2i, 1 + 3i, 2 + 4i, 3 + 1i, 3 + 2i, 2 + 1i}
		complex128Slice = []complex128{1 + 2i, 1 + 1i, 2 + 3i, 3 + 3i, 2 + 2i, 1 + 3i, 2 + 4i, 3 + 1i, 3 + 2i, 2 + 1i}
		stringSlice     = []string{"bbbb", "aa", "..", "cccc", "aaa", "!!", "cc", "c", "...", "a", "!", ".", "!!!", "bb", "bbb", "!!!!", "ccc", "aaaa", "....", "b"}
	)

	boolBackup := boolSlice
	fmt.Printf("boolSlice:\n%v - orig\n", boolSlice)
	mysort.SortSlice(boolSlice, true)
	fmt.Printf("%v - asc\n", boolSlice)
	boolSlice = boolBackup
	mysort.SortSlice(boolSlice, false)
	fmt.Printf("%v - desc\n\n", boolSlice)

	intBackup := intSlice
	fmt.Printf("intSlice:\n%v - orig\n", intSlice)
	mysort.SortSlice(intSlice, true)
	fmt.Printf("%v - asc\n", intSlice)
	intSlice = intBackup
	mysort.SortSlice(intSlice, false)
	fmt.Printf("%v - desc\n\n", intSlice)

	int8Backup := int8Slice
	fmt.Printf("int8Slice:\n%v - orig\n", int8Slice)
	mysort.SortSlice(int8Slice, true)
	fmt.Printf("%v - asc\n", int8Slice)
	int8Slice = int8Backup
	mysort.SortSlice(int8Slice, false)
	fmt.Printf("%v - desc\n\n", int8Slice)

	int16Backup := int16Slice
	fmt.Printf("int16Slice:\n%v - orig\n", int16Slice)
	mysort.SortSlice(int16Slice, true)
	fmt.Printf("%v - asc\n", int16Slice)
	int16Slice = int16Backup
	mysort.SortSlice(int16Slice, false)
	fmt.Printf("%v - desc\n\n", int16Slice)

	int32Backup := int32Slice
	fmt.Printf("int32Slice:\n%v - orig\n", int32Slice)
	mysort.SortSlice(int32Slice, true)
	fmt.Printf("%v - asc\n", int32Slice)
	int32Slice = int32Backup
	mysort.SortSlice(int32Slice, false)
	fmt.Printf("%v - desc\n\n", int32Slice)

	int64Backup := int64Slice
	fmt.Printf("int64Slice:\n%v - orig\n", int64Slice)
	mysort.SortSlice(int64Slice, true)
	fmt.Printf("%v - asc\n", int64Slice)
	int64Slice = int64Backup
	mysort.SortSlice(int64Slice, false)
	fmt.Printf("%v - desc\n\n", int64Slice)

	uintBackup := uintSlice
	fmt.Printf("uintSlice:\n%v - orig\n", uintSlice)
	mysort.SortSlice(uintSlice, true)
	fmt.Printf("%v - asc\n", uintSlice)
	uintSlice = uintBackup
	mysort.SortSlice(uintSlice, false)
	fmt.Printf("%v - desc\n\n", uintSlice)

	uint8Backup := uint8Slice
	fmt.Printf("uint8Slice:\n%v - orig\n", uint8Slice)
	mysort.SortSlice(uint8Slice, true)
	fmt.Printf("%v - asc\n", uint8Slice)
	uint8Slice = uint8Backup
	mysort.SortSlice(uint8Slice, false)
	fmt.Printf("%v - desc\n\n", uint8Slice)

	uint16Backup := uint16Slice
	fmt.Printf("uint16Slice:\n%v - orig\n", uint16Slice)
	mysort.SortSlice(uint16Slice, true)
	fmt.Printf("%v - asc\n", uint16Slice)
	uint16Slice = uint16Backup
	mysort.SortSlice(uint16Slice, false)
	fmt.Printf("%v - desc\n\n", uint16Slice)

	uint32Backup := uint32Slice
	fmt.Printf("uint32Slice:\n%v - orig\n", uint32Slice)
	mysort.SortSlice(uint32Slice, true)
	fmt.Printf("%v - asc\n", uint32Slice)
	uint32Slice = uint32Backup
	mysort.SortSlice(uint32Slice, false)
	fmt.Printf("%v - desc\n\n", uint32Slice)

	uint64Backup := uint64Slice
	fmt.Printf("uint64Slice:\n%v - orig\n", uint64Slice)
	mysort.SortSlice(uint64Slice, true)
	fmt.Printf("%v - asc\n", uint64Slice)
	uint64Slice = uint64Backup
	mysort.SortSlice(uint64Slice, false)
	fmt.Printf("%v - desc\n\n", uint64Slice)

	uintptrBackup := uintptrSlice
	fmt.Printf("uintptrSlice:\n%v - orig\n", uintptrSlice)
	mysort.SortSlice(uintptrSlice, true)
	fmt.Printf("%v - asc\n", uintptrSlice)
	uintptrSlice = uintptrBackup
	mysort.SortSlice(uintptrSlice, false)
	fmt.Printf("%v - desc\n\n", uintptrSlice)

	byteBackup := byteSlice
	fmt.Printf("byteSlice:\n%v - orig\n", byteSlice)
	mysort.SortSlice(byteSlice, true)
	fmt.Printf("%v - asc\n", byteSlice)
	byteSlice = byteBackup
	mysort.SortSlice(byteSlice, false)
	fmt.Printf("%v - desc\n\n", byteSlice)

	runeBackup := runeSlice
	fmt.Printf("runeSlice:\n%v - orig\n", runeSlice)
	mysort.SortSlice(runeSlice, true)
	fmt.Printf("%v - asc\n", runeSlice)
	runeSlice = runeBackup
	mysort.SortSlice(runeSlice, false)
	fmt.Printf("%v - desc\n\n", runeSlice)

	float32Backup := float32Slice
	fmt.Printf("float32Slice:\n%v - orig\n", float32Slice)
	mysort.SortSlice(float32Slice, true)
	fmt.Printf("%v - asc\n", float32Slice)
	float32Slice = float32Backup
	mysort.SortSlice(float32Slice, false)
	fmt.Printf("%v - desc\n\n", float32Slice)

	float64Backup := float64Slice
	fmt.Printf("float64Slice:\n%v - orig\n", float64Slice)
	mysort.SortSlice(float64Slice, true)
	fmt.Printf("%v - asc\n", float64Slice)
	float64Slice = float64Backup
	mysort.SortSlice(float64Slice, false)
	fmt.Printf("%v - desc\n\n", float64Slice)

	complex64Backup := complex64Slice
	fmt.Printf("complex64Slice:\n%v - orig\n", complex64Slice)
	mysort.SortSlice(complex64Slice, true)
	fmt.Printf("%v - asc\n", complex64Slice)
	complex64Slice = complex64Backup
	mysort.SortSlice(complex64Slice, false)
	fmt.Printf("%v - desc\n\n", complex64Slice)

	complex128Backup := complex128Slice
	fmt.Printf("complex128Slice:\n%v - orig\n", complex128Slice)
	mysort.SortSlice(complex128Slice, true)
	fmt.Printf("%v - asc\n", complex128Slice)
	complex128Slice = complex128Backup
	mysort.SortSlice(complex128Slice, false)
	fmt.Printf("%v - desc\n\n", complex128Slice)

	stringBackup := stringSlice
	fmt.Printf("stringSlice:\n%v - orig\n", stringSlice)
	mysort.SortSlice(stringSlice, true)
	fmt.Printf("%v - asc\n", stringSlice)
	stringSlice = stringBackup
	mysort.SortSlice(stringSlice, false)
	fmt.Printf("%v - desc\n\n", stringSlice)

}
