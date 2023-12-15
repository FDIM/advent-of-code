function isVerticalMirrorAt(pattern: string[][], y: number) {
  for (let x = 0; x < pattern[y].length; x++) {
    let lower = y - 1;
    let upper = y;
    while (lower >= 0 && upper < pattern.length) {
      if (pattern[y][lower] !== pattern[y][upper]) {
        return false;
      }
      lower--;
      upper++;
    }
  }
  return true;
}

function isHorizontalMirrorAt(pattern: string[][], x: number) {
  for (let y = 0; y < pattern.length; y++) {
    let lower = x - 1;
    let upper = x;
    while (lower >= 0 && upper < pattern.length) {
      if (pattern[lower][x] !== pattern[upper][x]) {
        return false;
      }
      lower--;
      upper++;
    }
  }
  return true;
}

function indexOfVerticalMirror(pattern: string[][]) {
  for (let y = 1; y < pattern.length - 1; y++) {
    if (isVerticalMirrorAt(pattern, y)) {
      return y;
    }
  }
  return 0;
}

function indexOfHorizontalMirror(pattern: string[][]) {
  for (let x = 1; x < pattern.length - 1; x++) {
    if (isHorizontalMirrorAt(pattern, x)) {
      return x;
    }
  }
  return 0;
}

function run(input: string) {
  const patterns: string[][][] = [];
  let buffer: string[][] = [];
  input.split("\n").map((line) => {
    const l = line.trim().split('');
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
  console.info(patterns);

  const sum = patterns.reduce((result, pattern) => {
    result += (indexOfHorizontalMirror(pattern) ?? 0);
    result += (indexOfVerticalMirror(pattern) ?? 0) * 100;
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

// console.info(run(require('fs').readFileSync(__dirname + '/input.txt', 'utf-8')));
