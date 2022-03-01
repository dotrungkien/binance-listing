require('dotenv').config();
const request = require('request');
fs = require('fs');

const { Telegraf } = require('telegraf');

const bot = new Telegraf(process.env.TELE_KEY);

const getLocalMarketData = () => {
  const data = fs.readFileSync('symbols.json');
  return JSON.parse(data);
};

const saveMarketData = (symbols) => {
  fs.writeFile('symbols.json', JSON.stringify(symbols), function (err, data) {
    if (err) {
      return console.log(err);
    }
  });
};

const sendMessage = (message) => {
  bot.telegram.sendMessage(process.env.CHAT_ID, message);
};

const getNewData = (currentArr, newArr) => {
  if (!currentArr || !newArr) return [];
  let diff = [];
  newArr.forEach((symbol) => {
    if (!currentArr.includes(symbol)) {
      diff.push(symbol);
    }
  });
  return diff;
};

const main = () => {
  const url = 'https://api.binance.com/api/v3/exchangeInfo';
  request.get(url, (error, response, body) => {
    if (response.statusCode === 200) {
      let symbols = [];
      JSON.parse(body).symbols.forEach((s) => {
        const symbol = s.symbol;
        if (symbol.slice(-4) === 'USDT') symbols.push(symbol);
      });

      const currentSymbols = getLocalMarketData();
      let newSymbols = getNewData(currentSymbols, symbols);
      if (newSymbols.length === 0) {
        console.log('no new symbol');
      } else {
        console.log('new symbol found!');
        let newSymbolsString = newSymbols.join(' ');
        sendMessage(`New symbol found: ${newSymbolsString}`);
        saveMarketData(symbols);
      }
    }
  });
};

main();
