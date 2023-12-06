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
	R.compose(
		R.ifElse(R.all(R.equals(true)), R.always(R.add(acc, R.nth(0, val))), R.always(acc)),
		R.chain(
			R.map(
				([amount, color]) => R.cond([
					[R.equals("red"), R.always(R.lte(amount, 12))],
					[R.equals("green"), R.always(R.lte(amount, 13))],
					[R.equals("blue"), R.always(R.lte(amount, 14))]
				])(color)
			)
		),
		R.nth(1),
	)(val), 0)(parsed)

console.log(sum)

