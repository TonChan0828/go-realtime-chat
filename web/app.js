ws.send(
  JSON.stringify({
    username: "sho",
    content: input.value,
    timestamp: new Date(),
  }),
);

ws.onmessage = (event) => {
  const msg = JSON.parse(event.data);

  switch (msg.type) {
    case "join":
      log.textContent = `🟢 ${msg.username} joined`;
      break;
    case "leave":
      log.textContent = `🔴 ${msg.username} left`;
      break;
    case "message":
      log.textContent += `${msg.username}: ${msg.content}\n`;
      break;
    case "system":
      log.textContent = `[system] ${msg.content}\n`;
      break;
  }
};
