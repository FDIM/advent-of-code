// Returns 1 if the lines intersect, otherwise 0. In addition, if the lines
// intersect the intersection point may be stored in the floats i_x and i_y.
function getLineIntersection(
  p0_x: number,
  p0_y: number,
  p1_x: number,
  p1_y: number,
  p2_x: number,
  p2_y: number,
  p3_x: number,
  p3_y: number
) {
  let s1_x, s1_y, s2_x, s2_y;
  s1_x = p1_x - p0_x;
  s1_y = p1_y - p0_y;
  s2_x = p3_x - p2_x;
  s2_y = p3_y - p2_y;

  let s: number, t: number;
  s =
    (-s1_y * (p0_x - p2_x) + s1_x * (p0_y - p2_y)) /
    (-s2_x * s1_y + s1_x * s2_y);
  t =
    (s2_x * (p0_y - p2_y) - s2_y * (p0_x - p2_x)) /
    (-s2_x * s1_y + s1_x * s2_y);

  if (s >= 0 && s <= 1 && t >= 0 && t <= 1) {
    // Collision detected
    return {
      x: p0_x + t * s1_x,
      y: p0_y + t * s1_y,
      z: 0,
    };
  }

  return false; // No collision
}

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
  initial: Vector;
  velocity: Vector;
  start: Vector;
  end: Vector;
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
  const min = new Vector(lower, lower, NaN);
  const max = new Vector(upper, upper, NaN);
  const stones = input.split("\n").map((line) => {
    const [pos, velocityInput] = line.split("@");
    const initial = new Vector(pos);
    const velocity = new Vector(velocityInput);
    const stone: HailStone = {
      initial,
      velocity,
      start: initial,
      end: extrapolate(initial, velocity, max.x, max.y),
    };
    return stone;
  }) as HailStone[];

  // console.info(stones);
  let collisions = 0;
  for (let i = 0; i < stones.length; i++) {
    const stone1 = stones[i];
    for (let j = i + 1; j < stones.length; j++) {
      const stone2 = stones[j];
      const collisionPoint = getLineIntersection(
        stone1.start.x,
        stone1.start.y,
        stone1.end.x,
        stone1.end.y,
        stone2.start.x,
        stone2.start.y,
        stone2.end.x,
        stone2.end.y
      );
      if (
        collisionPoint &&
        collisionPoint.x >= min.x &&
        collisionPoint.x <= max.x &&
        collisionPoint.y >= min.y &&
        collisionPoint.y <= max.y
      ) {
        collisions++;
        console.info("collision", collisionPoint);
        // console.info(stone1);
        // console.info(stone2);
      }
    }
  }

  return collisions;
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

console.info(
  run(require("fs").readFileSync(__dirname + "/input.txt", "utf-8"), 200000000000000, 400000000000000)
);
