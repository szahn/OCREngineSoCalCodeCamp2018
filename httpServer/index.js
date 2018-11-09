const express = require('express');
const app = express();
var cors = require('cors');
const bodyParser = require('body-parser');
const fileUpload = require('express-fileupload');
const grpcClient = require("./grpcClient");

const PORT = 8282;

app.use(cors());
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));
app.use(fileUpload());

app.post('/upload', (req, res, next) => {
    let imageFile = req.files.file;

    const destFilename = `${__dirname}/../temp/${req.body.filename}`;
    
    imageFile.mv(destFilename, function(err) {
        if (err) {
            return res.status(500).send(err);
        }

        grpcClient.OCR({
            filename:destFilename
        }, (err, response) => {
            if (err != null) {
                return res.status(500).send(err);
            }
            else{
                res.json(response);
            }
        });

    });

}); 

app.listen(8282, () => {
    console.log(`Listening on port ${PORT}`);
});

module.exports = app;