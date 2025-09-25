import { Link } from "@tanstack/react-router"

const HeroSection = () => {
  return (
    <div className="w-full flex items-center flex-col gap-8 pt-32 pb-8 justify-center">
        <h1 className="text-5xl font-extrabold text-center">Belajar, bermain, dan berkembang bersama BooBoo Story!</h1>
        <img src='/images/home/landing-page.svg' />
        <p className="text-5xl font-bold text-center">
            Siapkah kamu belajar bersama BooBoo?
        </p>
        <div className="flex gap-4 w-full max-w-md mx-auto">
            <Link 
              to="/register" 
              className="flex-1 py-3 rounded-lg shadow-lg shadow-accent-foreground/50 bg-primary text-white text-lg font-semibold text-center tracking-wide hover:scale-[1.02] active:scale-[0.98] transition-all duration-150"
            >
                Ikut Petualangan!
            </Link>
            <Link 
              to="/login" 
              className="flex-1 py-3 rounded-lg border-2 border-primary text-center text-lg font-semibold tracking-wide hover:bg-primary/5 hover:scale-[1.02] active:scale-[0.98] transition-all duration-150"
            >
                Ayo Mulai
            </Link>
        </div>
    </div>
  )
}

export default HeroSection
