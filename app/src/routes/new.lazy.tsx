import { createLazyFileRoute } from "@tanstack/react-router";

const Page = () => {
  return (
    <main>
      <h1>タスクの作成</h1>
      <form>
        <div>
          <label htmlFor="">タスク名</label>
          <input type="text" name="" id="" />
        </div>
        <div>
          <label htmlFor="">タスク内容</label>
          <textarea name="" id=""></textarea>
        </div>
        <div>
          <button type="submit">作成</button>
        </div>
      </form>
    </main>
  );
};

export const Route = createLazyFileRoute("/new")({
  component: Page,
});
