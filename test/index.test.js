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

const spiritsWindow = '#spirit-window';
const generateSpiritsButton = spiritsWindow + ' button';
const battleButton = '#show-battle-window-button';
const startBattleButton = `#start-battle-button`;
const battleOutput = '#battle-text';

test('generated spirits', async () => {
  // Click generate button.
  await page.waitForSelector(generateSpiritsButton);
  await page.click(generateSpiritsButton);

  // Make sure generated spirits show up.
  await page.waitForFunction(`document.querySelector("${spiritsWindow}").childElementCount > 1`);

  // Click on battle button.
  await page.waitForSelector(battleButton);
  await page.click(battleButton);

  // Click on start battle.
  await page.waitForSelector(startBattleButton);
  await page.click(startBattleButton);

  // Get battle output.
  await page.waitForSelector(battleOutput);
  await page.waitForFunction(`document.querySelector("${battleOutput}").innerText.length > 0`);
  const output = await page.$eval(battleOutput, (e) => e.innerText);
  expect(output).toMatch(/^> summary/);
});
