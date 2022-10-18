package knothash

import "github.com/richardc/advent-go/slices"

func Hash(input []byte) [16]byte {
	data := [256]byte{}
	for i := range data {
		data[i] = byte(i)
	}
	input = append(input, []byte{17, 31, 73, 47, 23}...)

	cur := 0
	skip := 0
	for round := 0; round < 64; round++ {
		for _, length := range input {
			length := int(length)
			for i, j := cur, cur+length-1; i < j; i, j = i+1, j-1 {
				data[i%len(data)], data[j%len(data)] = data[j%len(data)], data[i%len(data)]
			}

			cur = (cur + length + skip) % len(data)
			skip++
		}
	}

	dense := [16]byte{}
	for d := range dense {
		dense[d] = slices.Fold(data[d*16:(d+1)*16], 0, func(a, b byte) byte { return a ^ b })
	}

	return dense
}
