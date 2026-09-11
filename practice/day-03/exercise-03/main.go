package main

func main() {
}

func runeFrequency(text string) map[rune]int {
	mmap := make(map[rune]int)

	for _, ch := range text {
		mmap[ch]++
	}

	return mmap
}
