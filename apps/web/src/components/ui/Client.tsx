import { useEffect, useState, type ReactNode } from "react";

interface ClientProps {
  children: ReactNode;
  fallback?: ReactNode;
}

export const Client = ({ children, fallback = null }: ClientProps) => {
  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

  if (!mounted) return <>{fallback}</>;
  return <>{children}</>;
};
