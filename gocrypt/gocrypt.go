package gocrypt

import (
	"fmt"
)

func Encrypt_Pbox(plainText string) string {
	var cipherText string

	blocks := splitIntoBlocks(plainText, "plain")

	for _, block := range blocks {
		cipherText += fmt.Sprintf("%02x", Permutate_P_box(block, "in"))
	}
	return cipherText
}

func Decrypt_Pbox(cipherText string) string {
	var plainText string

	blocks := splitIntoBlocks(cipherText, "cipher")

	for _, block := range blocks {
		plainText += string(Permutate_P_box(block, "out"))
	}
	return plainText
}

// функція для розбиття відкритого або шифро- тексту на блоки по 8 біт
func splitIntoBlocks(text string, mode string) []byte {
	if mode == "plain" {
		blocks := make([]byte, len(text))

		for i := 0; i < len(text); i++ {
			blocks[i] = text[i]
		}

		return blocks
	} else if mode == "cipher" {
		blocks := make([]byte, len(text)/2)

		for i := 0; i < len(text)/2; i++ {
			blocks[i] = byte(HexToInt(text[i*2 : i*2+2]))
		}
		return blocks
	}
	return make([]byte, 0)
}

func Permutate_P_box(block byte, mode string) byte {
	// тетрада для старших чотирьох бітів
	nibble1 := (0xF0 & block) >> 4
	// тетрада для молодших чотирьох бітів
	nibble2 := 0x0F & block
	//fmt.Printf("%04b\n", nibble1)
	nibble1 = shuffleBits(nibble1, mode)
	//fmt.Printf("%04b\n", nibble1)
	//fmt.Printf("%04b\n", nibble2)
	nibble2 = shuffleBits(nibble2, mode)
	//fmt.Printf("%04b\n", nibble2)
	return (nibble1 << 4) | nibble2
}

func shuffleBits(num byte, mode string) byte {
	var bit1, bit2, bit3, bit4 byte
	if mode == "in" {
		// "змішаємо кожен біт тетради згідно з рисунком ../imgs/my_pbox_implementation.png"
		bit1 = (num & (1 << 0)) << 2 // переміщення 1-го біта на позицію 3
		bit2 = (num & (1 << 1)) >> 1 // переміщення 2-го біта на позицію 1
		bit3 = (num & (1 << 2)) << 1 // переміщення 3-го біта на позицію 4
		bit4 = (num & (1 << 3)) >> 2 // переміщення 4-го біта на позицію 2
	} else if mode == "out" {
		// "розмішаємо" позицію кожного біта назад
		bit1 = (num & (1 << 2)) >> 2 // переміщення 3-го біта (раніше 1-го) назад на позицію 1
		bit2 = (num & (1 << 0)) << 1 // переміщення 1-го біта (раніше 2-го) назад на позицію 2
		bit3 = (num & (1 << 3)) >> 1 // переміщення 4-го біта (раніше 3-го) назад на позицію 3
		bit4 = (num & (1 << 1)) << 2 // переміщення 2-го біта (раніше 4-го) назад на позицію 4
	}

	// Об'єднуємо біти разом
	return bit1 | bit2 | bit3 | bit4
}

/*
HELPERS
*/
func HexToInt(hex string) uint64 {
	if len(hex) > 16 {
		return 0
	}

	var result uint64

	for _, char := range hex {
		result = result << 4 // Переміщення на 4 біти вліво, резервуємо одне число 16 розряду -> 0001 0000
		switch {
		case '0' <= char && char <= '9':
			result += uint64(char - '0')
		case 'a' <= char && char <= 'f':
			result += uint64(char - 'a' + 10)
		case 'A' <= char && char <= 'F':
			result += uint64(char - 'A' + 10)
		default:
			return 0
		}
	}

	return result
}
