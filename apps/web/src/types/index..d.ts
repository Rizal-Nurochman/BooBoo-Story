interface User {
  id: string;
  name: string;
  email: string;
  avatar: string;
  role: 'Reader' | 'Creator' | 'Admin'; 
  level: number;
  points: number;
  streak: number;
}
