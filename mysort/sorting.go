package mysort

import (
	"errors"
	"reflect"
)

//Sort sorts slices of basic go types using quicksort in Hoare's implementation
func SortSlice(a interface{}, asc bool) error {
	switch sl := a.(type) {
	case []bool:
		sortBool(sl, 0, len(sl)-1, asc)

	case []int, []int8, []int16, []int32:
		val := reflect.ValueOf(sl)
		slice := make([]int64, val.Len())
		for i := 0; i < val.Len(); i++ {
			slice[i] = val.Index(i).Int()
		}

		sortInts(slice, 0, val.Len()-1, asc)
		changeType(sl, slice)

	case []uint, []uint8, []uint16, []uint32, []uint64, []uintptr:
		val := reflect.ValueOf(sl)
		slice := make([]int64, val.Len())
		for i := 0; i < val.Len(); i++ {
			slice[i] = int64(val.Index(i).Uint())
		}

		sortInts(slice, 0, val.Len()-1, asc)
		changeType(sl, slice)

	case []int64:
		sortInts(sl, 0, len(sl)-1, asc)

	case []float32:
		slice := make([]float64, len(sl))
		for n, c := range sl {
			slice[n] = float64(c)
		}

		sortFloats(slice, 0, len(slice)-1, asc)
		for n, c := range slice {
			sl[n] = float32(c)
		}

	case []float64:
		sortFloats(sl, 0, len(sl)-1, asc)

	case []complex64:
		slice := make([]complex128, len(sl))
		for n, c := range sl {
			slice[n] = complex128(c)
		}

		sortComplexReal(slice, 0, len(slice)-1, asc)
		var smallHigh, low int
		for low < len(slice)-1 {
			smallHigh = getRealNumberLimits(slice, low)
			sortComplexImag(slice, low, smallHigh, asc)
			low = smallHigh + 1
		}
		for n, c := range slice {
			sl[n] = complex64(c)
		}

	case []complex128:
		sortComplexReal(sl, 0, len(sl)-1, asc)
		var smallHigh, low int
		for low < len(sl)-1 {
			smallHigh = getRealNumberLimits(sl, low)
			sortComplexImag(sl, low, smallHigh, asc)
			low = smallHigh + 1
		}

	case []string:
		sortString(sl, 0, len(sl)-1, asc)

	default:
		return errors.New("mysort: unsupported type")
	}
	return nil
}

func partitionBool(a []bool, low, high int, asc bool) int {
	pivot := a[int((high+low)/2)]
	low -= 1
	high += 1
	for {
		if asc {
			low += 1
			for !a[low] && pivot {
				low += 1
			}
			high -= 1
			for a[high] && !pivot {
				high -= 1
			}
		} else {
			low += 1
			for a[low] && !pivot {
				low += 1
			}
			high -= 1
			for !a[high] && pivot {
				high -= 1
			}
		}

		if low >= high {
			return high
		}
		a[low], a[high] = a[high], a[low]
	}
}

func sortBool(a []bool, low, high int, asc bool) {
	if low < high {
		p := partitionBool(a, low, high, asc)
		sortBool(a, low, p, asc)
		sortBool(a, p+1, high, asc)
	}
}

func changeType(tp interface{}, a []int64) interface{} {
	switch sl := tp.(type) {
	case []int:
		for n, c := range a {
			sl[n] = int(c)
		}
		return sl

	case []int8:
		for n, c := range a {
			sl[n] = int8(c)
		}
		return sl

	case []int16:
		for n, c := range a {
			sl[n] = int16(c)
		}
		return sl

	case []int32:
		for n, c := range a {
			sl[n] = int32(c)
		}
		return sl

	case []int64:
		return sl

	case []uint:
		for n, c := range a {
			sl[n] = uint(c)
		}
		return sl

	case []uint8:
		for n, c := range a {
			sl[n] = uint8(c)
		}
		return sl

	case []uint16:
		for n, c := range a {
			sl[n] = uint16(c)
		}
		return sl

	case []uint32:
		for n, c := range a {
			sl[n] = uint32(c)
		}
		return sl

	case []uint64:
		for n, c := range a {
			sl[n] = uint64(c)
		}
		return sl

	case []uintptr:
		for n, c := range a {
			sl[n] = uintptr(c)
		}
		return sl

	}
	return nil
}

