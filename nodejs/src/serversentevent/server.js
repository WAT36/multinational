import express from "express";

const app = express();
const PORT = 4000;

app.use((_, res, next) => {
  res.setHeader("Access-Control-Allow-Origin", "*"); // CORS 許可
  next();
});

app.get("/events", (req, res) => {
  res.setHeader("Content-Type", "text/event-stream");
  res.setHeader("Cache-Control", "no-cache");
  res.setHeader("Connection", "keep-alive");

  const send = () => {
    const now = new Date().toLocaleTimeString();
    res.write(`data: 現在時刻は ${now} です\n\n`);
  };

  // 最初のメッセージ
  send();

  // 3秒ごとにメッセージ送信
  const interval = setInterval(send, 3000);

  // クライアントが切断したら終了
  req.on("close", () => {
    clearInterval(interval);
  });
});

app.listen(PORT, () => {
  console.log(`✅ SSE サーバーが http://localhost:${PORT} で起動中`);
});
