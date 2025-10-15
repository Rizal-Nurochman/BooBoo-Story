import { Link } from '@tanstack/react-router'
import { Avatar, AvatarFallback, AvatarImage } from '../ui/avatar'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Button } from '../ui/button'
import { ArrowLeftToLine } from 'lucide-react'

const user = {
  name: "Kristin Watson",
  email: "kristin@example.com",
  role: "Creator",
  level: 5,
  points: 320,
  avatar: "https://github.com/shadcn.png",
}

const handleLogout = () => {
  console.log("User logged out (dummy)")
}

const UserDropdown = () => {
  return (
    <div>
      <DropdownMenu>
        <DropdownMenuTrigger asChild>
          <Avatar className="cursor-pointer hover:scale-105 transition-all duration-100">
            <AvatarImage src={user.avatar} alt={user.name} />
            <AvatarFallback>{user.name?.charAt(0) || "U"}</AvatarFallback>
          </Avatar>
        </DropdownMenuTrigger>

        <DropdownMenuContent align="end" className="w-60 p-2">
          <DropdownMenuLabel className="pb-2">
            <div className="flex items-center gap-3">
              <Avatar className="w-10 h-10 ring-2 ring-muted">
                <AvatarImage src={user.avatar} alt={user.name} />
                <AvatarFallback className="text-lg">
                  {user.name?.charAt(0) || "U"}
                </AvatarFallback>
              </Avatar>
              <div className="flex flex-col">
                <span className="text-base font-semibold leading-tight">{user.name}</span>
                <span className="text-xs text-muted-foreground">{user.email}</span>
                <span className="text-xs text-muted-foreground mt-1">
                  Lv.{user.level} • {user.points} pts
                </span>
              </div>
            </div>
          </DropdownMenuLabel>

          <DropdownMenuSeparator />

          <DropdownMenuItem asChild>
            <Link to="/user/profile">Profile</Link>
          </DropdownMenuItem>

          <DropdownMenuItem asChild>
            <Link to="/user/bookmarks">Bookmarks</Link>
          </DropdownMenuItem>

          <DropdownMenuItem asChild>
            <Link to="/user/achievements">Achievements</Link>
          </DropdownMenuItem>

          {user.role === "Creator" && (
            <DropdownMenuItem asChild>
              <Link to="/user/my-stories">My Stories</Link>
            </DropdownMenuItem>
          )}

          <DropdownMenuSeparator />

          <Button variant={'destructive'} className='w-full flex gap-2 items-center cursor-pointer'>
            <ArrowLeftToLine size={20} />
            <span>Logout</span>
          </Button>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  )
}

export default UserDropdown
