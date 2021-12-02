package aocutils

import "io/ioutil"

const inputFile string = "input.txt"

func InputToSliceByte() []byte {
	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	return file
}

func InputToSliceInt() []int {
	return ToSliceInt(InputToSliceByte())
}
