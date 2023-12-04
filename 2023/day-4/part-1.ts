function computeScore(line: string) {
  let score = 0;
  let [winningInput, myInput] = line.split("|");
  const winningNumbers = winningInput
    .substring(winningInput.indexOf(":") + 1)
    .split(' ').filter(n => n !== '');
  const myNumbers = myInput.trim().split(' ').filter(n => n !== '');;
  myNumbers.forEach((num) => {
    if (winningNumbers.includes(num)) {
      if (!score) {
        score = 1;
      } else {
        score *= 2;
      }
    }
  });
  return score;
}

function run(input: string) {
  const lines = input.split("\n");
  const sum = lines.reduce((result, line) => {
    result += computeScore(line);
    return result;
  }, 0);

  return sum;
}

console.info(
  run(
    `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
  )
);

console.info(run(require('fs').readFileSync(__dirname + '/input.txt', 'utf-8')));
