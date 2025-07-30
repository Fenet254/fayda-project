// index.js
const express = require("express");
const cookieSession = require("cookie-session");
const dotenv = require("dotenv");
const { Issuer, generators } = require("openid-client"); // Use correct imports

const dbRoutes = require("./backend/routes"); // Your API routes

dotenv.config();

const app = express();

app.use(express.json());

app.use(
  cookieSession({
    name: "session",
    keys: ["some very secret key"], // Replace with secure key in production
    maxAge: 24 * 60 * 60 * 1000, // 24 hours
  })
);

let client;
let clientReady = false;

(async () => {
  try {
    const issuer = await Issuer.discover("https://esignet.ida.fayda.et");
    client = new issuer.Client({
      client_id: process.env.CLIENT_ID,
      redirect_uris: [process.env.REDIRECT_URI],
      response_types: ["code"],
    });
    clientReady = true;
    console.log("✅ OIDC Client is ready");
  } catch (err) {
    console.error("❌ Failed to initialize OIDC client:", err);
  }
})();

// MySQL route setup
app.use("/api", dbRoutes); // Visit /api/test-db to test database

// Home route
app.get("/", (req, res) => {
  if (req.session.tokenSet) {
    res.send(`
      <h1>Welcome!</h1>
      <p>You are logged in.</p>
      <a href="/profile">View Profile</a><br>
      <a href="/logout">Logout</a>
    `);
  } else {
    res.send('<h1>Home</h1><a href="/login">Login with Fayda ID</a>');
  }
});

// Login route
app.get("/login", (req, res) => {
  if (!clientReady) return res.status(503).send("OIDC client not ready");

  const codeVerifier = generators.codeVerifier();
  const codeChallenge = generators.codeChallenge(codeVerifier);

  req.session.codeVerifier = codeVerifier;

  const authUrl = client.authorizationUrl({
    scope: "openid profile email",
    code_challenge: codeChallenge,
    code_challenge_method: "S256",
  });

  res.redirect(authUrl);
});

// Callback route
app.get("/callback", async (req, res) => {
  if (!clientReady) return res.status(503).send("OIDC client not ready");

  const params = client.callbackParams(req);

  try {
    const tokenSet = await client.callback(process.env.REDIRECT_URI, params, {
      code_verifier: req.session.codeVerifier,
    });

    req.session.tokenSet = tokenSet;
    req.session.userinfo = await client.userinfo(tokenSet.access_token);

    res.redirect("/");
  } catch (err) {
    console.error("❌ Callback error:", err);
    res.status(500).send("Authentication failed.");
  }
});

// Profile route
app.get("/profile", (req, res) => {
  if (!req.session.tokenSet) return res.redirect("/");

  res.send(`
    <h1>Profile</h1>
    <pre>${JSON.stringify(req.session.userinfo, null, 2)}</pre>
    <a href="/">Home</a><br>
    <a href="/logout">Logout</a>
  `);
});

// Logout route
app.get("/logout", (req, res) => {
  req.session = null;
  res.redirect("/");
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log(`🚀 Server started on http://localhost:${PORT}`);
});
