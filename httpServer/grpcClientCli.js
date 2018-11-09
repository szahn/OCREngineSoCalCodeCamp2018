const grpcClient = require("./grpcClient");

grpcClient.OCR({
    filename:"/home/szahn/go/src/ocr-engine/api/public/document.pdf"
}, (err, response)=>{
    if (err != null){
        console.log(err);
    }
    else{
        console.log(response);
    }
});
