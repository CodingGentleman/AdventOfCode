var input = [16,1,2,0,4,2,7,1,2,14]

// get the median as a starting point
input = input.sort((a, b) => a - b)
var median = input[input.length/2]

// map array to objects array with pos (position) and a count
var a = Array.from(new Set(input)).map(v => ({pos:v, count: input.filter(f => f === v).length}));

function part1Distance(index) {
	return a.map(v => Math.abs(index-v.pos)*v.count).reduce((a,b) => a+b)
}

function gaussSum(n) {
	return n * (n+1) / 2
}
function part2Distance(index) {
	return a.map(v => gaussSum(Math.abs(index-v.pos))*v.count).reduce((a,b) => a+b)
}

var resultPart1 = Number.MAX_SAFE_INTEGER
var resultPart2 = Number.MAX_SAFE_INTEGER

for(let i = median; i < median+10000; i++) {
	// calc the fuel needed
  resultPart1 = Math.min(resultPart1, part1Distance(i), part1Distance(median-i))
  resultPart2 = Math.min(resultPart2, part2Distance(i), part2Distance(median-i))
}

console.log('Part1: ' + (37 == resultPart1))
console.log('Part2: ' + (168 == resultPart2))
