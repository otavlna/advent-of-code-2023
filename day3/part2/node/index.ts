import * as fs from "fs"

const data = fs.readFileSync("../../input.txt", "utf8")

const grid = data.split("\n").map(row => row.split(""))

let sum = 0
for (let y = 0; y < grid.length; y++) {
	for (let x = 0; x < grid[y].length; x++) {
		if (grid[y][x] === "*") {
			const neighbors = neighboringNumbers(y, x)
			if (neighbors.length === 2) {
				sum += neighbors[0] * neighbors[1]
			}

		}
	}

}

console.log(sum)

function neighboringNumbers(y: number, x: number): number[] {
	const neighborOffsets = [
		[-1, 0],
		[-1, 1],
		[0, 1],
		[1, 1],
		[1, 0],
		[1, -1],
		[0, -1],
		[-1, -1],
	]
	const neighbors = []
	const numbers = [
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	]

	for (const offset of neighborOffsets) {
		try {
			const char = grid[y + offset[0]][x + offset[1]]
			if (numbers.includes(char)) {
				let num = ""
				let negativeOffset = -1
				let positiveOffset = 0
				while (true) {
					const numChar = grid[y + offset[0]][x + offset[1] + negativeOffset]
					if (!isNaN(parseInt(numChar))) {
						num = numChar + num
						negativeOffset--
					}
					else {
						break;
					}
				}
				while (true) {
					const numChar = grid[y + offset[0]][x + offset[1] + positiveOffset]
					if (!isNaN(parseInt(numChar))) {
						num += numChar
						positiveOffset++
					}
					else {
						neighbors.push(parseInt(num))
						break;
					}
				}
			}

		}
		catch (err) { }
	}

	return [...new Set(neighbors)]
}


