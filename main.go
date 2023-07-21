package main

import (
	"fmt"
	"strings"
)

var (
	animals = []string{"dog", "cat", "bat", "cock", "cow", "pig", "fox", "ant", "bird", "lion", "wolf", "deer",
		"bear", "frog", "hen", "mole", "duck", "goat"}
	maxNameLength      = 0
	minNameLength      = len(animals[0])
	animalsWithShuffle = make(map[string][]string, 0)
)

func init() {
	for _, animal := range animals {
		if len(animal) > maxNameLength {
			maxNameLength = len(animal)
		}
		if len(animal) < minNameLength {
			minNameLength = len(animal)
		}
	}

	for _, animal := range animals {
		animalsWithShuffle[animal] = permutation([]rune(animal), 0, len(animal)-1)
	}

}

func main() {

	max_animals("goatcode")
	max_animals("cockdogwdufrbir")
	max_animals("dogdogdog")
	max_animals("cockdogwdufrbirtaccatdoodddognnnehhenattaentantlioxwlforede")
	max_animals("caasdsdocasdkdogwdufrbdiastaccdtdoodfoooofooforofrofgorfjkorkforfjkofrasddocatoyoidfhmemoledouckgoatasdannehhenattaesdtantlioxasdasdasforede")
}

func permutation(sampleRune []rune, left, right int) []string {
	str := make([]string, 0)
	if left == right {
		str = append(str, string(sampleRune))

	} else {
		for i := left; i <= right; i++ {
			sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
			str = append(str, permutation(sampleRune, left+1, right)...)
			sampleRune[left], sampleRune[i] = sampleRune[i], sampleRune[left]
		}
	}

	return str
}

func ChunksString(s string, chunkSize int) []string {
	if chunkSize <= 0 {
		return nil
	}

	var chunks []string
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}

	return chunks
}

func max_animals(txt string) {
	resultAnimal := make([]string, 0)

mainLoop:
	for {
		for i := minNameLength; i <= maxNameLength; i++ {
			nameChunks := ChunksString(txt, i)
			for _, nameChunk := range nameChunks {

				nameChunkPermuts := permutation([]rune(nameChunk), 0, len(nameChunk)-1)
				for _, nameChunkPermut := range nameChunkPermuts {
					for anim, animShuff := range animalsWithShuffle {
						for _, v := range animShuff {
							if nameChunkPermut == v {
								var found bool
								for _, ranim := range resultAnimal {
									if ranim == anim {
										found = true
									}
								}
								if !found {
									resultAnimal = append(resultAnimal, anim)
								}
								txt = strings.Replace(txt, nameChunkPermut, "", 1)
								continue mainLoop
							}
						}
					}
				}

			}

		}
		if len(txt) >= minNameLength {
			txt = txt[minNameLength:]
		} else {
			break mainLoop
		}

	}
	fmt.Println("Number of animals is", len(resultAnimal), resultAnimal)

}
