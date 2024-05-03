let clientWS = new WebSocket("ws://localhost:8085/ws")

var send_msg_btn = document.getElementById("send-message")
send_msg_btn.onclick = () => {
    clientWS.send(document.getElementById("user-message").value)
}