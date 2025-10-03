"use client"

import { motion, Variants } from "motion/react"

const arrAuthor = [
  { id: 1, name: "Kristin Watson", image: "/images/home/pic-1.png", role: "Chef" },
  { id: 2, name: "Jenny Wilson", image: "/images/home/pic-2.png", role: "Artist" },
  { id: 3, name: "Jacob Jones", image: "/images/home/pic-3.png", role: "Teacher" },
  { id: 4, name: "Savannah Nguyen", image: "/images/home/pic-4.png", role: "Writer" },
]

// Animation variants with proper typing
const easeOutFast: [number, number, number, number] = [0.25, 0.1, 0.25, 1]
const bounceEase: [number, number, number, number] = [0.68, -0.55, 0.265, 1.55]

const fadeInUp: Variants = {
  hidden: { opacity: 0, y: 20 },
  visible: {
    opacity: 1,
    y: 0,
    transition: {
      duration: 0.6,
      ease: easeOutFast
    }
  }
}

const staggerContainer: Variants = {
  hidden: {},
  visible: {
    transition: {
      staggerChildren: 0.15,
    }
  }
}

const playfulCardEnter = (index: number): Variants => ({
  hidden: { 
    opacity: 0, 
    x: -50, 
    rotate: -2 
  },
  visible: { 
    opacity: 1, 
    x: 0, 
    rotate: 0,
    transition: {
      duration: 0.6,
      delay: index * 0.1,
      ease: bounceEase
    }
  }
})

const playfulHover = {
  scale: 1.05,
  y: -8,
  rotate: 0.5,
  transition: { 
    duration: 0.3, 
    ease: easeOutFast 
  }
}

const AuthorSection = () => {
  return (
    <div className="bg-[#E08F62] pt-20 sm:pt-24 md:pt-32 pb-6 sm:pb-8 md:pb-12 lg:pb-16 overflow-hidden">
      <div className="w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 relative">
        {/* Decorative images with floating animation */}
        <motion.img 
          src="/images/home/ic-1.png" 
          alt="images-ic-1" 
          className="absolute -top-7 left-4 sm:left-8 size-16 sm:size-20 md:size-28"
          animate={{
            y: [0, -8, 0],
          }}
          transition={{
            duration: 2,
            repeat: Infinity,
            ease: "easeInOut"
          }}
        />
        <motion.img 
          src="/images/home/ic-2.png" 
          alt="images-ic-2" 
          className="absolute bottom-1/2 -right-6 sm:-right-10 h-10 sm:h-12 md:h-16 rotate-12"
          animate={{
            y: [0, -8, 0],
          }}
          transition={{
            duration: 2,
            repeat: Infinity,
            ease: "easeInOut"
          }}
        />
        
        {/* Heading with fade in animation */}
        <motion.h3 
          className="text-3xl sm:text-4xl md:text-5xl lg:text-6xl font-Inter font-semibold text-white text-center leading-snug tracking-wide mb-12 sm:mb-14 md:mb-16 px-4"
          variants={fadeInUp}
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true, amount: 0.3 }}
        >
          We aim to help children{" "}
          <span className="text-yellow-400 italic block mt-2">
            discover the joy of creative
          </span>
          learning and grow into well-rounded individuals.
        </motion.h3>

        {/* Author grid with stagger animation */}
        <motion.div 
          className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 sm:gap-8 justify-center"
          variants={staggerContainer}
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true, amount: 0.2 }}
        >
          {arrAuthor.map((author, index) => (
            <motion.div
              key={author.id}
              className="flex flex-col items-center text-center p-4 sm:p-6 rounded-2xl"
              variants={playfulCardEnter(index)}
              whileHover={playfulHover}
            >
              <motion.img
                src={author.image}
                alt={author.name}
                className="h-auto w-48 sm:w-56 md:w-64 object-cover rounded-full shadow-lg"
                whileHover={{ scale: 1.1, rotate: 5 }}
                transition={{ duration: 0.3 }}
              />
              <h4 className="mt-4 text-xl sm:text-2xl font-semibold text-white">
                {author.name}
              </h4>
              <p className="text-sm text-gray-100">{author.role}</p>
            </motion.div>
          ))}
        </motion.div>
      </div>
    </div>
  )
}

export default AuthorSection