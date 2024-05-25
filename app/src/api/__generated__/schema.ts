/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export interface paths {
  "/tasks": {
    /** タスク一覧を取得する */
    get: operations["listTasks"];
    /** タスクを作成する */
    post: operations["createTask"];
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    Task: {
      id: string;
      title: string;
      description: string;
      status: components["schemas"]["TaskStatus"];
    };
    /** @enum {string} */
    TaskStatus: "todo" | "inProgress" | "done";
    Error: {
      code: string;
      message: string;
    };
  };
  responses: never;
  parameters: never;
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export interface operations {

  /** タスク一覧を取得する */
  listTasks: {
    responses: {
      /** @description タスク一覧を返す */
      200: {
        content: {
          "application/json": components["schemas"]["Task"][];
        };
      };
      /** @description エラー */
      default: {
        content: {
          "application/json": components["schemas"]["Error"];
        };
      };
    };
  };
  /** タスクを作成する */
  createTask: {
    requestBody?: {
      content: {
        "application/json": {
          /** @description タスク名 */
          title: string;
          /** @description タスクの説明 */
          description?: string;
        };
      };
    };
    responses: {
      /** @description 作成したタスクのID */
      200: {
        content: {
          "application/json": {
            id: string;
          };
        };
      };
      /** @description エラー */
      default: {
        content: {
          "application/json": components["schemas"]["Error"];
        };
      };
    };
  };
}