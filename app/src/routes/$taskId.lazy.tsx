import { client } from "@/api/api";
import { createLazyFileRoute, useRouter, Link } from "@tanstack/react-router";
import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { cn } from "@/utils";
import { useForm } from "react-hook-form";
import { components } from "@/api/__generated__/schema";
import { useEffect } from "react";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";

const formSchema = z.object({
  title: z
    .string()
    .min(1, { message: "タスク名を入力してください" })
    .max(30, "タスク名を30文字以内で入力してください"),
  description: z
    .string()
    .max(500, { message: "タスク内容を500文字以内で入力してください" }),
  status: z.string(),
});

type FormValues = z.infer<typeof formSchema>;

type Status = components["schemas"]["TaskStatus"];

type StatusValue = {
  [K in Status]: K;
};

const statusValue = {
  todo: "todo",
  inProgress: "inProgress",
  done: "done",
} as const satisfies StatusValue;

const convertStatus = (status: string) => {
  switch (status) {
    case statusValue.todo:
    case statusValue.inProgress:
    case statusValue.done:
      return status;
    default:
      return statusValue.todo;
  }
};

const Page = () => {
  const form = useForm<FormValues>({
    mode: "onBlur",
    resolver: zodResolver(formSchema),
    defaultValues: {
      title: "",
      description: "",
      status: statusValue.todo,
    },
  });
  const router = useRouter();
  const { taskId } = Route.useParams();

  const updateTask = form.handleSubmit(async (data) => {
    try {
      await client.PUT("/tasks/{taskId}", {
        params: {
          path: { taskId },
        },
        body: {
          title: data.title,
          description: data.description,
          status: convertStatus(data.status),
        },
      });
      router.navigate({ to: "/" });
    } catch (error) {
      form.setError("root", { message: "通信エラーが発生しました" });
    }
  });

  useEffect(() => {
    client
      .GET("/tasks/{taskId}", {
        params: { path: { taskId } },
      })
      .then(({ data }) => {
        if (typeof data === "undefined") return;

        form.setValue("title", data.title);
        form.setValue("description", data.description);
        form.setValue("status", convertStatus(data.status));
      });
  }, [form, taskId]);

  const {
    formState: { errors, isSubmitting },
  } = form;

  return (
    <main className="p-4">
      <Breadcrumb className="mb-4">
        <BreadcrumbList>
          <BreadcrumbItem>
            <BreadcrumbLink asChild>
              <Link to="/">TOP</Link>
            </BreadcrumbLink>
          </BreadcrumbItem>
          <BreadcrumbSeparator />
          <BreadcrumbItem>
            <BreadcrumbPage>タスクの編集</BreadcrumbPage>
          </BreadcrumbItem>
        </BreadcrumbList>
      </Breadcrumb>

      <h1 className={cn("mb-3", "text-lg", "font-bold")}>タスクの編集</h1>
      <Form {...form}>
        <form
          onSubmit={updateTask}
          className={cn("flex", "flex-col", "gap-y-4")}
        >
          <FormField
            control={form.control}
            name="title"
            render={({ field }) => (
              <FormItem>
                <FormLabel>タスク名</FormLabel>
                <FormControl>
                  <Input {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="description"
            render={({ field }) => (
              <FormItem>
                <FormLabel>タスク内容</FormLabel>
                <FormControl>
                  <Textarea {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <FormField
            control={form.control}
            name="status"
            render={({ field }) => (
              <FormItem>
                <FormLabel>ステータス</FormLabel>
                <Select
                  onValueChange={field.onChange}
                  defaultValue={field.value}
                >
                  <FormControl>
                    <SelectTrigger>
                      <SelectValue />
                    </SelectTrigger>
                  </FormControl>
                  <SelectContent>
                    <SelectItem value={statusValue.todo}>Todo</SelectItem>
                    <SelectItem value={statusValue.inProgress}>
                      In Progress
                    </SelectItem>
                    <SelectItem value={statusValue.done}>Done</SelectItem>
                  </SelectContent>
                </Select>
                <FormMessage />
              </FormItem>
            )}
          />
          <div>
            <Button type="submit" disabled={isSubmitting}>
              保存
            </Button>
          </div>
          <p
            aria-live="assertive"
            className={cn("text-red-500", "text-sm", "empty:mt-0", "mt-2")}
          >
            {errors.root?.message}
          </p>
        </form>
      </Form>
    </main>
  );
};

export const Route = createLazyFileRoute("/$taskId")({
  component: Page,
});
