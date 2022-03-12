const {fail} = require('assert');

const puppeteer = require('puppeteer');

let browser, page;

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

const generateButton = '.component-spirit-window > button';
const spiritsInput = '.component-spirit-window > div';
const battleButton = '.component-navigation > button + button';
const battleOutput = '.component-battle-screen';

test('generated spirits', async () => {
  // Click generate button.
  await page.waitForSelector(generateButton);
  await page.click(generateButton);

  // Make sure generated spirits show up.
  await page.waitForSelector(spiritsInput)
  await page.waitForFunction(`document.querySelector("${spiritsInput}").innerText.length > 0`);

  // Click on battle button.
  await page.waitForSelector(battleButton);
  await page.click(battleButton);

  // Get battle output.
  await page.waitForSelector(battleOutput);
  await page.waitForFunction(`document.querySelector("${battleOutput}").innerText.length > 0`);
  const output = await page.$eval(battleOutput, (e) => e.innerText);
  expect(output).toMatch(/^> summary/);
});
