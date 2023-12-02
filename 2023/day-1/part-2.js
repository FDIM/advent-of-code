function computeCalibrationValue(line) {
  const map = {
    'one': 1,
    'two': 2,
    'three': 3,
    'four': 4,
    'five': 5,
    'six': 6,
    'seven': 7,
    'eight': 8,
    'nine': 9
  };
  const words = Object.keys(map);
  let firstDigit = undefined;
  let lastDigit = undefined;
  for (let i = 0; i< line.length; i++) {
    for (let j = 1; j <= 9; j++) {
      const char = line.charAt(i);
      const isChar = `${j}` === char;
      const word = !isChar && words[j - 1];
      const isWord = !isChar && line.indexOf(word, i) === i;

      if (isChar || isWord) {
        let digit = j;
        if(isWord) {
          digit = map[word];
          i+= word.length - 1;
        }
        if (!firstDigit) {
          firstDigit = digit;
        }
        lastDigit = digit;
        break;
      }
    }
  }
  return parseInt(`${firstDigit}${lastDigit}`, 10);
}

function run(input) {
  const lines = input.split('\n');
  const sum = lines.reduce((result, line) => {
    result+= computeCalibrationValue(line);
    return result;
  }, 0);

  return sum;
}

console.info(run(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`));


console.info(run(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`));

console.info(run(require('fs').readFileSync(__dirname + '/input.txt', 'utf-8')));