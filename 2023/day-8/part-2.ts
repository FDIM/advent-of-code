interface MapNode {
  name: string;
  left: string;
  right: string;
}

function run(input: string) {
  const lines = input.split("\n");
  const instructions =
    lines
      .shift()
      ?.trim()
      .split("")
      .map((d) => (d === "R" ? "right" : "left")) ?? [];
  lines.shift();
  const nodes: Record<string, MapNode> = {};
  lines.map((line) => {
    line = line.replace(/\s/g, "");
    const index = line.indexOf("=");
    const name = line.substring(0, index).trim();
    const [left, right] = line.substring(index + 2, line.length - 1).split(",");
    nodes[name] = { name, left, right };
  });

  const current: MapNode[] = Object.keys(nodes)
    .filter((n) => n.endsWith("A"))
    .map((n) => nodes[n]);
  let steps = 0;
  for (let i = 0; i < instructions.length; i++) {
    const dir = instructions[i];
    for (let n = 0; n < current.length; n++) {
      const node = current[n];
      current[n] = nodes[node[dir]];
    }
    steps++;
    // if (current.filter((n) => n.name.endsWith("Z")).length === 4) {
    //   console.info(4, current);
    // }
    // if (current.filter((n) => n.name.endsWith("Z")).length === 5) {
    //   console.info(5, current);
    // }
    if (current.every((n) => n.name.endsWith("Z"))) {
      break;
    }
    if (i + 1 === instructions.length) {
      i = -1; // restart loop
    }
  }
  return steps;
}
console.info(new Date().toJSON());
console.info(
  run(
    `LR

    11A = (11B, XXX)
    11B = (XXX, 11Z)
    11Z = (11B, XXX)
    22A = (22B, XXX)
    22B = (22C, 22C)
    22C = (22Z, 22Z)
    22Z = (22B, 22B)
    XXX = (XXX, XXX)`
  )
);
console.info(new Date().toJSON());

console.info(
  run(require("fs").readFileSync(__dirname + "/input.txt", "utf-8"))
);
console.info(new Date().toJSON());