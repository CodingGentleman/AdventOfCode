var input = [3,4,3,1,2]

// map array to objects array with value (e.g. numbers) and a count
var a = Array.from(new Set(input)).map(v => ({value:v, count: input.filter(f => f === v).length}));
// all possible keys
var uq = [0,1,2,3,4,5,6,7,8]

for (var i = 0; i < 256; i++) {
  // count 0s aka how many 8s to push
  var newA = a.filter(v => v.value == 0).map(v => v.count).reduce((a,b)=>a+b,0)
  // before pushing, reduce value by 1 or set to 6 if 0
  a = a.map(v => v.value == 0 ? {value:6, count:v.count} : {value:v.value - 1, count:v.count})
  // push 8s
  a.push({value:8, count:newA})
  // reduce array by summing the count of the unique keys
	a = uq.map(v => ({value:v, count: a.filter(w => w.value === v).map(w => w.count).reduce((a,b)=>a+b,0)}))
}
var result = a.map(v => v.count).reduce((a,b)=>a+b,0)
console.log(26984457539 == result)
