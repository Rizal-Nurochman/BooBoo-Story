import { BE_API_URL } from "@/lib/be";
import type { ApiResponse } from "@/types/api";

interface FetchOptions {
  method?: "GET" | "POST" | "PUT" | "PATCH" | "DELETE";
  data?: any;
  headers?: HeadersInit;
}

export async function apiRequest<T = any>(
  endpoint: string,
  options: FetchOptions = {}
): Promise<ApiResponse<T>> {
  const { method = "GET", data, headers } = options;

  try {
    const res = await fetch(`${BE_API_URL}${endpoint}`, {
      method,
      headers: {
        "Content-Type": "application/json",
        ...headers,
      },
      body: method !== "GET" ? JSON.stringify(data) : undefined,
    });

    const result = await res.json();

    if (!res.ok) {
      return {
        status: "error",
        message: result?.message || "Request failed",
        errors: result?.errors || null,
        data: null,
        meta: null,
      };
    }

    return result;
  } catch (error) {
    return {
      status: "error",
      message: error instanceof Error ? error.message : "Unexpected error occurred",
      data: null,
      errors: error,
      meta: null,
    };
  }
}


export const sleep = (ms: number = 1000) => new Promise((resolve) => setTimeout(resolve, ms));