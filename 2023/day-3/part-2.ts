const numbers = "0123456789".split("");

function getNearbyGearSymbolPosition(
  map: string[][],
  y: number,
  x: number,
  number: string[]
): string | undefined {
  let pos: string | undefined;

  function checkLocation(yy: number, xx: number) {
    if (map[yy]?.[xx] === "*") {
      pos = `${yy}x${xx}`;
    }
  }
  // left side
  checkLocation(y, x - 1);
  checkLocation(y - 1, x - 1);
  checkLocation(y + 1, x - 1);
  // above or below
  number.forEach((_, xx) => {
    checkLocation(y - 1, x + xx);
    checkLocation(y + 1, x + xx);
  });
  // right side
  checkLocation(y, x + number.length);
  checkLocation(y - 1, x + number.length);
  checkLocation(y + 1, x + number.length);
  return pos;
}

function findAndCollectGears(
  gears: Record<string, number[]>,
  map: string[][],
  y: number
) {
  let x = 0;
  let tmp: string[] = [];
  while (x <= map[y].length) {
    if (numbers.includes(map[y][x])) {
      tmp.push(map[y][x]);
    } else if (tmp.length) {
      const pos = getNearbyGearSymbolPosition(map, y, x - tmp.length, tmp);
      if (pos) {
        const number = parseInt(tmp.join(""), 10);
        if (gears[pos]) {
          gears[pos].push(number);
        } else {
          gears[pos] = [number];
        }
      }
      tmp = [];
    }
    x++;
  }
}

function run(input: string) {
  const gears: Record<string, number[]> = {};
  const map = input.split("\n").map((line) => line.split(""));
  map.forEach((_, y) => {
    findAndCollectGears(gears, map, y);
  });
  const sum = Object.keys(gears).reduce((result, key) => {
    result += gears[key].length === 2 ? gears[key][0] * gears[key][1] : 0;
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
