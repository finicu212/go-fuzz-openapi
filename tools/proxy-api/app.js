const { createProxyMiddleware } = require('http-proxy-middleware');
const express = require('express');

const app = express()

const myProxy = createProxyMiddleware({
  target: 'https://petstore3.swagger.io/api/v3/',
  changeOrigin: true
});

const proxyDelay = function (req, res, next) {
  const delay = req.headers["x-proxy-delay"];
  if (delay) {
    const delayTime = parseDelayTime(delay);

    // Send JSON response to the user immediately
    res.status(200).json({ href: req.originalUrl, status: "delayed" });

    setTimeout(() => {
      next();
    }, delayTime);
  } else {
    next();
  }
};

function parseDelayTime(delay) {
  const unit = delay.slice(-1).toLowerCase();
  const time = parseInt(delay.slice(0, -1));
  switch (unit) {
    case 's':
      return time * 1000;
    case 'm':
      return time * 60 * 1000;
    case 'h':
      return time * 60 * 60 * 1000;
    default:
      return time;
  }
}

const port = process.argv[2] || 3000; // check if port is provided as argument, otherwise use 3000 as default
app.use('/', proxyDelay, myProxy);
app.listen(port, () => {
  console.log(`Server listening on port ${port}`);
});
