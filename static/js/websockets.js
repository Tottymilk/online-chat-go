let clientWS = new WebSocket("ws://localhost:8085/ws")


clientWS.onopen = () => {
    clientWS.send(" yay somebody joined :) ")
}

clientWS.onmessage = (event) => {
    console.log("message recieved: ", event.data)
    chat_box.textContent = chat_box.textContent + event.data
}

window.addEventListener('beforeunload', () => {
    clientWS.send(" somebody just left :( ")
    clientWS.close()
})

var chat_box = document.getElementById("chat-box")
var input_field = document.getElementById("user-message")
var send_msg_form = document.getElementById("send-message-form")

htmx.on(send_msg_form, "htmx:afterRequest", () => {
    console.log("message had been sent: ", input_field.value)
    // TODO: make date time appear from the database datestamp, this is unreliable
    clientWS.send(" " + getCurrentDateTime() + " | " + input_field.value)
    input_field.value = ""   
})

function getCurrentDateTime() {
    let now = new Date();
    let year = now.getFullYear();
    let month = (now.getMonth() + 1).toString().padStart(2, '0');
    let day = now.getDate().toString().padStart(2, '0');
    let hours = now.getHours().toString().padStart(2, '0');
    let minutes = now.getMinutes().toString().padStart(2, '0');
    let dateTime = day + "-" + month + "-" + year + " " + hours + ":" + minutes;
    return dateTime;
}
