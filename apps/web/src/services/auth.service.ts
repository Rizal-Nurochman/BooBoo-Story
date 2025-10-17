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

  async me(cookie?:string): Promise<ApiResponse<{ id: string; email: string }>> {
    return apiRequest("/auth/me", {
      method:'GET',
      cookie
    });
  },
  async register(data: registerSchemaType): Promise<ApiResponse<{ name: string; email: string }>> {
    return apiRequest<{ name: string; email: string }>("/auth/register", {
      method: "POST",
      data,
    });
  }
};
