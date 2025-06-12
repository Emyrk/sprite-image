const express = require('express');
const puppeteer = require('puppeteer');
const querystring = require('querystring');

const app = express();
const PORT = 3000;
// https://crates.io/crates/lpcg
// https://liberatedpixelcup.github.io/Universal-LPC-Spritesheet-Character-Generator

app.get('/render.png', async (req, res) => {
  const q = querystring.stringify(req.query)
  const url = `https://liberatedpixelcup.github.io/Universal-LPC-Spritesheet-Character-Generator/#?body=Body_color_light&head=Human_male_light&${q}`;

  console.log("New request", q)
  console.log(`Rendering: ${url}`);

  let browser;
  try {
    browser = await puppeteer.launch({
      headless: 'new', // newer headless mode
    //   args: ['--no-sandbox', '--disable-setuid-sandbox'] // safer for some envs
    });

    const page = await browser.newPage();
    await page.goto(url, { waitUntil: 'networkidle2', timeout: 60000 });

    // Optional: wait for canvas or some selector you care about
    await page.waitForSelector('canvas');

    // Capture full page, or target specific element
    const element = await page.$('canvas');  // <-- capture only the canvas
    const screenshotBuffer = await element.screenshot({ type: 'png' });

    res.set('Content-Type', 'image/png');
    res.send(screenshotBuffer);
  } catch (err) {
    console.error(err);
    res.status(500).send('Error generating image');
  } finally {
    if (browser) await browser.close();
  }
});

app.listen(PORT, () => {
  console.log(`Server is running at http://localhost:${PORT}`);
});
