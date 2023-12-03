const numbers = "0123456789".split("");

function isRelevantSymbol(char: string | undefined) {
  return char !== undefined && char !== "." && !numbers.includes(char);
}

function hasSymbolNearby(
  map: string[][],
  y: number,
  x: number,
  number: string[]
) {
  return (
    // left side
    isRelevantSymbol(map[y][x - 1]) ||
    isRelevantSymbol(map[y - 1]?.[x - 1]) ||
    isRelevantSymbol(map[y + 1]?.[x - 1]) ||
    // above or below
    number.some((_, xx) => {
      return isRelevantSymbol(map[y - 1]?.[x + xx]) || isRelevantSymbol(map[y + 1]?.[x + xx])
    }) ||
    // right side
    isRelevantSymbol(map[y][x + number.length]) ||
    isRelevantSymbol(map[y - 1]?.[x + number.length]) ||
    isRelevantSymbol(map[y + 1]?.[x + number.length])
  );
}

function computeSymbolsValue(map: string[][], y: number) {
  let sum = 0;
  let x = 0;
  let tmp: string[] = [];
  while (x <= map[y].length) {
    if (numbers.includes(map[y][x])) {
      tmp.push(map[y][x]);
    } else if (tmp.length) {
      if (hasSymbolNearby(map, y, x - tmp.length, tmp)) {
        sum += parseInt(tmp.join(""), 10);
      }
      tmp = [];
    }
    x++;
  }
  return sum;
}

function run(input: string) {
  const map = input.split("\n").map((line) => line.split(""));
  const sum = map.reduce((result, _, y) => {
    result += computeSymbolsValue(map, y);
    return result;
  }, 0);

  return sum;
}

console.info(
  run(
    `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
  )
);

console.info(run(require('fs').readFileSync(__dirname + '/input.txt', 'utf-8')));
