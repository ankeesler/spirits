const {fail} = require('assert');

const puppeteer = require('puppeteer');

let browser, page;

jest.setTimeout(1000 * 10);

beforeAll(async () => {
  const baseUrl = process.env.SPIRITS_TEST_URL;
  if (!baseUrl) {
    fail("must set 'process.env.SPIRITS_TEST_URL'");
  }

  browser = await puppeteer.launch({headless: false});
  page = await browser.newPage();
  await page.goto(baseUrl);
  page.setDefaultTimeout(30000);
});

afterAll(async () => {
  if (browser) await browser.close();
});

const generateSpiritsButton = '#generate-spirits-button';
const spiritsInput = '#spirits-text';
const battleButton = '#battle-navigation';
const startBattleButton = '#start-battle-button';
const battleOutput = '#battle-text';

test('generated spirits', async () => {
  // Click generate button.
  await page.waitForSelector(generateSpiritsButton);
  await page.click(generateSpiritsButton);

  // Make sure generated spirits show up.
  await page.waitForSelector(spiritsInput)
  await page.waitForFunction(`document.querySelector("${spiritsInput}").innerText.length > 0`);

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
