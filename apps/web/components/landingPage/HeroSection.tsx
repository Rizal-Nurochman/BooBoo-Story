'use client'
import Image from 'next/image'
import React from 'react'
import Link from 'next/link'
import { motion } from 'motion/react'

const springs = [
  {
    img: '/images/home/spring.png',
    class: 'w-28 h-auto absolute top-[70%] -right-[5%]'
  },
  {
    img: '/images/home/spring.png',
    class: 'w-28 h-auto absolute top-[70%] -left-[5%]'
  },
  {
    img: '/images/home/spring.png',
    class: 'w-28 h-auto absolute top-[30%] -left-[10%]'
  }
]

const containerVariants = {
  hidden: { opacity: 0 },
  visible: {
    opacity: 1,
    transition: {
      delayChildren: 0.2,
      staggerChildren: 0.15
    }
  }
}

const springVariants = {
  hidden: { 
    opacity: 0,
    scale: 0,
    y: -100,
    rotate: 360
  },
  visible: (i: number) => ({
    opacity: 1,
    scale: 1,
    y: 0,
    rotate: 0,
    transition: {
      type: "spring" as const,
      stiffness: 400,
      damping: 15,
      delay: i * 0.2,
      bounce: 0.6
    }
  })
}

const titleVariants = {
  hidden: { 
    opacity: 0,
    y: 50,
    scale: 0.8
  },
  visible: {
    opacity: 1,
    y: 0,
    scale: 1,
    transition: {
      type: "spring" as const,
      stiffness: 100,
      damping: 15
    }
  }
}

const imageVariants = {
  hidden: { 
    opacity: 0,
    scale: 0.3,
    rotateY: -90
  },
  visible: {
    opacity: 1,
    scale: 1,
    rotateY: 0,
    transition: {
      type: "spring" as const,
      stiffness: 100,
      damping: 20
    }
  }
}

const subtitleVariants = {
  hidden: { 
    opacity: 0,
    y: 30,
    scale: 0.9
  },
  visible: {
    opacity: 1,
    y: 0,
    scale: 1,
    transition: {
      type: "spring" as const,
      stiffness: 120,
      damping: 12
    }
  }
}

const buttonContainerVariants = {
  hidden: { opacity: 0 },
  visible: {
    opacity: 1,
    transition: {
      delayChildren: 0.3,
      staggerChildren: 0.2
    }
  }
}

const buttonVariants = {
  hidden: { 
    opacity: 0,
    y: 40,
    scale: 0.8
  },
  visible: {
    opacity: 1,
    y: 0,
    scale: 1,
    transition: {
      type: "spring" as const,
      stiffness: 200,
      damping: 15
    }
  }
}

const HeroSection = () => {
  return (
    <motion.div
      variants={containerVariants}
      initial="hidden"
      animate="visible"
      className="w-full max-w-6xl mx-auto pt-24 pb-10 space-y-8 relative"
    >
      {springs.map((img, i) => (
        <motion.div
          key={i}
          variants={springVariants}
          custom={i}
        >
          <Image
            src={img.img}
            className={img.class}
            width={50}
            height={20}
            alt={'img-spring' + i}
          />
        </motion.div>
      ))}

      <motion.h1
        variants={titleVariants}
        className="text-xl sm:text-2xl md:text-4xl 2xl:text-5xl text-center font-bold text-primary"
      >
        Belajar, bermain, dan berkembang bersama BooBoo Story!
      </motion.h1>

      <motion.div variants={imageVariants}>
        <Image
          src={'/images/home/landing-page.svg'}
          alt="childenrs image"
          width={200}
          height={200}
          className="w-full max-w-3xl mx-auto h-auto"
        />
      </motion.div>

      <motion.p
        variants={subtitleVariants}
        className="text-4xl font-semibold text-primary text-center"
      >
        Siapkah kamu belajar bersama BooBoo?
      </motion.p>

      <motion.div
        variants={buttonContainerVariants}
        className="flex gap-4 items-center justify-center"
      >
        <Link href="/register">
          <motion.button
            variants={buttonVariants}
            whileHover={{ 
              scale: 1.05,
              y: -2,
              boxShadow: "0 10px 25px rgba(0,0,0,0.1)"
            }}
            whileTap={{ scale: 0.95 }}
            className="
              relative overflow-hidden
              px-6 py-3 text-2xl font-semibold
              text-background bg-primary
              rounded-lg cursor-pointer
              transition-all duration-300 ease-out
              hover:shadow-lg hover:bg-primary/90
            "
          >
            <span className="relative z-10">Get Started</span>
          </motion.button>
        </Link>
        <Link href="/login">
          <motion.button
            variants={buttonVariants}
            whileHover={{ 
              scale: 1.05,
              y: -2,
              boxShadow: "0 10px 25px rgba(0,0,0,0.1)"
            }}
            whileTap={{ scale: 0.95 }}
            className="
              relative overflow-hidden
              px-6 py-3 text-2xl font-semibold
              text-primary ring-2 ring-primary
              rounded-lg cursor-pointer
              transition-all duration-300 ease-out
              hover:shadow-lg hover:bg-primary hover:text-background
            "
          >
            <span className="relative z-10">Join With Us</span>
          </motion.button>
        </Link>
      </motion.div>
    </motion.div>
  )
}

export default HeroSection