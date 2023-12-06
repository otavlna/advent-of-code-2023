import * as fs from "fs"
import * as R from "ramda"

const data = fs.readFileSync("../../input.txt", "utf8")

const parsed = R.compose(
	R.map(R.compose(
		R.adjust(0, R.compose(parseInt, R.nth(1), R.split(" "))),
		R.adjust(1, R.compose(R.map(R.map(R.compose(R.adjust(0, parseInt), R.split(" ")))), R.map(R.map(R.trim)), R.map(R.split(",")), R.split(";")))
	)),
	R.dropLast(2),
	R.compose(R.map(R.split(":")), R.split("\n"))
)(data)

const sum = R.reduce((acc, val) =>
	acc + R.compose(
		R.product,
		R.reduce((acc, val) => [R.max(acc[0], val[0]), R.max(acc[1], val[1]), R.max(acc[2], val[2])], [1, 1, 1]),
		R.map(
			R.reduce((acc, val) => [R.max(acc[0], val[0]), R.max(acc[1], val[1]), R.max(acc[2], val[2])], [1, 1, 1])
		),
		R.map(
			R.compose(
				R.map(
					([amount, color]) => R.cond([
						[R.equals("red"), R.always([amount, 1, 1])],
						[R.equals("green"), R.always([1, amount, 1])],
						[R.equals("blue"), R.always([1, 1, amount])]
					])(color)
				)
			),
		),
		R.nth(1),
	)(val), 0)(parsed)

console.log(sum)

