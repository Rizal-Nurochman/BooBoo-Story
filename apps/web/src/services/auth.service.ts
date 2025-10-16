import { apiRequest } from "@/lib/api";
import { LoginSchemaType, registerSchemaType } from "@/schemas/auth.schema";
import type { ApiResponse } from "@/types/api";

export const AuthService = {
  async login(data: LoginSchemaType): Promise<ApiResponse<{ token: string }>> {
    return apiRequest<{ token: string }>("/auth/login", {
      method: "POST",
      data,
    });
  },

  async me(): Promise<ApiResponse<{ id: string; email: string }>> {
    return apiRequest("/auth/me");
  },
  async register(data: registerSchemaType): Promise<ApiResponse<{ name: string; email: string }>> {
    return apiRequest<{ name: string; email: string }>("/auth/register", {
      method: "POST",
      data,
    });
  }
};
