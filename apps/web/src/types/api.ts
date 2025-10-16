type Meta ={
  page: number;
  limit: number;
  total: number;
  [key: string]: any;
}

export type ApiResponse<T = any> ={
  status: "success" | "error";
  message: string;
  data?: T | null;
  errors?: any | null;
  meta?: Meta | null;
}
