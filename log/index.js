var joke = require('joke')();

module.exports = joke.pipe(joke.stringify()).pipe(process.stdout);
