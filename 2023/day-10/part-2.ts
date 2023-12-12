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
  My = "#",
  EnclosedGround = "I",
  NotEnclosedGround = "0",
}

// found this in the broad internet
// https://alienryderflex.com/polygon/
function intersects (x: number, y: number, coords: Coords[]) {

  let i: number, j=coords.length-1 ;
  let odd: any = false;

  // yes, this is slow but it does the job!
  let pX = coords.map(p => p.x);
  let pY = coords.map(p => p.y);

  for (i=0; i<coords.length; i++) {
      if ((pY[i]< y && pY[j]>=y ||  pY[j]< y && pY[i]>=y)
          && (pX[i]<=x || pX[j]<=x)) {
            odd ^= (pX[i] + (y-pY[i])*(pX[j]-pX[i])/(pY[j]-pY[i])) < x as any; 
      }

      j=i; 
  }

return odd;
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
  return paths[0];
}

function getInnerGround(map: MapSymbol[][], path: Coords[]) {
  let min: Coords = { x: 99999, y: 99999 };
  let max: Coords = { x: -1, y: -1 };
  path.forEach((p) => {
    if (p.x < min.x) {
      min.x = p.x;
    }
    if (p.y < min.y) {
      min.y = p.y;
    }
    if (max.x < p.x) {
      max.x = p.x;
    }
    if (max.y < p.y) {
      max.y = p.y;
    }
  });

  const ground: Coords[] = [];
  for (let y = min.y; y < max.y; y++) {
    for (let x = min.x; x < max.x; x++) {
      if (!path.find(c => c.x === x && c.y === y) && intersects(x, y, path)) {
        ground.push({ x, y });
      }
    }
  }
  return ground;
}

function debugMap(map: MapSymbol[][]) {
  require('fs').appendFileSync('./out.txt', map.map(l => l.join('')).join('\r\n') + '\r\n\r\n', 'utf-8')

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
  }) as MapSymbol[][];
  debugMap(map);
  const longestPath = findLongestPath(map, pos);
  longestPath.forEach(p => {
    map[p.y][p.x] = MapSymbol.My;
  });
  debugMap(map);
  const innerGround = getInnerGround(map, longestPath);
  innerGround.forEach(p => {
    map[p.y][p.x] = MapSymbol.EnclosedGround;
  });
 debugMap(map);

  console.info(innerGround.length);
  return "";
}

require('fs').writeFileSync('./out.txt', '');

console.info(
  run(
    `..........
    .S-------7.
    .|F-----7|.
    .||.....||.
    .||.....||.
    .|L-7.F-J|.
    .|..|.|..|.
    .L--J.L--J.
    ...........`
  )
);

console.info(
  run(
    `.........
    .S------7.
    .|F----7|.
    .||....||.
    .||....||.
    .|L-7F-J|.
    .|..||..|.
    .L--JL--J.
    ..........`
  )
);

console.info(
  run(
    `.F----7F7F7F7F-7....
    .|F--7||||||||FJ....
    .||.FJ||||||||L7....
    FJL7L7LJLJ||LJ.L-7..
    L--J.L7...LJS7F-7L7.
    ....F-J..F7FJ|L7L7L7
    ....L7.F7||L7|.L7L7|
    .....|FJLJ|FJ|F7|.LJ
    ....FJL-7.||.||||...
    ....L---J.LJ.LJLJ...`
  )
);
console.info(
  run(
    `FF7FSF7F7F7F7F7F---7
    L|LJ||||||||||||F--J
    FL-7LJLJ||||||LJL-77
    F--JF--7||LJLJ7F7FJ-
    L---JF-JLJ.||-FJLJJ7
    |F|F-JF---7F7-L7L|7|
    |FFJF7L7F-JF7|JL---7
    7-L-JL7||F7|L7F-7F7|
    L.L7LFJ|||||FJL7||LJ
    L7JLJL-JLJLJL--JLJ.L`
  )
);
console.info(run(require('fs').readFileSync(__dirname + '/input.txt', 'utf-8')));
