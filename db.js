const mysql = require("mysql2");

const connection = mysql.createConnection({
  host: "localhost",
  user: "root",
  password: "Ru0810/15", // replace with your MySQL password
  database: "fayda_queue_system",
});

connection.connect((err) => {
  if (err) {
    console.error("❌ MySQL connection error:", err.message);
  } else {
    console.log("✅ Connected to MySQL Database!");
  }
});

module.exports = connection;
