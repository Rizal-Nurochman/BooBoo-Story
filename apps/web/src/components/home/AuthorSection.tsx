import { motion, Variants } from "motion/react"

const arrAuthor = [
  { id: 1, name: "Kristin Watson", image: "/images/home/pic-1.png", role: "Chef" },
  { id: 2, name: "Jenny Wilson", image: "/images/home/pic-2.png", role: "Artist" },
  { id: 3, name: "Jacob Jones", image: "/images/home/pic-3.png", role: "Teacher" },
  { id: 4, name: "Savannah Nguyen", image: "/images/home/pic-4.png", role: "Writer" },
]

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
  y: -6,
  rotate: 0.5,
  transition: { 
    duration: 0.3, 
    ease: easeOutFast 
  }
}

const AuthorSection = () => {
  return (
    <div className="bg-[#E08F62] pt-12 sm:pt-16 md:pt-20 pb-4 sm:pb-6 md:pb-8 lg:pb-10 overflow-hidden">
      <div className="w-full max-w-5xl mx-auto px-3 sm:px-4 lg:px-6 relative">
        <motion.img 
          src="/images/home/ic-1.png" 
          alt="images-ic-1" 
          className="absolute -top-5 left-3 sm:left-6 size-10 sm:size-14 md:size-20"
          animate={{
            y: [0, -6, 0],
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
          className="absolute bottom-1/2 -right-4 sm:-right-6 h-8 sm:h-10 md:h-12 rotate-12"
          animate={{
            y: [0, -6, 0],
          }}
          transition={{
            duration: 2,
            repeat: Infinity,
            ease: "easeInOut"
          }}
        />
        
        <motion.h3 
          className="text-xl sm:text-2xl md:text-3xl lg:text-4xl font-Inter font-semibold text-white text-center leading-snug tracking-wide mb-8 sm:mb-10 md:mb-12 px-2"
          variants={fadeInUp}
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true, amount: 0.3 }}
        >
          Kami ingin mengajak anak-anak
          <span className="text-yellow-400 italic block mt-1 text-base sm:text-lg">menjelajahi dunia imajinasi melalui cerita-cerita seru</span>
          dan menemukan keseruan belajar dari setiap kisah yang menginspirasi.
        </motion.h3>

        <motion.div 
          className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 sm:gap-6 justify-center"
          variants={staggerContainer}
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true, amount: 0.2 }}
        >
          {arrAuthor.map((author, index) => (
            <motion.div
              key={author.id}
              className="flex flex-col items-center text-center p-3 sm:p-4 rounded-xl"
              variants={playfulCardEnter(index)}
              whileHover={playfulHover}
            >
              <motion.img
                src={author.image}
                alt={author.name}
                className="h-auto w-32 sm:w-40 md:w-48 object-cover rounded-full shadow-lg"
                whileHover={{ scale: 1.1, rotate: 5 }}
                transition={{ duration: 0.3 }}
              />
              <h4 className="mt-3 text-lg sm:text-xl font-semibold text-white">
                {author.name}
              </h4>
              <p className="text-xs text-gray-100">{author.role}</p>
            </motion.div>
          ))}
        </motion.div>
      </div>
    </div>
  )
}

export default AuthorSection
