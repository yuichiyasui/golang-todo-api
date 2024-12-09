import { client } from "@/api/api";
import { cn } from "@/utils";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { createLazyFileRoute } from "@tanstack/react-router";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Link } from "@/components/ui/link";

const formSchema = z.object({
  email: z.string().email("メールアドレスの形式で入力してください"),
});
type FormValues = z.infer<typeof formSchema>;

const Page = () => {
  const form = useForm<FormValues>({
    mode: "onBlur",
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
    },
  });

  const submit = form.handleSubmit(async (data) => {
    try {
      await client.POST("/users/sign-up/email", {
        body: {
          email: data.email,
        },
      });
    } catch (error) {
      form.setError("root", { message: "通信エラーが発生しました" });
    }
  });

  return (
    <main className="p-4">
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
        <h1 className={cn("mb-3", "text-lg", "font-bold")}>
          ユーザー登録のメールの送信
        </h1>
        {(() => {
          if (!form.formState.isSubmitSuccessful) {
            return (
              <>
                <Form {...form}>
                  <form
                    onSubmit={submit}
                    className={cn("flex", "flex-col", "gap-y-4")}
                  >
                    <FormDescription>
                      ユーザー登録を行うメールアドレスを入力してください。
                    </FormDescription>
                    <FormField
                      control={form.control}
                      name="email"
                      render={({ field }) => {
                        return (
                          <FormItem>
                            <FormLabel>メールアドレス</FormLabel>
                            <FormControl>
                              <Input {...field} />
                            </FormControl>
                            <FormMessage />
                          </FormItem>
                        );
                      }}
                    />
                    <div>
                      <Button
                        type="submit"
                        disabled={form.formState.isSubmitting}
                      >
                        メールを送信
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
                <hr className={cn("my-4")} />
                <p className={cn("text-sm")}>
                  登録済みの方は
                  <Link to="/sign-in">ログイン</Link>へ
                </p>
              </>
            );
          }

          return (
            <div>
              <p>
                メールを送信しました。メール内のリンクをクリックしてユーザー登録を完了してください。
              </p>
            </div>
          );
        })()}
      </div>
    </main>
  );
};

export const Route = createLazyFileRoute("/sign-up/email")({
  component: Page,
});
