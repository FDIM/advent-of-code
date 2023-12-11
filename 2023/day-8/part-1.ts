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

  let node: MapNode = nodes["AAA"];
  let steps = 0;
  for (let i = 0; i < instructions.length; i++) {
    const dir = instructions[i];
    node = nodes[node[dir]];
    steps++;
    if (node.name === "ZZZ") {
      break;
    }
    if (i + 1 === instructions.length) {
      i = -1; // restart loop
    }
  }
  return steps;
}

console.info(
  run(
    `RL

    AAA = (BBB, CCC)
    BBB = (DDD, EEE)
    CCC = (ZZZ, GGG)
    DDD = (DDD, DDD)
    EEE = (EEE, EEE)
    GGG = (GGG, GGG)
    ZZZ = (ZZZ, ZZZ)`
  )
);

console.info(
  run(`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`)
);

console.info(
  run(require("fs").readFileSync(__dirname + "/input.txt", "utf-8"))
);
