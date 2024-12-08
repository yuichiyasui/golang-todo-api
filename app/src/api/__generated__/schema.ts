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
  "/tasks/{taskId}": {
    /** タスク詳細を取得する */
    get: operations["getTaskDetail"];
    /** タスクを更新する */
    put: operations["updateTask"];
  };
  "/users/sign-up/email": {
    /** 会員登録用のメールを送信する */
    post: operations["sendSignUpEmail"];
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
    /** @description 会員登録用のメール送信リクエスト */
    SignUpEmailRequest: {
      /**
       * Format: email
       * @description 送信先のメールアドレス
       */
      email: string;
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
  /** タスク詳細を取得する */
  getTaskDetail: {
    parameters: {
      path: {
        /** @description タスクID */
        taskId: string;
      };
    };
    responses: {
      /** @description タスク詳細を返す */
      200: {
        content: {
          "application/json": components["schemas"]["Task"];
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
  /** タスクを更新する */
  updateTask: {
    parameters: {
      path: {
        /** @description タスクID */
        taskId: string;
      };
    };
    requestBody?: {
      content: {
        "application/json": {
          /** @description タスク名 */
          title: string;
          /** @description タスクの説明 */
          description?: string;
          status: components["schemas"]["TaskStatus"];
        };
      };
    };
    responses: {
      /** @description 更新したタスクを返す */
      200: {
        content: {
          "application/json": components["schemas"]["Task"];
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
  /** 会員登録用のメールを送信する */
  sendSignUpEmail: {
    requestBody?: {
      content: {
        "application/json": components["schemas"]["SignUpEmailRequest"];
      };
    };
    responses: {
      /** @description メールの送信に成功 */
      200: {
        content: never;
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
