interface Coords {
  x: number;
  y: number;
}

enum MapSymbol {
  Space = ".",
  Galaxy = "#",
  Tag = '@',
}

function debugMap(map: MapSymbol[][]) {
  require("fs").appendFileSync(
    "./out.txt",
    map.map((l) => l.join("")).join("\r\n") + "\r\n\r\n",
    "utf-8"
  );
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

function expandMap(map: MapSymbol[][]) {
  for (let y = 0; y < map.length; y++) {
    if (map[y].every((i) => i === MapSymbol.Space)) {
      map.splice(y, 0, [...map[y]]);
      y += 1;
    }
  }

  for (let x = 0; x < map[0].length; x++) {
    if (map.every((i) => i[x] === MapSymbol.Space)) {
      map.forEach((row) => {
        row.splice(x, 0, MapSymbol.Space);
      });
      x += 1;
    }
  }
}

function findDistance(g1: Coords, g2: Coords) {
  return Math.max(g1.y, g2.y) - Math.min(g1.y, g2.y) + Math.max(g1.x, g2.x) - Math.min(g1.x, g2.x);
}

function run(input: string) {
  const map = input.split("\n").map((line, y) => {
    const cells = line.trim().split("");
    return cells;
  }) as MapSymbol[][];
  debugMap(map);

  expandMap(map);
  
  const galaxies = findGalaxies(map);
  galaxies.forEach(p => {
    map[p.y][p.x] = MapSymbol.Tag;
  });

  debugMap(map);
  console.info('Galaxies count:', galaxies.length);
  const pairs: Coords[][] = [];
  galaxies.forEach(g1 => {
    galaxies.forEach(g2 => {
      if (g1 !== g2 && !pairs.find(p => p.includes(g1) && p.includes(g2) )) {
        pairs.push([g1, g2]);
      }
    });
  });
  const total = pairs.reduce((res, [g1, g2]) => {
    const dist = findDistance(g1, g2);
    // console.info('pair', g1, g2, dist);
    res += dist;
    return res;
  }, 0);
  return total;
}

require("fs").writeFileSync("./out.txt", "");
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