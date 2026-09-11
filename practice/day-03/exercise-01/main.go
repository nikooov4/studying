package main

func main() {
}

func wordFrequency(words []string) map[string]int {
	mmap := make(map[string]int)

	if len(words) == 0 {
		return mmap
	}

	for _, v := range words {
		mmap[v]++
	}

	return mmap
}
