// WebSocket モジュールを import
import { WebSocketServer } from "ws";

// ポート 3000 で WebSocket サーバーを起動
const wss = new WebSocketServer({ port: 3000 });

// クライアントが接続した時の処理
wss.on("connection", (ws) => {
  console.log("クライアントが接続しました。");

  // クライアントからメッセージを受信
  ws.on("message", (message) => {
    console.log(message);

    // すべてのクライアントにメッセージをブロードキャスト
    wss.clients.forEach((client) => {
      if (client.readyState === ws.OPEN) {
        client.send(message);
      }
    });
  });

  // クライアントが切断した時の処理
  ws.on("close", () => {
    console.log("クライアントが切断しました。");
  });
});