func partitionInts(a []int64, low, high int, asc bool) int {
	pivot := a[int((high+low)/2)]
	low -= 1
	high += 1

	for {
		if asc {
			low += 1
			for a[low] < pivot {
				low += 1
			}
			high -= 1
			for a[high] > pivot {
				high -= 1
			}
		} else {
			low += 1
			for a[low] > pivot {
				low += 1
			}
			high -= 1
			for a[high] < pivot {
				high -= 1
			}
		}

		if low >= high {
			return high
		}
		a[low], a[high] = a[high], a[low]
	}
}

func sortInts(a []int64, low, high int, asc bool) {
	if low < high {
		p := partitionInts(a, low, high, asc)
		sortInts(a, low, p, asc)
		sortInts(a, p+1, high, asc)
	}
}

func partitionFloats(a []float64, low, high int, asc bool) int {
	pivot := a[int((high+low)/2)]
	low -= 1
	high += 1

	for {
		if asc {
			low += 1
			for a[low] < pivot {
				low += 1
			}
			high -= 1
			for a[high] > pivot {
				high -= 1
			}
		} else {
			low += 1
			for a[low] > pivot {
				low += 1
			}
			high -= 1
			for a[high] < pivot {
				high -= 1
			}
		}

		if low >= high {
			return high
		}
		a[low], a[high] = a[high], a[low]
	}
}

func sortFloats(a []float64, low, high int, asc bool) {
	if low < high {
		p := partitionFloats(a, low, high, asc)
		sortFloats(a, low, p, asc)
		sortFloats(a, p+1, high, asc)
	}
}

func partitionComplexReal(a []complex128, low, high int, asc bool) int {
	pivot := a[int((high+low)/2)]
	low -= 1
	high += 1

	for {
		if asc {
			low += 1
			for real(a[low]) < real(pivot) {
				low += 1
			}
			high -= 1
			for real(a[high]) > real(pivot) {
				high -= 1
			}
		} else {
			low += 1
			for real(a[low]) > real(pivot) {
				low += 1
			}
			high -= 1
			for real(a[high]) < real(pivot) {
				high -= 1
			}
		}

		if low >= high {
			return high
		}
		a[low], a[high] = a[high], a[low]
	}
}

func sortComplexReal(a []complex128, low, high int, asc bool) {
	if low < high {
		p := partitionComplexReal(a, low, high, asc)
		sortComplexReal(a, low, p, asc)
		sortComplexReal(a, p+1, high, asc)
	}
}

func getRealNumberLimits(a []complex128, low int) (high int) {
	lowReal := real(a[low])
	highReal := real(a[low+1])
	for high = low + 2; lowReal == highReal; high++ {
		if high == len(a) {
			return high - 1
		}
		highReal = real(a[high])
	}
	return high - 2
}

func partitionComplexImag(a []complex128, low, high int, asc bool) int {
	pivot := a[int((high+low)/2)]
	low -= 1
	high += 1

	for {
		if asc {
			low += 1
			for imag(a[low]) < imag(pivot) {
				low += 1
			}
			high -= 1
			for imag(a[high]) > imag(pivot) {
				high -= 1
			}
		} else {
			low += 1
			for imag(a[low]) > imag(pivot) {
				low += 1
			}
			high -= 1
			for imag(a[high]) < imag(pivot) {
				high -= 1
			}
		}

		if low >= high {
			return high
		}
		a[low], a[high] = a[high], a[low]
	}
}

func sortComplexImag(a []complex128, low, high int, asc bool) {
	if low < high {
		p := partitionComplexImag(a, low, high, asc)
		sortComplexImag(a, low, p, asc)
		sortComplexImag(a, p+1, high, asc)
	}
}

func partitionString(a []string, low, high int, asc bool) int {
	pivot := a[int((high+low)/2)]
	low -= 1
	high += 1

	for {
		if asc {
			low += 1
			for a[low] < pivot {
				low += 1
			}
			high -= 1
			for a[high] > pivot {
				high -= 1
			}
		} else {
			low += 1
			for a[low] > pivot {
				low += 1
			}
			high -= 1
			for a[high] < pivot {
				high -= 1
			}
		}

		if low >= high {
			return high
		}
		a[low], a[high] = a[high], a[low]
	}
}

func sortString(a []string, low, high int, asc bool) {
	if low < high {
		p := partitionString(a, low, high, asc)
		sortString(a, low, p, asc)
		sortString(a, p+1, high, asc)
	}
}

