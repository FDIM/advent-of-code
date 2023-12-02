const max = {
  red: 12,
  green: 13,
  blue: 14
};

function parse(line) {
  const colonIndex = line.indexOf(':');
  const id = parseInt(line.substring('Game'.length, colonIndex));
  const rounds = line.substr(colonIndex + 1).split(';').map(round => {
    const res = {};
    round.split(',').forEach(hand => {
      const [number, ball] = hand.trim().split(' ');
      res[ball] = parseInt(number, 10);
    });
    return res;
  });

  return { id, rounds };
}

function isGamePossible(game) {
  return game.rounds.every(round => {
    return (!round.red || round.red <= max.red) && (!round.green || round.green <= max.green) && (!round.blue || round.blue <= max.blue) 
  });
}

function run(input) {
  const lines = input.split('\n');
  const sum = lines.reduce((result, line) => {
    const game = parse(line);
    result+= isGamePossible(game) ? game.id: 0;
    return result;
  }, 0);

  return sum;
}

console.info(run(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`));

console.info(run(require('fs').readFileSync(__dirname + '/input.txt', 'utf-8')));