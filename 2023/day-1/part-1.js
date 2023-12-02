function computeCalibrationValue(line) {
  const digits = '0123456789'.split('');
  let firstDigit = undefined;
  let lastDigit = undefined;
  for (let i = 0; i< line.length; i++) {
    const char = line.charAt(i);
    if (digits.includes(char)) {
      const digit = parseInt(char, 10);
      if (!firstDigit) {
        firstDigit = digit;
      }
      lastDigit = digit;
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

console.info(run(require('fs').readFileSync(__dirname + '/input.txt', 'utf-8')));