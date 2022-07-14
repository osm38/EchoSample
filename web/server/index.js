const express = require('express')
const web = express()
const request = require('request')
const port = process.env.PORT || 8080
const urlPrefix = 'http://app:1323'

web.get('/', (req, res) => {
    res.send('Hello World!');
});

web.get(/.+/, (req, res) => {
    const options = {
        method: req.method,
        url: urlPrefix + req.path,
    };
    console.log(options['method'])
    console.log(options['url'])
    request(options, function(err, response, body) {
        res.send(body);
    })
})

web.listen(port, () => {
    console.log(`listening on *:${port}`);
});