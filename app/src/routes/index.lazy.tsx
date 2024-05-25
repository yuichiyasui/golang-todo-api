import { createLazyFileRoute, Link } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import { components } from "../api/__generated__/schema";
import { client } from "../api/api";

const Page = () => {
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
    <main className="py-4 px-6">
      <h1 className="text-xl font-bold mb-4">Task App</h1>
      <ul className="mb-4">
        {tasks.map((t) => (
          <li key={t.id}>
            <p>{t.id}</p>
            <p>{t.title}</p>
            <p>{t.description}</p>
            <p>{t.status}</p>
          </li>
        ))}
      </ul>
      <div>
        <Link
          to="/new"
          className="bg-blue-500 px-4 py-1 rounded text-white font-bold"
        >
          タスクを追加
        </Link>
      </div>
    </main>
  );
};

export const Route = createLazyFileRoute("/")({
  component: Page,
});
