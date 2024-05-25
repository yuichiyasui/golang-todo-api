import createClient from "openapi-fetch";
import type { paths } from "./__generated__/schema";

const API_URL: string = import.meta.env.VITE_API_URL;

export const client = createClient<paths>({ baseUrl: API_URL });
