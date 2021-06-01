const puppeteer = require('puppeteer');
const fs = require('fs-extra');

(async () => {
  const browser = await puppeteer.launch({ headless: true });
  const page = await browser.newPage();

  const html = fs.readFileSync(`${__dirname}/index.html`, 'utf8');
  await page.setContent(html, { waitUntil: 'domcontentloaded' });

  // const pdfBuffer = await page.pdf({ format: 'A4' });

  await page.pdf({
    path: `${__dirname}/BurgessCoverLetter.pdf`,
    format: 'A4',
    printBackground: true,
  });

  await browser.close();
})();
