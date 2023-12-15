package slices

func RotateRightTimes(text string, times int) string {
	for i := 1; i <= times; i++ {
		text = RotateRight(text)
	}
	return text
}

func RotateRight(text string) string {
	size := len(text)
	rotatedText := ""
	for fromIndex := range text {
		toIndex := (size - 1 + fromIndex) % size
		char := string(text[toIndex])
		rotatedText = rotatedText + char
	}
	return rotatedText
}

func RotateRightVersion2(text string) string {
	size := len(text)
	var rotatedChars []rune = make([]rune, size)
	for posicionNueva := range text {
		posicionOriginal := (size - 1 + posicionNueva) % size
		char := rune(text[posicionOriginal])
		rotatedChars[posicionNueva] = char
	}
	return string(rotatedChars) // hay una compatibilidad entre []rune y string, tmb entre []byte y string, y de esta forma podemos transformar de un tipo a otro con "facilidad"
}

func RotateRightVersion3(text string) string {
	size := len(text)
	if size > 1 {
		primerCaracterDelStringRotado := string(text[size-1])
		elRestoDeLosCaracteresDelStringRotato := text[0 : size-1]
		return primerCaracterDelStringRotado + elRestoDeLosCaracteresDelStringRotato
	}
	return text
}

func RotateRightVersion4(textOriginal string) string { // textoOrigial = abc
	size := len(textOriginal)
	if size > 1 {
		primerCaracterDelStringRotado := string(textOriginal[size-1])
		textRotado := primerCaracterDelStringRotado // c
		for i := 0; i < size-1; i++ {
			textRotado = textRotado + string(textOriginal[i]) // 1: ca, 2: cab
		}
		return textRotado
	}
	return textOriginal
}

func RotateLeft(text string) string {
	size := len(text)
	rotatedText := ""
	for i := range text {
		char := string(text[(size+1+i)%size])
		rotatedText = rotatedText + char
	}
	return rotatedText
}
