// /auth/faydaAuth.js
const { Issuer, generators } = require("openid-client");

let client;

async function setupClient() {
  const faydaIssuer = await Issuer.discover("https://verifayda-issuer-url.com"); // 🔁 Replace this with the actual URL from VeriFayda docs

  client = new faydaIssuer.Client({
    client_id: "YOUR_CLIENT_ID", // 🔁 Replace with real Client ID
    client_secret: "YOUR_CLIENT_SECRET", // 🔁 Replace with real secret
    redirect_uris: ["http://localhost:3000/callback"],
    response_types: ["code"],
  });

  return client;
}

module.exports = {
  setupClient,
  getClient: () => client,
};
