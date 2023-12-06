import * as fs from "fs"

const data = fs.readFileSync("../../input.txt", "utf8")
const grid = data.split("\n").map(row => row.split(""))

let sum = 0
for (let y = 0; y < grid.length; y++) {
	for (let x = 0; x < grid[y].length; x++) {
		if (!isNaN(parseInt(grid[y][x])) && isNeighboringSymbol(y, x)) {
			let num = ""
			let negativeOffset = -1
			while (true) {
				if (!isNaN(parseInt(grid[y][x + negativeOffset]))) {
					num = grid[y][x + negativeOffset] + num
					negativeOffset--
				}
				else {
					break;
				}
			}
			while (true) {
				if (!isNaN(parseInt(grid[y][x]))) {
					num += grid[y][x]
					x++
				}
				else {
					sum += parseInt(num)
					break;
				}
			}
		}
	}

}

console.log(sum)

function isNeighboringSymbol(y: number, x: number): boolean {
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
	const symbols = ["/", "#", "%", "+", "*", "@", "$", "=", "-", "&"]

	for (const offset of neighborOffsets) {
		try {
			if (symbols.includes(grid[y + offset[0]][x + offset[1]])) {
				return true
			}

		}
		catch (err) { }
	}

	return false
}


