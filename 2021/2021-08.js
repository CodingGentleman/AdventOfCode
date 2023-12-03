function part1(data) {
	return data.split(/\r?\n/).map(v => v.split(' | ')[1].split(' ').map(v => v.length).filter(v => [2,3,4,7].includes(v))).flat().length
}

function part2(data) {
	return data.split(/\r?\n/).map(v => part2Line(v)).reduce((a,b) => a+b)
}

function part2Line(line) {
	let solution = []
  const getVal = (num) => solution.find(x => x.num == num).val
  const getNum = (val) => solution.find(x => [...val].sort().toString() == [...x.val].sort().toString()).num
  const valContainsNum = (num, v) => [...getVal(num)].every(e => [...v].includes(e))
  const numContainsVal = (num, v) => [...v].every(e => [...getVal(num)].includes(e))

	let data = line.split(' | ')
  let numbers = data[0].split(' ')
  
  // one = 2L
  solution.push({num: 1, val: numbers.filter(v => v.length == 2)[0]})
  // four = 4L
  solution.push({num: 4, val: numbers.filter(v => v.length == 4)[0]})
  // seven = 3L
  solution.push({num: 7, val: numbers.filter(v => v.length == 3)[0]})
  // eight = 7L
  solution.push({num: 8, val: numbers.filter(v => v.length == 7)[0]})
  // nine = 6L contains 'four'
  solution.push({num: 9, val: numbers.filter(v => v.length == 6 && valContainsNum(4,v))[0]})
  // three = 5L contains 'one'
  solution.push({num: 3, val: numbers.filter(v => v.length == 5 && valContainsNum(1,v))[0]})
  // six = 6L not contains 'one'
  solution.push({num: 6, val: numbers.filter(v => v.length == 6 && !valContainsNum(1,v))[0]})
  // zero = 6L not six or nine 
  solution.push({num: 0, val: numbers.filter(v => v.length == 6 && v != getVal(9) && v != getVal(6))[0]})
  // five = 'six' contains 5L
  solution.push({num: 5, val: numbers.filter(v => v.length == 5 && numContainsVal(6,v))[0]})
  // two = 5L not three or five 
  solution.push({num: 2, val: numbers.filter(v => v.length == 5 && v != getVal(3) && v != getVal(5))[0]})

  return parseInt(data[1].split(' ').map(v => getNum(v)).join(''))
}
