const {fail} = require('assert');
const fs = require('fs');
const path = require('path');

const puppeteer = require('puppeteer');

let browser, page;

const fixture = (name) => {
  return fs.readFileSync(path.join(__dirname, 'fixture', name), {encoding: 'utf-8'});
};

// Worst case scenario we have to build the whole container and wait for it to start.
jest.setTimeout(1000 * 10);

beforeAll(async () => {
  const baseUrl = process.env.SPIRITS_TEST_URL;
  if (!baseUrl) {
    fail("must set 'process.env.SPIRITS_TEST_URL'");
  }

  browser = await puppeteer.launch();
  page = await browser.newPage();
  await page.goto(baseUrl);
});

afterAll(async () => {
  if (browser) await browser.close();
});

test('can run a battle', async () => {
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
