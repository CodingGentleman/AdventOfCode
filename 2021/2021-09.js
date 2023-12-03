function part1(data) {
	let lines = data.split(/\r?\n/).map(v => '9'+v+'9')
  lines.unshift('9'.repeat(lines[0].length))
  lines.push('9'.repeat(lines[0].length))
  lines = lines.map(v => v.split(''))
  let lowPoints = []
  for(let i = 1; i < lines.length-1; i++) {
    for(let j = 1; j < lines[i].length-1; j++) {
      let cur = lines[i][j]
      if(cur < lines[i-1][j] && cur < lines[i][j+1] && cur < lines[i+1][j] && cur < lines[i][j-1]) {
        lowPoints.push(parseInt(cur))
      }
    }
  }
  return lowPoints.map(v => v+1).reduce((a,b) => a+b)
}

let input = 
'2199943210\n'+
'3987894921\n'+
'9856789892\n'+
'8767896789\n'+
'9899965678'
console.log(15 == part1(input))
