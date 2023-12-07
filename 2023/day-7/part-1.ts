interface Game {
  hand: string;
  bid: number;
  handType: string;
}

const powerMap: Record<string, number> = {
  A: 14,
  K: 13,
  Q: 12,
  J: 11,
  T: 10,
  9: 9,
  8: 8,
  7: 7,
  6: 6,
  5: 5,
  4: 4,
  3: 3,
  2: 2,
  FiveOfAKind: 10,
  FourOfAKind: 9,
  FullHouse: 8,
  ThreeOfAKind: 7,
  TwoPair: 6,
  OnePair: 5,
  HighCard: 4,
};

function getHandType(hand: string) {
  const sortedHand = hand
      .split("")
      .sort((a, b) => powerMap[b] - powerMap[a])
      .join("");
  let pairsCount = 0;
  let potentialFullHouse = false;
  for (let i = 0; i < sortedHand.length; i++) {
    const identicalCardCount = sortedHand.replace(
      new RegExp(`[^${sortedHand.charAt(i)}]`, "g"),
      ""
    ).length;
    if (identicalCardCount === 5) {
      return "FiveOfAKind";
    } else if (identicalCardCount === 4) {
      return "FourOfAKind";
    } else if (identicalCardCount === 3) {
      potentialFullHouse = true;
      i += 2;
    } else if (identicalCardCount === 2) {
      pairsCount++;
      i++;
    }
  }
  if (potentialFullHouse && pairsCount == 1) {
    return "FullHouse";
  } else if (potentialFullHouse) {
    return "ThreeOfAKind";
  } else if (pairsCount === 2) {
    return "TwoPair";
  } else if (pairsCount === 1) {
    return "OnePair";
  }
  return "HighCard";
}

function compare(a: Game, b: Game) {
  const aPower = powerMap[a.handType] as number;
  const bPower = powerMap[b.handType] as number;
  if (aPower > bPower) {
    return 1;
  } else if (aPower < bPower) {
    return -1;
  }
  for (let i = 0; i < a.hand.length; i++) {
    const ap = powerMap[a.hand.charAt(i)] as number;
    const bp = powerMap[b.hand.charAt(i)] as number;
    if (ap == bp) {
      continue;
    }
    return ap - bp;
  }
  return 0;
}

function run(input: string) {
  const games = input.split("\n").map((line) => {
    const [hand, bid] = line.trim().split(" ");
    return {
      hand,
      handType: getHandType(hand),
      bid: parseInt(bid, 10),
    };
  });
  games.sort(compare);
  // console.info(games);
  const total = games.reduce((res, game, index) => {
    res += game.bid * (index + 1);
    return res;
  }, 0);
  return total;
}

console.info(
  run(
    `32T3K 765
    T55J5 684
    KK677 28
    KTJJT 220
    QQQJA 483`
  )
);

console.info(run(require('fs').readFileSync(__dirname + '/input.txt', 'utf-8')));
