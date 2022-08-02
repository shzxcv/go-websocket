const socket = new WebSocket("ws://localhost:8000/ws");

socket.addEventListener("open", function() {
  add_message("フロントからのこんにちは");
});

socket.addEventListener("message", function(event) {
  add_message(event.data);
})

const submit = document.getElementById("submit");
submit.addEventListener("click", function() {
  const input = document.getElementById("input");
  const send_msg = input.value;
  add_message(send_msg);
  socket.send(send_msg);
})

function add_message(message) {
  const message_list = document.getElementById("messages");
  const new_element = document.createElement("p");
  new_element.textContent = message;
  message_list.appendChild(new_element);
}
