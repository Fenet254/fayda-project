const express = require("express");
const db = require("../db");

const router = express.Router();

router.get("/test-db", (req, res) => {
  db.query("SELECT NOW()", (err, results) => {
    if (err) {
      console.error("DB error:", err);
      return res.status(500).send("Database error");
    }
    res.send(results);
  });
});

module.exports = router;