//SortStructs sorts user-made structs, given that "a" is a pointer to slice of structs
//and key's type (which is name of a field by which struct would be sorted) is one of the basic GO types
//which will be sorted in ascending order when asc is true and the other way around, fields must be exported
func SortStructs(a interface{}, key string, asc bool) error {
	var (
		fName    string
		fPos     int
		valSlice []reflect.Value
	)

	if a == nil {
		return errors.New("mysort: given interface is empty")
	}

	structSlicePointer := reflect.ValueOf(a)
	if structSlicePointer.Kind() != reflect.Ptr {
		return errors.New("mysort: given interface is not pointer")
	} else if structSlicePointer.Elem().Kind() != reflect.Slice {
		return errors.New("mysort: given interface is not pointer to slice")
	}

	//append single structs here
	for i := 0; i < structSlicePointer.Elem().Len(); i++ {
		valSlice = append(valSlice, structSlicePointer.Elem().Index(i))
	}

	if valSlice[0].Kind() != reflect.Struct {
		return errors.New("mysort: interface is not a struct")
	}

	//search for key here
	sl := valSlice[0]
	for ; fPos < sl.NumField(); fPos++ { //range over fields and search for match with key
		fName = sl.Type().Field(fPos).Name
		if fName == key {
			break
		}
	}

	if fPos == sl.NumField() && fName != key {
		return errors.New("mysort: key not found")
	} else if !basicGoType(sl.FieldByName(key)) {
		return errors.New("mysort: key is not a basic go type")
	}

	sortReflect(valSlice, 0, len(valSlice)-1, key, asc)
	return nil
}

func basicGoType(a reflect.Value) bool {
	return a.Kind() == reflect.Bool ||
		a.Kind() == reflect.Int || a.Kind() == reflect.Int8 || a.Kind() == reflect.Int16 || a.Kind() == reflect.Int32 || a.Kind() == reflect.Int64 ||
		a.Kind() == reflect.Uint || a.Kind() == reflect.Uint8 || a.Kind() == reflect.Uint16 || a.Kind() == reflect.Uint32 || a.Kind() == reflect.Uint64 ||
		a.Kind() == reflect.Complex64 || a.Kind() == reflect.Complex128 || a.Kind() == reflect.Float32 || a.Kind() == reflect.Float64 || a.Kind() == reflect.String
}

func aLessThanBReflect(a reflect.Value, b reflect.Value) bool {
	switch a.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		int1 := a.Int()
		int2 := b.Int()
		return int1 < int2
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uint1 := a.Uint()
		uint2 := b.Uint()
		return uint1 < uint2
	case reflect.Bool:
		bool1 := a.Bool()
		bool2 := b.Bool()
		return !bool1 && bool2
	case reflect.Float32, reflect.Float64:
		float1 := a.Float()
		float2 := b.Float()
		return float1 < float2
	case reflect.Complex64, reflect.Complex128:
		complex1 := a.Complex()
		complex2 := b.Complex()
		if real(complex1) == real(complex2) {
			return imag(complex1) < imag(complex2)
		}
		return real(complex1) < real(complex2)
	case reflect.String:
		str1 := a.String()
		str2 := b.String()
		return str1 < str2
	}
	return false
}

//This is Hoare partition scheme adapted for this code
func partitionReflect(a []reflect.Value, low, high int, key string, asc bool) int {
	pivot := a[int((high+low)/2)]
	low -= 1
	high += 1

	for {
		if asc {
			low += 1
			for a[low].FieldByName(key) != pivot.FieldByName(key) && aLessThanBReflect(a[low].FieldByName(key), pivot.FieldByName(key)) {
				low += 1
			}
			high -= 1
			for a[high].FieldByName(key) != pivot.FieldByName(key) && !aLessThanBReflect(a[high].FieldByName(key), pivot.FieldByName(key)) {
				high -= 1
			}
		} else {
			low += 1
			for a[low].FieldByName(key) != pivot.FieldByName(key) && !aLessThanBReflect(a[low].FieldByName(key), pivot.FieldByName(key)) {
				low += 1
			}
			high -= 1
			for a[high].FieldByName(key) != pivot.FieldByName(key) && aLessThanBReflect(a[high].FieldByName(key), pivot.FieldByName(key)) {
				high -= 1
			}
		}

		if low >= high {
			return high
		}

		//allocate memory for struct and copy a[low]'s value here
		//couldn't do temp := a[low].Interface() because it shared 1 memory adress
		temp := reflect.New(a[low].Type()).Interface()
		temp = a[low].Interface()
		a[low].Set(a[high])
		a[high].Set(reflect.ValueOf(temp))
	}
}

func sortReflect(a []reflect.Value, low, high int, key string, asc bool) {
	if low < high {
		p := partitionReflect(a, low, high, key, asc)
		sortReflect(a, low, p, key, asc)
		sortReflect(a, p+1, high, key, asc)
	}
}
