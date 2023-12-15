interface SpringRecord {
  springs: SpringType[];
  sequence: number[];
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

function getPossibleOptions(types: SpringType[], length: number): SpringType[][] {
  if (length === 0) {
    return [];
  } else if (length === 1) {
    return types.map(t => [t]);
  } else {
    const options = getPossibleOptions(types, length - 1);
    const copy: SpringType[][] = [];
    types.forEach(type => {
      options.forEach(op => {
        copy.push([...op, type]);
      });
    })
    return copy;
  }
}

function computeArrangements(record: SpringRecord): SpringType[][] {
  const locations: number[] = record.springs.reduce((res, spring, index) => {
    if (spring === SpringType.Unknown) {
      res.push(index);
    }
    return res;
  }, [] as number[]);
  
  const variants = [SpringType.Damaged, SpringType.Operational];
  const options = getPossibleOptions(variants, locations.length);
  // console.info(options);

  const arrangements = options.map(op => {
    const copy = [...record.springs];
    locations.forEach((loc, index) => {
      copy[loc] = op[index];
    });
    return copy;
  });

  return arrangements;
}

function run(input: string) {
  // console.info(
  //   checkIntegrity("....##...#.###...".split("") as SpringType[], [2, 1, 3])
  // );

  const records: SpringRecord[] = input.split("\n").map((line, y) => {
    const [springs, sequence] = line.trim().split(" ");
    return {
      springs: [springs,springs,springs,springs,springs].join('?').split("") as SpringType[],
      sequence: [sequence, sequence, sequence, sequence, sequence].join(',').split(",").map((n) => parseInt(n, 10)),
    };
  });

  

  const total = records.reduce((res, record) => {
    const regex = new RegExp(`^[\\\.]*${ record.springs.map(s => {
      return `[\#]{${s}}`
    }).join('[\\\.]+') }[\\\.]*$`);
    const arrangements = computeArrangements(record).filter(a => checkIntegrity(a, regex));
    // console.info(record.springs.join(""), record.sequence);

    // console.info(
    //   arrangements
    //     .map((a) => a.join("") +':' + checkIntegrity(a, record.sequence))
    //     .join("\n")
    // );
    // console.info("");
    res +=
      arrangements.length + 0
      // (arrangements[0]?.filter((s) => s === SpringType.Damaged).length ?? 0);
    return res;
  }, 0);
  return total;
}

console.info(new Date().toJSON());
console.info(
  run(
    `???.### 1,1,3
    .??..??...?##. 1,1,3
    ?#?#?#?#?#?#?#? 1,3,1,6
    ????.#...#... 4,1,1
    ????.######..#####. 1,6,5
    ?###???????? 3,2,1`
  )
);
console.info(new Date().toJSON());
console.info(
  run(require("fs").readFileSync(__dirname + "/input.txt", "utf-8"))
);
console.info(new Date().toJSON());
