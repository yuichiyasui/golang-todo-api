import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { z } from "zod";
import { Input } from "@/components/ui/input";
import { Link } from "@/components/ui/link";
import { cn } from "@/utils";
import { zodResolver } from "@hookform/resolvers/zod";
import { createFileRoute } from "@tanstack/react-router";
import { useForm } from "react-hook-form";
import { client } from "@/api/api";

const formSchema = z
  .object({
    username: z.string().min(1, "ユーザー名を入力してください"),
    password: z.string().min(8, "パスワードを8文字以上で入力してください"),
    confirmationPassword: z
      .string()
      .min(1, "確認用パスワードを入力してください"),
  })
  .refine((data) => data.password === data.confirmationPassword, {
    message: "パスワードが一致しません",
  });

type FormValues = z.infer<typeof formSchema>;

const Page = () => {
  const form = useForm<FormValues>({
    mode: "onBlur",
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
      password: "",
      confirmationPassword: "",
    },
  });

  const searchParams = Route.useSearch();

  const signUp = form.handleSubmit(async (data) => {
    try {
      await client.POST("/users/sign-up", {
        body: {
          username: data.username,
          password: data.password,
          token: searchParams.token,
        },
      });
    } catch (error) {
      form.setError("root", {
        message: "ユーザー登録に失敗しました",
      });
    }
  });

  return (
    <main>
      <div
        className={cn(
          "w-96",
          "mx-auto",
          "my-20",
          "border",
          "border-slate-200",
          "p-8",
        )}
      >
        <h1 className={cn("mb-3", "text-lg", "font-bold")}>ユーザー登録</h1>
        {(() => {
          if (!form.formState.isSubmitSuccessful) {
            return (
              <Form {...form}>
                <form
                  onSubmit={signUp}
                  className={cn("flex", "flex-col", "gap-y-4")}
                >
                  <FormField
                    control={form.control}
                    name="username"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>ユーザー名</FormLabel>
                        <FormControl>
                          <Input {...field} aria-required />
                        </FormControl>
                        <FormDescription>
                          他のユーザーに表示される名前です。
                        </FormDescription>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name="password"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>パスワード</FormLabel>
                        <FormControl>
                          <Input type="password" aria-required {...field} />
                        </FormControl>
                        <FormDescription>
                          パスワードを8文字以上で入力してください
                        </FormDescription>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name="confirmationPassword"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>パスワードの確認</FormLabel>
                        <FormControl>
                          <Input type="password" aria-required {...field} />
                        </FormControl>
                        <FormDescription>
                          パスワードをもう一度入力してください
                        </FormDescription>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <div>
                    <Button
                      type="submit"
                      disabled={form.formState.isSubmitting}
                    >
                      登録
                    </Button>
                  </div>
                  <p
                    aria-live="assertive"
                    className={cn(
                      "text-red-500",
                      "text-sm",
                      "empty:hidden",
                      "mt-2",
                    )}
                  >
                    {form.formState.errors.root?.message}
                  </p>
                </form>
              </Form>
            );
          }

          return (
            <div>
              <p>ユーザー登録が完了しました。ログインしてください。</p>
              <div>
                <Link to="/sign-in">ログイン</Link>
              </div>
            </div>
          );
        })()}
      </div>
    </main>
  );
};

const searchParamsSchema = z.object({
  token: z.string().catch(""),
});

export const Route = createFileRoute("/sign-up/")({
  component: Page,
  validateSearch: (search: Record<string, unknown>) => {
    return searchParamsSchema.parse(search);
  },
});
