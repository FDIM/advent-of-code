interface Coords {
  x: number;
  y: number;
}

interface EmptySpace{
  cols: number[];
  rows: number[];
}

enum MapSymbol {
  Space = ".",
  Galaxy = "#",
  Tag = "@",
}

function findGalaxies(map: MapSymbol[][]) {
  const coords: Coords[] = [];
  for (let y = 0; y < map.length; y++) {
    for (let x = 0; x < map[y].length; x++) {
      if (map[y][x] === MapSymbol.Galaxy) {
        coords.push({ x, y });
      }
    }
  }
  return coords;
}

function findEmptySpaces(map: MapSymbol[][]) {
  const empty: EmptySpace = {
    cols: [],
    rows: [],
  };
  for (let y = 0; y < map.length; y++) {
    if (map[y].every((i) => i === MapSymbol.Space)) {
      empty.rows.push(y);
    }
  }
  for (let x = 0; x < map[0].length; x++) {
    if (map.every((i) => i[x] === MapSymbol.Space)) {
      empty.cols.push(x);
    }
  }
  return empty;
}

function findDistance(g1: Coords, g2: Coords, empty: EmptySpace, scale: number) {
  const minX = Math.min(g1.x, g2.x);
  const maxX = Math.max(g1.x, g2.x);
  const minY = Math.min(g1.y, g2.y);
  const maxY = Math.max(g1.y, g2.y);
  let distance = (
    Math.max(g1.y, g2.y) -
    Math.min(g1.y, g2.y) +
    Math.max(g1.x, g2.x) -
    Math.min(g1.x, g2.x)
  );
  empty.cols.forEach(col => {
    if(minX < col && maxX > col) {
      distance += (scale - 1);
    }
  });
  empty.rows.forEach(row => {
    if(minY < row && maxY > row) {
      distance += (scale - 1);
    }
  });
  return distance;
}

function run(input: string) {
  const map = input.split("\n").map((line, y) => {
    const cells = line.trim().split("");
    return cells;
  }) as MapSymbol[][];

  const galaxies = findGalaxies(map);
  galaxies.forEach((p) => {
    map[p.y][p.x] = MapSymbol.Tag;
  });

  console.info("Galaxies count:", galaxies.length);
  const pairs: Coords[][] = [];
  galaxies.forEach((g1) => {
    galaxies.forEach((g2) => {
      if (g1 !== g2 && !pairs.find((p) => p.includes(g1) && p.includes(g2))) {
        pairs.push([g1, g2]);
      }
    });
  });

  const empty = findEmptySpaces(map);
  const scale = 1000000;
  const total = pairs.reduce((res, [g1, g2]) => {
    const dist = findDistance(g1, g2, empty, scale);
    // console.info('pair', g1, g2, dist);
    res += dist;
    return res;
  }, 0);
  return total;
}

// require("fs").writeFileSync("./out.txt", "");

console.info(new Date().toJSON());
console.info(
  run(
    `...#......
    .......#..
    #.........
    ..........
    ......#...
    .#........
    .........#
    ..........
    .......#..
    #...#.....`
  )
);
console.info(new Date().toJSON());
console.info(
  run(require("fs").readFileSync(__dirname + "/input.txt", "utf-8"))
);
console.info(new Date().toJSON());
