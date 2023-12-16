require("fs").writeFileSync("./out.txt", "");

enum MapSymbol {
  RoundRock = "O",
  CubeRock = "#",
  Ground = ".",
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

function rollRockNorth(map: MapSymbol[][], x: number, y: number) {
  while (map[y - 1]?.[x] === MapSymbol.Ground) {
    map[y - 1][x] = MapSymbol.RoundRock;
    map[y][x] = MapSymbol.Ground;
    y--;
  }
}

function rollAllRocksNorth(map: MapSymbol[][]) {
  for (let y = 0; y < map.length; y++) {
    for (let x = 0; x < map[y].length; x++) {
      if (map[y][x] === MapSymbol.RoundRock) {
        rollRockNorth(map, x, y);
      }
    }
  }
}

function run(input: string) {
  const map = input.split("\n").map((line) => {
    return line.trim().split("");
  }) as MapSymbol[][];

  debugMap(map);
  rollAllRocksNorth(map);
  
  debugMap(map);
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

console.info(
  run(
    `O....#....
    O.OO#....#
    .....##...
    OO.#O....O
    .O.....O#.
    O.#..O.#.#
    ..O..#O..O
    .......O..
    #....###..
    #OO..#....`
  )
);

console.info(
  run(require("fs").readFileSync(__dirname + "/input.txt", "utf-8"))
);
