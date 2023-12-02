
function parse(line) {
  const requirement = {
    red: 0,
    green: 0,
    blue: 0
  };
  const colonIndex = line.indexOf(':');
  const id = parseInt(line.substring('Game'.length, colonIndex));
  const rounds = line.substr(colonIndex + 1).split(';').map(round => {
    const res = {};
    round.split(',').forEach(hand => {
      const [number, ball] = hand.trim().split(' ');
      res[ball] = parseInt(number, 10);
      if(requirement[ball] < res[ball]) {
        requirement[ball] = res[ball];
      }
    });
    return res;
  });

  return { id, rounds, requirement };
}

function computeCubesPower(game) {
  return game.requirement.red * game.requirement.green * game.requirement.blue;
}

function run(input) {
  const lines = input.split('\n');
  const sum = lines.reduce((result, line) => {
    const game = parse(line);
    result+= computeCubesPower(game)
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