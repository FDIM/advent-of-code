function computeDistance(hold: number, duration: number) {
  const speed = hold;
  return (duration - hold) * speed;
}

function computeDifferentWaysToWin(duration: number, maxDistance: number) {
  let chancesToWin = 0;
  for (let i = 0; i < duration; i++) {
    const distance = computeDistance(i, duration);
    if (distance > maxDistance) {
      chancesToWin++;
    }
  }
  return chancesToWin;
}

function run(input: string) {
  const [times, distances] = input.split("\n").map((line) => {
    return line
      .substr(line.indexOf(":") + 1)
      .trim()
      .split(/[\s]+/)
      .map((n) => parseInt(n, 10));
  });
  console.info(times);
  console.info(distances);
  const total = times.reduce((res, time, index) => {
    const waysToWin = computeDifferentWaysToWin(time, distances[index]);
    if (waysToWin > 0) {
      res *= waysToWin;
    }
    return res;
  }, 1);
  return total;
}

console.info(
  run(
    `Time:      7  15   30
    Distance:  9  40  200`
  )
);

console.info(run(require('fs').readFileSync(__dirname + '/input.txt', 'utf-8')));
