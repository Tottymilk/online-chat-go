clientWS = WebSocket("ws://localhost:8085/")

var send_msg_btn = document.getElementById("send-message")
send_msg_btn.onclick = function() {
    clientWS.send(document.getElementById("user-message").value)
}