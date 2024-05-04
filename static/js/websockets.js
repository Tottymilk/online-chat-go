let clientWS = new WebSocket("ws://localhost:8085/ws")

window.addEventListener('beforeunload', () => {
    clientWS.send("somebody just left :(")
})

var chat_box = document.getElementById("chat-box")

clientWS.onmessage = (event) => {
    console.log("message recieved: ", event.data)
    chat_box.textContent = chat_box.textContent + event.data
}

var input_field = document.getElementById("user-message")
var send_msg_btn = document.getElementById("send-message")
var send_msg_form = document.getElementById("send-message-form")


clientWS.onopen = () => {
    clientWS.send("yay somebody joined :)")
}