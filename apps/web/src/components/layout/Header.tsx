import { buttonTap, fadeIn, slideDown } from "@/lib/animations";
import { Link, useLocation } from "@tanstack/react-router";
import { motion } from "motion/react";

const Header = () => {

  const { pathname }=useLocation()

  if(pathname.includes('/app/')) return null

  return (
    <motion.header 
      variants={slideDown}
      initial="hidden"
      animate="visible"
      className="w-full absolute top-0 left-0 py-0.5 md:py-1.5 z-50"
    >
      <nav className="w-full max-w-[90%] md:max-w-[90%] xl:max-w-[80%] px-4 md:px-6 lg:px-8 mx-auto flex justify-between items-center">
        <motion.div
          variants={fadeIn}
          initial="hidden"
          animate="visible"
          transition={{ delay: 0.2 }}
        >
          <Link
            to="/"
            className="cursor-pointer hover:scale-105 transition-all duration-100 flex items-center"
          >
            <img
              src="/images/core/logo.png"
              alt="image-logo"
              className="h-16 md:h-20 lg:h-24 xl:h-28 w-auto block"
            />
          </Link>
        </motion.div>

        <motion.div
          variants={fadeIn}
          initial="hidden"
          animate="visible"
          transition={{ delay: 0.3 }}
        >
          <motion.button whileTap={buttonTap}>
            <Link
              to="/auth/login"
              className="px-4 py-1.5 md:px-6 md:py-2 lg:px-8 lg:py-1.5 shadow-lg shadow-accent-foreground/50 rounded-md text-xs md:text-sm lg:text-md xl:text-lg font-semibold cursor-pointer bg-primary text-white hover:bg-primary/90 transition-all duration-100 flex items-center justify-center"
            >
              Login
            </Link>
          </motion.button>
        </motion.div>
      </nav>
    </motion.header>
  )
}

export default Header