function forecastNextValue(numbers: number[]) {
  const stack: number[][] = [numbers];
  let current = stack[0];
  let next: number[] = [];
  let i = 0;

  while (i <= current.length - 1) {
    if (i === current.length - 1) {
      stack.push(next);
      current = next;
      next = [];
      i = 0;
      if (current.every((n) => n === 0)) {
        break;
      }
    } else {
      next.push(current[i + 1] - current[i]);
      i++;
    }
  }
  // console.info(stack);
  let entry = stack.pop();
  let prediction = 0;
  while (entry) {
    prediction = entry[0] - prediction;
    entry = stack.pop();
  }
  return prediction;
}

function run(input: string) {
  const map = input.split("\n");
  const sum = map.reduce((result, numbers) => {
    result += forecastNextValue(
      numbers
        .trim()
        .split(" ")
        .map((n) => parseInt(n, 10))
    );
    return result;
  }, 0);

  return sum;
}

console.info(
  run(
    `10 13 16 21 30 45`
  )
);

console.info(run(require('fs').readFileSync(__dirname + '/input.txt', 'utf-8')));
