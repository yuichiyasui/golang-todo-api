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
import { Label } from "@/components/ui/label";
import { Textarea } from "@/components/ui/textarea";
import { cn } from "@/utils";
import { createLazyFileRoute, useRouter, Link } from "@tanstack/react-router";
import { useForm } from "react-hook-form";

type FormData = {
  title: string;
  description: string;
};

const titleFieldId = "title";
const titleFieldErrorId = "title-error";
const descriptionFieldId = "description";
const descriptionFieldErrorId = "description-error";

const Page = () => {
  const {
    register,
    handleSubmit,
    setError,
    formState: { errors, isSubmitting },
  } = useForm<FormData>({
    mode: "onBlur",
  });
  const router = useRouter();

  const createTask = handleSubmit(async (data) => {
    try {
      await client.POST("/tasks", {
        body: {
          title: data.title,
          description: data.description,
        },
      });
      router.navigate({ to: "/" });
    } catch (error) {
      setError("root", { message: "通信エラーが発生しました" });
    }
  });

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
      <form onSubmit={createTask} className={cn("flex", "flex-col", "gap-y-4")}>
        <div>
          <Label htmlFor={titleFieldId}>タスク名</Label>
          <Input
            {...register("title", {
              required: {
                value: true,
                message: "タスク名を入力してください",
              },
              maxLength: {
                value: 30,
                message: "タスク名を30文字以内で入力してください",
              },
            })}
            id={titleFieldId}
            aria-errormessage={titleFieldErrorId}
            aria-required
            aria-invalid={!!errors.title}
          />
          <p
            id={titleFieldErrorId}
            aria-live="polite"
            className={cn("text-red-500", "text-sm", "empty:mt-0", "mt-2")}
          >
            {errors.title?.message}
          </p>
        </div>
        <div>
          <Label htmlFor={descriptionFieldId}>タスク内容</Label>
          <Textarea
            {...register("description", {
              maxLength: {
                value: 500,
                message: "タスク内容を500文字以内で入力してください",
              },
            })}
            id={descriptionFieldId}
            aria-errormessage={descriptionFieldErrorId}
            aria-invalid={!!errors.description}
          />
          <p id={descriptionFieldErrorId} aria-live="polite">
            {errors.description?.message}
          </p>
        </div>
        <div>
          <Button type="submit" disabled={isSubmitting}>
            作成
          </Button>
        </div>
        <p aria-live="assertive">{errors.root?.message}</p>
      </form>
    </main>
  );
};

export const Route = createLazyFileRoute("/new")({
  component: Page,
});
