import { createLazyFileRoute, Link } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import { components } from "@/api/__generated__/schema";
import { client } from "@/api/api";
import { Button } from "@/components/ui/button";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { cn } from "@/utils";

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
    <main className={cn("py-4", "px-6")}>
      <h1 className={cn("text-lg", "font-bold", "mb-3")}>タスク一覧</h1>
      <div className={cn("border", "rounded")}>
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className={cn("w-20")}>ID</TableHead>
              <TableHead>タスク名</TableHead>
              <TableHead>状態</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {(() => {
              if (tasks.length === 0) {
                return (
                  <TableRow>
                    <TableCell colSpan={3} className={cn("text-center")}>
                      タスクがありません
                    </TableCell>
                  </TableRow>
                );
              }

              return tasks.map((t) => (
                <TableRow key={t.id}>
                  <TableCell>{t.id}</TableCell>
                  <TableCell>
                    <p>
                      <Link to={`/$taskId`} params={{ taskId: t.id }}>
                        {t.title}
                      </Link>
                    </p>
                    {t.description && (
                      <p className={cn("text-xs", "mt-1", "text-slate-400")}>
                        {t.description}
                      </p>
                    )}
                  </TableCell>
                  <TableCell>{t.status}</TableCell>
                </TableRow>
              ));
            })()}
          </TableBody>
        </Table>
      </div>
      <Button asChild>
        <Link to="/new" className="mt-4">
          タスクを追加
        </Link>
      </Button>
    </main>
  );
};

export const Route = createLazyFileRoute("/")({
  component: Page,
});
