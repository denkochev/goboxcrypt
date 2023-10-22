package gocrypt

import "fmt"

func Encrypt_Pbox(source string) {
	fmt.Println("I'm billionare")
}

// функція для розбиття вхідного тексту на блоки по 8 біт
func splitIntoBlocks(plainText string) []byte {

	blocks := make([]byte, len(plainText))

	for i := 0; i < len(plainText); i++ {
		blocks[i] = plainText[i]
	}

	return blocks
}

func Permutate_P_box(block byte) byte {
	// тетрада для старших чотирьох бітів
	nibble1 := (0xF0 & block) >> 4
	// тетрада для молодших чотирьох бітів
	nibble2 := 0x0F & block
	//fmt.Printf("%04b\n", nibble1)
	nibble1 = shuffleBits(nibble1)
	//fmt.Printf("%04b\n", nibble1)
	//fmt.Printf("%04b\n", nibble2)
	nibble2 = shuffleBits(nibble2)
	//fmt.Printf("%04b\n", nibble2)
	return (nibble1 << 4) | nibble2
}

func shuffleBits(num byte) byte {
	// Виберемо кожний біт із джерела та змінимо його позицію
	bit1 := (num & (1 << 0)) << 2 // переміщення 1-го біта на позицію 3
	bit2 := (num & (1 << 1)) >> 1 // переміщення 2-го біта на позицію 1
	bit3 := (num & (1 << 2)) << 1 // переміщення 3-го біта на позицію 4
	bit4 := (num & (1 << 3)) >> 2 // переміщення 4-го біта на позицію 2

	// Об'єднуємо біти разом
	return bit1 | bit2 | bit3 | bit4
}
