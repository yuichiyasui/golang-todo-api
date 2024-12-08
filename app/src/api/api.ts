import createClient, { type Middleware } from "openapi-fetch";
import type { paths } from "./__generated__/schema";

const API_URL: string = import.meta.env.VITE_API_URL;

export class FetchError extends Error {
  response: Response;

  constructor(message: string, response: Response) {
    super(message);
    this.response = response;
  }
}

const raiseErrorMiddleware = {
  onResponse: async (res) => {
    if (!res.ok) {
      throw new FetchError("fetch error", res);
    }

    return res;
  },
} satisfies Middleware;

export const client = createClient<paths>({ baseUrl: API_URL });
client.use(raiseErrorMiddleware);
