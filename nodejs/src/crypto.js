const crypto = require("crypto");
const path = require("path");
const ENV_PATH = path.join(__dirname, "../.env");
require("dotenv").config({ path: ENV_PATH });

const timestamp = Date.now().toString();
const method = "GET";
const apiPath = "/v1/activeOrders";

console.log("ENV_TEST:" + process.env.ENV_TEST);
const secretKey = process.env.COIN_API_SECRETKEY;
const text = timestamp + method + apiPath;
const sign = crypto.createHmac("sha256", secretKey).update(text).digest("hex");

console.log("sign:");
console.log(sign);
