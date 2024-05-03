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
send_msg_btn.onclick = () => {
    console.log("message had been sent: ", document.getElementById("user-message").value)
    clientWS.send(document.getElementById("user-message").value)
    input_field.value = " " 
}

clientWS.onopen = () => {
    clientWS.send("yay somebody joined :)")
}