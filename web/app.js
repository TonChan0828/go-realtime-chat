ws.send(
  JSON.stringify({
    username: "sho",
    content: input.value,
    timestamp: new Date(),
  }),
);
