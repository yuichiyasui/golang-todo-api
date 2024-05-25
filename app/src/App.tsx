import { useEffect, useState } from "react";
import { client } from "./api/api";
import { components } from "./api/__generated__/schema";

export const App = () => {
  const [tasks, setTasks] = useState<components["schemas"]["Task"][]>([]);

  useEffect(() => {
    client.GET("/tasks").then(({ data }) => {
      if (typeof data === "undefined") {
        return;
      }

      setTasks(data);
    });
  }, []);

  return (
    <main>
      <h1>Task App</h1>
      <ul>
        {tasks.map((t) => (
          <li key={t.id}>
            <p>{t.id}</p>
            <p>{t.title}</p>
            <p>{t.description}</p>
            <p>{t.status}</p>
          </li>
        ))}
      </ul>
      {}
    </main>
  );
};
