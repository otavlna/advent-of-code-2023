import * as fs from "fs"

const data = fs.readFileSync("../../input.txt", "utf8")
const lines = data.split("\n")
let sum = 0

const wordDigitMap = {
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
};
const numberDigitMap = {
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

type IndexWithDigit = {
	digit: number;
	index: number;
}

function iterateOverMapAtLine(map: Record<string, number>, line: string, firstDigit: IndexWithDigit, secondDigit: IndexWithDigit): void {
	for (const key of Object.keys(map)) {
		const index = line.indexOf(key)
		const digit = map[key as keyof typeof map]
		if (index !== -1 && index < firstDigit.index) {
			firstDigit.index = index
			firstDigit.digit = digit
		}
		const lastIndex = line.lastIndexOf(key)
		if (lastIndex !== -1 && lastIndex > secondDigit.index) {
			secondDigit.index = lastIndex
			secondDigit.digit = digit
		}
	}
}


for (const line of lines) {
	const firstDigit: IndexWithDigit = { digit: 0, index: Infinity }
	const secondDigit: IndexWithDigit = { digit: 0, index: -1 }

	iterateOverMapAtLine(wordDigitMap, line, firstDigit, secondDigit)
	iterateOverMapAtLine(numberDigitMap, line, firstDigit, secondDigit)

	const calibrationValue = parseInt(`${firstDigit.digit}${secondDigit.digit}`)
	if (!isNaN(calibrationValue)) sum += calibrationValue
}

console.log(sum)

