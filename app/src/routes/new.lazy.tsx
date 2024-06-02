import { client } from "@/api/api";
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
import { createLazyFileRoute, useRouter, Link } from "@tanstack/react-router";
import { useForm } from "react-hook-form";
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
});
type FormValues = z.infer<typeof formSchema>;

const Page = () => {
  const form = useForm<FormValues>({
    mode: "onBlur",
    resolver: zodResolver(formSchema),
    defaultValues: {
      title: "",
      description: "",
    },
  });
  const router = useRouter();

  const createTask = form.handleSubmit(async (data) => {
    try {
      await client.POST("/tasks", {
        body: {
          title: data.title,
          description: data.description,
        },
      });
      router.navigate({ to: "/" });
    } catch (error) {
      form.setError("root", { message: "通信エラーが発生しました" });
    }
  });

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
            <BreadcrumbPage>タスクの作成</BreadcrumbPage>
          </BreadcrumbItem>
        </BreadcrumbList>
      </Breadcrumb>

      <h1 className={cn("mb-3", "text-lg", "font-bold")}>タスクの作成</h1>
      <Form {...form}>
        <form
          onSubmit={createTask}
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
          <div>
            <Button type="submit" disabled={isSubmitting}>
              作成
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

export const Route = createLazyFileRoute("/new")({
  component: Page,
});
