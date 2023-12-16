interface SpringRecord {
  springs: SpringType[];
  sequence: number[];
  regex: RegExp;
}

enum SpringType {
  Damaged = "#",
  Operational = ".",
  Unknown = "?",
}

function checkIntegrity(springs: SpringType[], regex: RegExp) {
  
  return regex.test(springs.join(''));
  // const seq = sequence.slice();
  // let i = 0;
  // let num = seq.shift();
  // let valid = true;
  // while (i < springs.length && valid) {
  //   const spring = springs[i];
  //   if (num === 0) {
  //     num = seq.shift();
  //     valid = spring === SpringType.Operational;
  //   } else if (num && num > 0 && spring === SpringType.Damaged) {
  //     num--;
  //   } else {
  //     valid = spring === SpringType.Operational || spring === undefined;
  //   }
  //   i++;
  // }
  // return valid && springs.length > 0;
}

function iteratePossibleOptions(variants: SpringType[], length: number, callback: (variant: SpringType[]) => void, value?: SpringType[]) {
  if (length === 0) {
    callback(value ?? []);
    return;
  }

  for(const variant of variants) {
    const copy = [...value ?? [], variant];
    iteratePossibleOptions(variants, length -1, callback, copy)
  };
}

function computeArrangements(record: SpringRecord, callback: (arrangement: SpringType[]) => void) {
  const locations: number[] = record.springs.reduce((res, spring, index) => {
    if (spring === SpringType.Unknown) {
      res.push(index);
    }
    return res;
  }, [] as number[]);

  const variants = [SpringType.Damaged, SpringType.Operational];

  const copy = [...record.springs];
  iteratePossibleOptions(variants, locations.length, (variant) =>{
    for(let i = 0; i< locations.length; i++) {
      copy[locations[i]] = variant[i];
    }
    if (checkIntegrity(copy, record.regex)) {
      callback(copy);
    }
  });

}

function run(input: string) {

  const records: SpringRecord[] = input.split("\n").map((line, y) => {
    const [springs, sequence] = line.trim().split(" ");
    const parsedSequence = sequence.split(",").map((n) => parseInt(n, 10));
    return {
      springs: springs.split("") as SpringType[],
      sequence: parsedSequence,
      regex: new RegExp(
        `^[\\\.]*${parsedSequence
          .map((s) => {
            return `[\#]{${s}}`;
          })
          .join("[\\.]+")}[\\\.]*$`
      ),
    };
  });

  if (process.argv.length <= 2) {
    require("fs").writeFileSync("./out.txt", "");

    const spawn = require("child_process").spawn;
    const chunkSize = Math.ceil(records.length / 8);
    for (let i = 0; i < records.length; i += chunkSize) {
      spawn("ts-node.cmd", ["./part-1.ts", i, i + chunkSize], { stdio: "inherit" });
    }
    return;
  }

  let end = parseInt(process.argv.pop() ?? "", 10);
  let start = parseInt(process.argv.pop() ?? "", 10);
  console.info('processing from', start, 'to', end);
  const total = records.slice(start, end).reduce((res, record, index) => {
    let arrangementsCount = 0;
    computeArrangements(record, () => {
      arrangementsCount++;
    });
    require("fs").appendFileSync(
      "./out.txt",
      `${record.springs.join("")} | ${record.sequence.join(",")} | ${
        arrangementsCount
      } | ${Math.round(index / records.length * 100)}%\r\n`,
      "utf-8"
    );
    res += arrangementsCount;
    return res;
  }, 0);
  return total;
}

// console.info(new Date().toJSON());
// console.info(
//   run(
//     `???.### 1,1,3
//     .??..??...?##. 1,1,3
//     ?#?#?#?#?#?#?#? 1,3,1,6
//     ????.#...#... 4,1,1
//     ????.######..#####. 1,6,5
//     ?###???????? 3,2,1`
//   )
// );
console.info(new Date().toJSON());
console.info(
  run(require("fs").readFileSync(__dirname + "/input.txt", "utf-8"))
);
console.info(new Date().toJSON());
