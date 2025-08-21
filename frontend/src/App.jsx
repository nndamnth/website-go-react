import { useEffect, useState } from "react";

function App() {
  const [message, setMessage] = useState("");

  useEffect(() => {
    fetch("/api/hello")   // ⚡ cukup '/api/hello', jangan full URL
      .then(res => res.json())
      .then(data => setMessage(data.message))
      .catch(err => console.error("Error fetching backend:", err));
  }, []);

  return (
    <div>
      <h1>Training Needs Analysis</h1>
      <p>Message from backend: {message}</p>
    </div>
  );
}

export default App;
