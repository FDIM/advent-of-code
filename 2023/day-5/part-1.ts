type LookupFn = (seed: number) => number | undefined;

function createMapLookup(
  destStart: number,
  sourceStart: number,
  count: number
): LookupFn {
  return (seed: number) => {
    if (seed >= sourceStart && seed < sourceStart + count) {
      return destStart + (seed - sourceStart);
    }
  };
}

function run(input: string) {
  const lines = input.split("\n");
  const seeds = lines
    .shift()
    ?.replace("seeds: ", "")
    .split(" ")
    .map((n) => parseInt(n, 10)) as number[];
  lines.shift();
  let category: LookupFn[] = [];
  const categoryKeys: string[] = [];
  const categories: Record<string, LookupFn[]> = {};
  for (const line of lines) {
    const input = line.trim();
    if (input.includes(":")) {
      category = [];
      const key = input.replace(" map:", "");
      categoryKeys.push(key);
      categories[key] = category;
    } else {
      const [destStart, sourceStart, count] = input
        .split(" ")
        .map((n) => parseInt(n, 10));
      category.push(createMapLookup(destStart, sourceStart, count));
    }
  }
  const locations: number[] = [];
  for (const seed of seeds) {
    let value = seed;
    for (const cat of categoryKeys) {
      for (const fn of categories[cat]) {
        const target = fn(value);
        if (target) {
          value = target;
          break;
        }
      }
    }
    locations.push(value);
  }
  // console.info(seeds);
  // console.info(categoryKeys);
  // console.info(categories);
  // console.info(locations);
  return locations.sort((a, b) => a - b ).shift();
}

console.info(
  run(
    `seeds: 79 14 55 13

    seed-to-soil map:
    50 98 2
    52 50 48
    
    soil-to-fertilizer map:
    0 15 37
    37 52 2
    39 0 15
    
    fertilizer-to-water map:
    49 53 8
    0 11 42
    42 0 7
    57 7 4
    
    water-to-light map:
    88 18 7
    18 25 70
    
    light-to-temperature map:
    45 77 23
    81 45 19
    68 64 13
    
    temperature-to-humidity map:
    0 69 1
    1 0 69
    
    humidity-to-location map:
    60 56 37
    56 93 4`
  )
);

console.info(run(require('fs').readFileSync(__dirname + '/input.txt', 'utf-8')));
