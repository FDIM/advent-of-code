import { Worker, isMainThread, parentPort, workerData } from "worker_threads";

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
  const steps = 0;
  return steps;
}

if (isMainThread) {
  const worker = new Worker(__filename.replace('.ts', '.js'), { workerData: ['foo'] });
  // worker.on("message", msg => {
  //   if (msg.type === "update") {
  //     console.log(arr);
  //   }
  // });
  // worker.postMessage({type: "init", arr});
}
else {
  // parentPort?.on("message", msg => {
    // if (msg.type === "init") {
    //   msg.arr[0] = 1001;
    //   parentPort.postMessage({type: "update"});
    // }
  // });
  console.info(workerData);
}

// console.info(
//   run(require("fs").readFileSync(__dirname + "/input.txt", "utf-8"))
// );
// console.info(new Date().toJSON());