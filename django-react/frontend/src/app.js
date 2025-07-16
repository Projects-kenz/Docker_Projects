import React, { useEffect, useState } from "react";

function App() {
  const [msg, setMsg] = useState("");

  useEffect(() => {
    fetch("http://localhost:8000/api/ping/")


      .then((res) => res.json())
      .then((data) => setMsg(data.message));
  }, []);

  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <h1>React + Django + Postgres</h1>
      <p>Backend says: <strong>{msg}</strong></p>
    </div>
  );
}

export default App;
