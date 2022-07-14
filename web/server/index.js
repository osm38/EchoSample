const express = require('express')
const web = express()
const request = require('request')
const port = process.env.PORT || 8080
const urlPrefix = 'http://app:1323'
const callback = (req, res) => {
    const options = {
        method: req.method,
        url: urlPrefix + req.originalUrl,
    };
    console.log('method:' + options['method'])
    console.log('url:' + options['url'])
    console.log('originalUrl:' + req.originalUrl)
    console.log('protocol:' + req.protocol)
    console.log('parmas:' + req.params)
    console.log('body:' + req.body)
    request(options, function(err, response, body) {
        res.send(body);
    })
}

web.use(express.json())

web.get('/', (req, res) => {
    res.send('Hello World!');
});

web.get(/.+/, callback);
web.post(/.+/, callback);

web.listen(port, () => {
    console.log(`listening on *:${port}`);
});