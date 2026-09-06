package utils

type reducetype func(interface{}) interface{}

type filtertype func(interface{}) bool

func InSlice(v string, sl []string) bool { _ = "STUB: not implemented"; return false }

func InSliceIface(v interface{}, sl []interface{}) bool { _ = "STUB: not implemented"; return false }

func SliceRandList(min, max int) []int { _ = "STUB: not implemented"; return nil }

func SliceMerge(slice1, slice2 []interface{}) (c []interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func SliceReduce(slice []interface{}, a reducetype) (dslice []interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func SliceRand(a []interface{}) (b interface{}) { _ = "STUB: not implemented"; return nil }

func SliceSum(intslice []int64) (sum int64) { _ = "STUB: not implemented"; return 0 }

func SliceFilter(slice []interface{}, a filtertype) (ftslice []interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func SliceDiff(slice1, slice2 []interface{}) (diffslice []interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func SliceIntersect(slice1, slice2 []interface{}) (diffslice []interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func SliceChunk(slice []interface{}, size int) (chunkslice [][]interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func SliceRange(start, end, step int64) (intslice []int64) { _ = "STUB: not implemented"; return nil }

func SlicePad(slice []interface{}, size int, val interface{}) []interface{} {
	_ = "STUB: not implemented"
	return nil
}

func SliceUnique(slice []interface{}) (uniqueslice []interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func SliceShuffle(slice []interface{}) []interface{} { _ = "STUB: not implemented"; return nil }
