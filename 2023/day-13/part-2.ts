require("fs").writeFileSync("./out.txt", "");

function debugMap(map: string[][], x: number, y: number) {
  require("fs").appendFileSync(
    "./out.txt",
    (x !== -1 ? "^".padStart(x + 1, " ") : "") +
      "\r\n" +
      map
        .map((l, yy) => {
          return `${yy === y ? ">" : " "}${l.join("")}${yy === y ? "<" : " "}`;
        })
        .join("\r\n") +
      "\r\n" +
      (x !== -1 ? "^".padStart(x + 1, " ") : "") +
      "\r\n\r\n",
    "utf-8"
  );
}

function isVerticalLineMirrorAt(pattern: string[][], x: number) {
  let wrong = 0;
  let count = 0;
  for (let y = 0; y < pattern.length; y++) {
    let lower = x - 1;
    let upper = x;
    while (lower >= 0 && upper < pattern[y].length) {
      if (pattern[y][lower] !== pattern[y][upper]) {
        wrong++;
      } else {
        count++;
      }
      lower--;
      upper++;
    }
  }
  return (count - wrong) === 0;
}

function isHorizontalLineMirrorAt(pattern: string[][], y: number) {
  let wrong = 0;
  let count = 0;
  for (let x = 0; x < pattern[y].length; x++) {
    let lower = y - 1;
    let upper = y;
    while (lower >= 0 && upper < pattern.length) {
      if (pattern[lower][x] !== pattern[upper][x]) {
        wrong++;
      } else {
        count++;
      }
      lower--;
      upper++;
    }
  }
  return (count - wrong) === 0;
}

function indexOfVerticalLineMirror(pattern: string[][]) {
  let lower = Math.floor(pattern[0].length / 2);
  let upper = lower + 1;
  while (lower >= 0 && upper < pattern[0].length) {
    if (isVerticalLineMirrorAt(pattern, lower)) {
      return lower;
    }
    if (isVerticalLineMirrorAt(pattern, upper)) {
      return upper;
    }
    lower--;
    upper++;
  }

  return -1;
}

function indexOfHorizontalLineMirror(pattern: string[][]) {
  let lower = Math.floor(pattern.length / 2);
  let upper = lower + 1;
  while (lower >= 0 && upper < pattern.length) {
    if (isHorizontalLineMirrorAt(pattern, lower)) {
      return lower;
    }
    if (isHorizontalLineMirrorAt(pattern, upper)) {
      return upper;
    }
    lower--;
    upper++;
  }
  return -1;
}

function run(input: string) {
  const patterns: string[][][] = [];
  let buffer: string[][] = [];
  input.split("\n").map((line) => {
    const l = line.trim().split("");
    if (l.length === 0) {
      patterns.push(buffer);
      buffer = [];
    } else {
      buffer.push(l);
    }
  });
  if (buffer.length) {
    patterns.push(buffer);
  }

  const sum = patterns.reduce((result, pattern) => {
    const horizontalIndex = indexOfHorizontalLineMirror(pattern);
    const verticalIndex = horizontalIndex === -1 ? indexOfVerticalLineMirror(pattern) : -1;
    if (verticalIndex !== -1) {
      result += verticalIndex;
    }
    if (horizontalIndex !== -1) {
      result += horizontalIndex * 100;
    }
    debugMap(pattern, verticalIndex, horizontalIndex);
    return result;
  }, 0);

  return sum;
}

console.info(
  run(
   `#.##..##.
    ..#.##.#.
    ##......#
    ##......#
    ..#.##.#.
    ..##..##.
    #.#.##.#.
    
    #...##..#
    #....#..#
    ..##..###
    #####.##.
    #####.##.
    ..##..###
    #....#..#`
  )
);

// console.info(
//   run(require("fs").readFileSync(__dirname + "/input.txt", "utf-8"))
// );
