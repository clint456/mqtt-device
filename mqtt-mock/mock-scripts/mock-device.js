function getRandomFloat(min, max) {
    return Math.random() * (max - min) + min;
}

const deviceName = "my-mqtt-device";
let message = "test-message";
let json = {"name" : "My JSON"};

schedule('*/15 * * * * *', ()=>{
    var data = {};
    data.randnum = getRandomFloat(25,29).toFixed(1);
    data.ping = "pong"
    data.message = "Hello World "

    publish( 'incoming/data/devices-topic/values', JSON.stringify(data));
});


subscribe( "command/devices-topic/#" , (topic, val) => {
    const words = topic.split('/');
    var cmd = words[2];
    var method = words[3];
    var uuid = words[4];
    var response = {};
    var data = val;

    if (method == "set") {
        switch(cmd) {
            case "message":
                message = data[cmd];
                break;
            case "json":
                json = data[cmd];
                break;
        }
    }else{
        switch(cmd) {
            case "ping":
                response.ping = "pong";
                break;
            case "message":
                response.message = message;
                break;
            case "randnum":
                response.randnum = 12.123;
                break;
            case "json":
                response.json = json;
                break;
        }
    }
    var sendTopic ="command/response/"+ uuid;
    publish( sendTopic, JSON.stringify(response));
    // publish('incoming/data/devices-topic/values',JSON.stringify(response))
});