import * as fs from "fs"

const data = fs.readFileSync("../../input.txt", "utf8")
const lines = data.split("\n")
let sum = 0

for (const line of lines) {
	const lineChars = line.split("")
	const firstCalibartionDigit = lineChars.find((char) => char.charCodeAt(0) > 47 && char.charCodeAt(0) < 58)
	const secondCalibartionDigit = lineChars.findLast((char) => char.charCodeAt(0) > 47 && char.charCodeAt(0) < 58)
	const calibrationValue = parseInt(`${firstCalibartionDigit}${secondCalibartionDigit}`)
	if (!isNaN(calibrationValue)) sum += calibrationValue
}

console.log(sum)

