const fs = require('fs');
const path = require('path');

const puppeteer = require('puppeteer');

const server = require('./server');

const fixture = (name) => {
  return fs.readFileSync(path.join(__dirname, 'fixture', name), {encoding: 'utf-8'});
};

let browser, page;

// Worst case scenario we have to build the whole container and wait for it to start.
jest.setTimeout(1000 * 10);

beforeAll(async () => {
  const baseUrl = await server.start((details) => {
    if (details.code !== 0) {
      console.log(`server errored: ${JSON.stringify(details, null, 2)}`);
    }
  });
  browser = await puppeteer.launch();
  page = await browser.newPage();
  await page.goto(baseUrl);
});

afterAll(async () => {
  server.stop();
  await browser.close();
});

test('can generate spirtis', async () => {
  const inputSelector = '.component-spirit-input';
  const outputSelector = '.component-battle-console';

  await page.waitForSelector(inputSelector);
  await page.type(inputSelector, fixture('good-spirits.json'));

  const getOutput = async () => {
    return await page.$eval(outputSelector, e => e.value);
  };
  const waitForOutput = async (length) => {
    await page.waitForFunction(`document.querySelector("${outputSelector}").value.length > ${length}`);
  };
  await page.waitForSelector(outputSelector);
  await waitForOutput(0);
  expect(await getOutput()).toEqual('⏳');
  await waitForOutput(10);
  expect(await getOutput()).toEqual(fixture('good-spirits.txt'));
});
