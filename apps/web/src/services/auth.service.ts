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

  async me(cookie?:string): Promise<ApiResponse<User>> {
    return apiRequest<User>("/auth/me", {
      method:'GET',
      cookie
    });
  },
  async register(data: registerSchemaType): Promise<ApiResponse<User>> {
    return apiRequest<User>("/auth/register", {
      method: "POST",
      data,
    });
  },
  async logout(): Promise<ApiResponse<null>> {
    return apiRequest<null>("/auth/logout", {
      method: "DELETE",
    });
  },
};
