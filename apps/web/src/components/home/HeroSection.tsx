import { Link } from "@tanstack/react-router";
import { motion } from "motion/react";
import { 
  staggerContainer,
  slideUp,
  fadeIn,
  scaleIn,
  buttonTap
} from "@/lib/animations";



const HeroSection = () => {

  return (
    <motion.div 
      className="w-full max-w-7xl mx-auto flex items-center flex-col gap-6 sm:gap-8 pt-14 sm:pt-16 md:pt-20 pb-6 sm:pb-8 md:pb-12 lg:pb-16 px-4 sm:px-6 lg:px-8 justify-center"
      variants={staggerContainer}
      initial="hidden"
      whileInView="visible"
      viewport={{ once: true, margin: "-100px" }}
    >
      <motion.h1 
        className="text-lg sm:text-xl md:text-2xl font-extrabold text-center leading-tight"
        variants={slideUp}
      >
        Belajar, bermain, dan berkembang bersama BooBoo Story!
      </motion.h1>
      
      <motion.img 
        src="/images/home/landing-page.svg" 
        className="w-full max-w-64 sm:max-w-72 md:max-w-80 lg:max-w-sm xl:max-w-md"
        variants={scaleIn}
        whileHover={{
          scale: 1.05,
          transition: { duration: 0.3, ease: [0.25, 0.1, 0.25, 1] },
        }}
      />
      
      <motion.p 
        className="text-lg sm:text-xl md:text-2xl font-bold text-center leading-tight"
        variants={slideUp}
      >
        Siapkah kamu belajar bersama BooBoo?
      </motion.p>
      
      <motion.div 
        className="flex flex-col sm:flex-row gap-3 sm:gap-4 w-full max-w-xs sm:max-w-md mx-auto"
        variants={fadeIn}
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true }}
        transition={{ delay: 0.3 }}
      >
        <motion.div 
          whileHover={{ 
            scale: 1.02, 
            y: -2,
            transition: { duration: 0.2, ease: [0.25, 0.1, 0.25, 1] }
          }} 
          whileTap={buttonTap}
          className="flex-1"
        >
          <Link 
            to="/auth/register" 
            className="block py-0.5 sm:py-1 lg:py-1.5 xl:py-2 rounded-lg shadow-lg shadow-accent-foreground/50 bg-primary text-white text-sm sm:text-base lg:text-md 2xl:text-md font-semibold text-center tracking-wide"
          >
            Ikut Petualangan!
          </Link>
        </motion.div>

        <motion.div 
          whileHover={{ 
            scale: 1.02, 
            y: -2,
            transition: { duration: 0.2, ease: [0.25, 0.1, 0.25, 1] }
          }} 
          whileTap={buttonTap}
          className="flex-1"
        >
          <Link 
            to="/auth/login" 
            className="block py-0.5 sm:py-1 lg:py-1.5 xl:py-2 rounded-lg border-2 border-primary text-sm sm:text-base lg:text-md 2xl:text-md font-semibold text-center tracking-wide hover:bg-primary/5"
          >
            Ayo Mulai
          </Link>
        </motion.div>
      </motion.div>
    </motion.div>
  );
};

export default HeroSection;