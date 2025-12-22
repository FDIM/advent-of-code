class Vector {
  x: number;
  y: number;
  z: number;

  constructor(x: number, y: number, z: number);
  constructor(input: string);
  constructor(input: string | number, y?: number, z?: number) {
    if (typeof input === "string") {
      const [x, y, z] = input.split(",");
      this.x = parseInt(x, 10);
      this.y = parseInt(y, 10);
      this.z = parseInt(z, 10);
    } else {
      this.x = input;
      this.y = y!;
      this.z = z!;
    }
  }
}

interface HailStone {
  start: Vector;
  velocity: Vector;
}

function extrapolate(
  start: Vector,
  velocity: Vector,
  x: number,
  y: number
): Vector {
  const time = Math.min(Math.abs(x / velocity.x), Math.abs(y / velocity.y));
  return new Vector(
    start.x + velocity.x * time,
    start.y + velocity.y * time,
    NaN
  );
}

function run(input: string, lower: number, upper: number) {
  const stones = input.split("\n").map((line) => {
    const [pos, velocityInput] = line.split("@");
    const start = new Vector(pos);
    const velocity = new Vector(velocityInput);
    const stone: HailStone = {
      start,
      velocity,
    };
    return stone;
  }) as HailStone[];

  let point = new Vector(0,0,0);
  
  console.info(stones);

  return point.x + point.y + point.z;
}

console.info(
  run(
    `19, 13, 30 @ -2,  1, -2
   18, 19, 22 @ -1, -1, -2
   20, 25, 34 @ -2, -2, -4
   12, 31, 28 @ -1, -2, -1
   20, 19, 15 @  1, -5, -3`
  ,7, 27)
);

// console.info(
//   run(require("fs").readFileSync(__dirname + "/input.txt", "utf-8"), 200000000000000, 400000000000000)
// );
