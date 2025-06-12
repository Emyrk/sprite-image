// const { createCanvas, loadImage } = require('canvas');
// const fs = require('fs');

// // create a 512x512 canvas
// const width = 512;
// const height = 512;
// const canvas = createCanvas(width, height);
// const ctx = canvas.getContext('2d');

// // draw something simple
// ctx.fillStyle = '#ffffff';
// ctx.fillRect(0, 0, width, height);

// ctx.fillStyle = '#ff0000';
// ctx.fillRect(100, 100, 300, 300);

// ctx.font = '48px sans-serif';
// ctx.fillStyle = '#000000';
// ctx.fillText('Hello!', 180, 300);

// // save to file
// const buffer = canvas.toBuffer('image/png');
// fs.writeFileSync('./output.png', buffer);

// console.log('Image saved!');


const puppeteer = require('puppeteer');
const fs = require('fs');

(async () => {
  const browser = await puppeteer.launch();
  const page = await browser.newPage();

  await page.goto('https://liberatedpixelcup.github.io/Universal-LPC-Spritesheet-Character-Generator/#?body=Body_color_light&head=Human_male_light'); // <-- your website

  const id = "#previewAnimations"

  // Wait for the div to appear
  await page.waitForSelector(id); // change selector

  // Get element handle
  const element = await page.$(id);

  // Screenshot the element
  const screenshot = await element.screenshot({ path: 'div-screenshot.png' });

  console.log('Div captured!');
  // const buffer = canvas.toBuffer('image/png');
  fs.writeFileSync('./output.png', screenshot);
  await browser.close();
})();