import { Link } from "@tanstack/react-router";

const Header = () => {
  return (
    <header className="w-full absolute top-0 left-0 py-1">
      <nav className="w-full max-w-[80%] px-8 mx-auto flex justify-between items-center">
        <Link
          to="/"
          className="cursor-pointer hover:scale-105 transition-all duration-100 flex items-center"
        >
          <img
            src="images/core/logo.png"
            alt="image-logo"
            className="h-32 w-auto block"
          />
        </Link>

        <Link
          to="/login"
          className="px-8 py-2 shadow-lg shadow-accent-foreground/50 rounded-md text-xl font-semibold cursor-pointer bg-primary text-white hover:bg-primary/90 transition-all duration-100 flex items-center justify-center"
        >
          Login
        </Link>
      </nav>
    </header>
  )
}

export default Header

