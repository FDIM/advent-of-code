interface Card {
  winning: string[];
  my: string[]
}

function parse(line: string): Card {
  let [winningInput, myInput] = line.split("|");
  const winningNumbers = winningInput
    .substring(winningInput.indexOf(":") + 1)
    .split(" ")
    .filter((n) => n !== "");
  const myNumbers = myInput
    .trim()
    .split(" ")
    .filter((n) => n !== "");
  return {
    winning: winningNumbers,
    my: myNumbers,
  };
}

function computeNumberOfWinningCards(card: Card) {
  let score = 0;

  card.my.forEach((num) => {
    if (card.winning.includes(num)) {
      score++;
    }
  });
  return score;
}

function checkCard(
  index: number,
  cards: Card[],
  copies: Record<number, number>
) {
  let winning = computeNumberOfWinningCards(cards[index]);
  if (winning > 0) {
    while (winning > 0) {
      if (!copies[index + winning]) {
        copies[index + winning] = 0;
      }
      copies[index + winning]++;
      winning--;
    }
  }
}

function run(input: string) {
  const cards = input.split("\n").map(parse);
  const copies: Record<number, number> = {};
  let extraCopies = 0;
  let i = 0;

  while (i < cards.length) {
    checkCard(i, cards, copies);

    if (copies[i]) {
      extraCopies += copies[i];
      for (let j = 0; j < copies[i]; j++) {
        checkCard(i, cards, copies);
      }
    }
    i++;
  }
  return cards.length + extraCopies;
}

console.info(
  run(
    `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
  )
);

const start = new Date().getTime();
console.info(
  run(require("fs").readFileSync(__dirname + "/input.txt", "utf-8"))
);
console.info("Took", new Date().getTime() - start);
