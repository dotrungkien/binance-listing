require('dotenv').config();
const request = require('request');
const { send } = require('process');
fs = require('fs');

const getLocalMarketData = () => {
  const data = fs.readFileSync('symbols.json');
  return JSON.parse(data);
};

const saveMarketData = symbols => {
  fs.writeFile('symbols.json', JSON.stringify(symbols), function (err, data) {
    if (err) {
      return console.log(err);
    }
  });
};

// curl -X POST -H "X-ChatWorkToken: 92b891ee2edef5c24527ca0890a133df" -d "body=Hello+ChatWork%21" "https://api.chatwork.com/v2/rooms/167988546/messages"

const sendCW = message => {
  const options = {
    method: 'POST',
    url: 'https://api.chatwork.com/v2/rooms/167988546/messages',
    headers: {
      'X-ChatWorkToken': process.env.CW_TOKEN,
    },
    form: { body: message },
  };

  function callback(error, response, body) {
    if (!error && response.statusCode == 200) {
      console.log(body);
    }
  }
  request.post(options, callback);
};

const getNewData = (currentArr, newArr) => {
  if (!currentArr || !newArr) return [];
  let diff = [];
  newArr.forEach(symbol => {
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
      JSON.parse(body).symbols.forEach(s => {
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
        sendCW(`[To:899965] new symbol found: ${newSymbolsString}`);
        saveMarketData(symbols);
      }
    }
  });
};

main();
