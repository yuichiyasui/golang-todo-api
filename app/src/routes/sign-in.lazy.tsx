import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Link } from "@/components/ui/link";
import { cn } from "@/utils";
import { zodResolver } from "@hookform/resolvers/zod";
import { createLazyFileRoute, useRouter } from "@tanstack/react-router";
import { useForm } from "react-hook-form";
import { z } from "zod";

const formSchema = z.object({
  email: z.string().email("メールアドレスの形式で入力してください"),
  password: z.string().min(1, "パスワードを入力してください"),
});

type FormValues = z.infer<typeof formSchema>;

const Page = () => {
  const form = useForm<FormValues>({
    mode: "onBlur",
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const router = useRouter();

  const signIn = form.handleSubmit(async () => {
    router.navigate({ to: "/" });
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
        <h1 className={cn("mb-3", "text-lg", "font-bold")}>ログイン</h1>
        <Form {...form}>
          <form onSubmit={signIn} className={cn("flex", "flex-col", "gap-y-4")}>
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>メールアドレス</FormLabel>
                  <FormControl>
                    <Input {...field} />
                  </FormControl>
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
                    <Input type="password" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <div>
              <Button type="submit">ログイン</Button>
            </div>
            <p
              aria-live="assertive"
              className={cn("text-red-500", "text-sm", "empty:hidden", "mt-2")}
            >
              {form.formState.errors.root?.message}
            </p>
          </form>
        </Form>
        <hr className={cn("my-4")} />
        <p className={cn("text-sm")}>
          初めての方は
          <Link to="/sign-up/email">ユーザー登録</Link>へ
        </p>
      </div>
    </main>
  );
};

export const Route = createLazyFileRoute("/sign-in")({
  component: Page,
});
