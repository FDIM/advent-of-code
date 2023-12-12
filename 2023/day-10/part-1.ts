interface Coords {
  x: number;
  y: number;
}

enum MapSymbol {
  Vertical = "|", // is a vertical pipe connecting north and south.
  Horizontal = "-", //is a horizontal pipe connecting east and west.
  NE = "L", // is a 90-degree bend connecting north and east.
  NW = "J", // is a 90-degree bend connecting north and west.
  SW = "7", // is a 90-degree bend connecting south and west.
  SE = "F", // is a 90-degree bend connecting south and east.
  Ground = ".",
  Start = "S",
}

function findPath(map: MapSymbol[][], start: Coords, dir: Coords) {
  let current = { x: start.x + dir.x, y: start.y + dir.y };
  let segments: Coords[] = [{ ...start }];
  let previous = { ...start };
  let tile = map[current.y]?.[current.x];
  let diff = { x: 0, y: 0 };
  while (tile && tile !== MapSymbol.Ground) {
    segments.push({ ...current });
    diff.x = 0;
    diff.y = 0;
    if (tile == MapSymbol.Vertical) {
      diff.y += previous.y < current.y ? 1 : -1;
    } else if (tile === MapSymbol.Horizontal) {
      diff.x += previous.x < current.x ? 1 : -1;
    } else if (tile === MapSymbol.NE) {
      if (previous.x > current.x) {
        diff.y--;
      } else {
        diff.x++;
      }
    } else if (tile === MapSymbol.NW) {
      if (previous.x < current.x) {
        diff.y--;
      } else {
        diff.x--;
      }
    } else if (tile === MapSymbol.SE) {
      if (previous.x > current.x) {
        diff.y++;
      } else {
        diff.x++;
      }
    } else if (tile === MapSymbol.SW) {
      if (previous.x < current.x) {
        diff.y++;
      } else {
        diff.x--;
      }
    }
    previous.x = current.x;
    previous.y = current.y;
    current.x += diff.x;
    current.y += diff.y;
    tile = map[current.y]?.[current.x];
    if (tile === MapSymbol.Start) {
      segments.push({ ...current });
      return segments;
    }
  }
  return [];
}

function findLongestPath(map: MapSymbol[][], start: Coords) {
  let paths: Coords[][] = [];
  paths.push(findPath(map, start, { x: -1, y: 0 }));
  paths.push(findPath(map, start, { x: 1, y: 0 }));
  paths.push(findPath(map, start, { x: 0, y: -1 }));
  paths.push(findPath(map, start, { x: 0, y: 1 }));
  paths.sort((a, b) => b.length - a.length);
  // console.info(map);
  // console.info(start);
  // console.info(paths);
  return (paths[0].length - 1) / 2;
}

function run(input: string) {
  const pos: Coords = { x: 0, y: 0 };
  const map = input.split("\n").map((line, y) => {
    const cells = line.trim().split("");
    const x = cells.indexOf(MapSymbol.Start);
    if (x !== -1) {
      pos.x = x;
      pos.y = y;
    }
    return cells;
  });
  return findLongestPath(map as MapSymbol[][], pos);
}

console.info(
  run(
    `.....
    .S-7.
    .|.|.
    .L-J.
    .....`
  )
);
console.info(
  run(
    `..F7.
    .FJ|.
    SJ.L7
    |F--J
    LJ...`
  )
);

console.info(run(require('fs').readFileSync(__dirname + '/input.txt', 'utf-8')));
