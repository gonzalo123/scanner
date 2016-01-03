var socket = io('http://localhost:5000/');
var myToken;

if (typeof(Storage) !== "undefined") {
    myToken = localStorage.getItem("myToken");
    if (!myToken) {
        for (myToken = ''; myToken.length < 32;) {
            myToken += Math.random().toString(36).substr(2, 1)
        }
        myToken = (new Date()).toISOString() + myToken;
        localStorage.setItem("myToken", myToken);
    }

    socket.on(myToken, function (data) {
        console.log(data);
        document.getElementById("text").innerHTML = data.text || '';
        document.getElementById("format").innerHTML = data.format || '';
    });

    new QRCode(document.getElementById("qrcode"), myToken);

}