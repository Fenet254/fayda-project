function bookToken() {
  const office = document.getElementById("office-select").value;
  if (!office) {
    alert("Please select an office.");
    return;
  }

  const token = Math.floor(Math.random() * 1000) + 1;
  const message = `Office: ${office}<br>Token Number: ${token}<br>Status: Waiting`;

  document.getElementById("token-info").innerHTML = message;
}
// Mocked token list (replace with real DB later)
const queues = [
  { id: 1, name: "Fenet", office: "Passport", token: "A102", status: "Waiting" },
  { id: 2, name: "Liya", office: "Tax", token: "B017", status: "Waiting" },
];

function renderQueue() {
  const container = document.getElementById("queue-container");
  container.innerHTML = "";

  queues.forEach((q, index) => {
    const div = document.createElement("div");
    div.className = "queue-item";
    div.innerHTML = `
      <strong>${q.name}</strong> - ${q.office} | Token: ${q.token} | Status: ${q.status}
      <button onclick="markServed(${index})">Mark as Served</button>
    `;
    container.appendChild(div);
  });
}

function markServed(index) {
  queues[index].status = "Served";
  renderQueue();
}

renderQueue();
