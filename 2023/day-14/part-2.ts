require("fs").writeFileSync("./out.txt", "");

enum MapSymbol {
  RoundRock = "O",
  CubeRock = "#",
  Ground = ".",
}

enum Direction {
  North,
  West,
  South,
  East,
}

function debugMap(map: MapSymbol[][]) {
  require("fs").appendFileSync(
    "./out.txt",
    map
      .map((l) => {
        return `${l.join("")}`;
      })
      .join("\r\n") + "\r\n\r\n",
    "utf-8"
  );
}

function rollRock(map: MapSymbol[][], dir: Direction, x: number, y: number) {
  let diffX = 0;
  let diffY = 0;
  switch (dir) {
    case Direction.North:
      diffY--;
      break;
    case Direction.South:
      diffY++;
      break;
    case Direction.West:
      diffX--;
      break;
    case Direction.East:
      diffX++;
      break;
  }

  while (map[y + diffY]?.[x + diffX] === MapSymbol.Ground) {
    map[y + diffY][x + diffX] = MapSymbol.RoundRock;
    map[y][x] = MapSymbol.Ground;
    y += diffY;
    x += diffX;
  }
}

function rollAllRocks(map: MapSymbol[][], dir: Direction) {
  let startY = 0;
  let endY = 0;
  let diffY = 0;
  let startX = 0;
  let endX = 0;
  let diffX = 0;

  switch (dir) {
    case Direction.North:
      startY = 0;
      endY = map.length;
      diffY++;
      startX = 0;
      endX = map[0].length;
      diffX++;
      break;
    case Direction.South:
      startY = map.length - 1;
      endY = 0;
      diffY--;
      startX = 0;
      endX = map[0].length;
      diffX++;
      break;
    case Direction.West:
      startY = 0;
      endY = map.length;
      diffY++;
      startX = 0;
      endX = map[0].length;
      diffX++;
      break;
    case Direction.East:
      startY = 0;
      endY = map.length;
      diffY++;
      startX = map[0].length - 1;
      endX = 0;
      diffX--;
      break;
  }

  for (let y = startY; diffY < 0 ? y >= endY : y < endY; y += diffY) {
    for (let x = startX; diffX < 0 ? x >= endX : x < endX; x += diffX) {
      if (map[y][x] === MapSymbol.RoundRock) {
        rollRock(map, dir, x, y);
      }
    }
  }
}

function computeLoad(map: MapSymbol[][]) {
    let load = 0;
    for (let y = 0; y < map.length; y++) {
        for (let x = 0; x < map[y].length; x++) {
          if (map[y][x] === MapSymbol.RoundRock) {
            load += map.length - y;
          }
        }
      }
    return load;
}

function run(input: string) {
  const map = input.split("\n").map((line) => {
    return line.trim().split("");
  }) as MapSymbol[][];

  debugMap(map);
  let cycles = 200000;
  let logpoints = 500;
  let increment = cycles / logpoints;
  let counter = 1;
  let loads: number[] = [];
  require("fs").writeFileSync("./loads2.txt", "");

  for (let i = 0; i < cycles; i++) {
    rollAllRocks(map, Direction.North);
    rollAllRocks(map, Direction.West);
    rollAllRocks(map, Direction.South);
    rollAllRocks(map, Direction.East);
    loads.push(computeLoad(map));
    if (i === (increment * counter)) {
      console.info('Reached', Math.round(i / cycles * 100) + '%');
      counter++;
      require("fs").appendFileSync("./loads2.txt", loads.join(', ')+'\r\n');

      loads = [];
    }
  }

  debugMap(map);
  const load = computeLoad(map);

  
  return load;
}

console.info(new Date().toJSON());
// console.info(
//   run(
//     `O....#....
//     O.OO#....#
//     .....##...
//     OO.#O....O
//     .O.....O#.
//     O.#..O.#.#
//     ..O..#O..O
//     .......O..
//     #....###..
//     #OO..#....`
//   )
// );
// console.info(new Date().toJSON());
console.info(
  run(require("fs").readFileSync(__dirname + "/input.txt", "utf-8"))
);
console.info(new Date().toJSON());
