import { Variants } from "motion/react";

// 🔑 Definisi cubic bezier sebagai tuple (bukan number[])
const easeOutFast: [number, number, number, number] = [0.25, 0.1, 0.25, 1];
const easeInOut: [number, number, number, number] = [0.4, 0, 0.2, 1];
const easeOutSlow: [number, number, number, number] = [0.22, 1, 0.36, 1];
const easeOut: [number, number, number, number] = [0.25, 0.1, 0.25, 1];
const easeIn: [number, number, number, number] = [0.42, 0, 1, 1];

export const fadeIn: Variants = {
  hidden: { opacity: 0 },
  visible: {
    opacity: 1,
    transition: {
      duration: 0.6,
      ease: easeOutFast
    }
  }
};

export const slideUp: Variants = {
  hidden: { opacity: 0, y: 30 },
  visible: {
    opacity: 1,
    y: 0,
    transition: {
      duration: 0.7,
      ease: easeOutFast
    }
  }
};

export const slideRight: Variants = {
  hidden: { opacity: 0, x: -30 },
  visible: {
    opacity: 1,
    x: 0,
    transition: {
      duration: 0.7,
      ease: easeOutFast
    }
  }
};

export const scaleIn: Variants = {
  hidden: { opacity: 0, scale: 0.8 },
  visible: {
    opacity: 1,
    scale: 1,
    transition: {
      duration: 0.5,
      ease: easeInOut
    }
  }
};

export const rotateIn: Variants = {
  hidden: { opacity: 0, rotate: -15 },
  visible: {
    opacity: 1,
    rotate: 0,
    transition: {
      duration: 0.6,
      ease: easeOutFast
    }
  }
};

export const staggerContainer: Variants = {
  hidden: {},
  visible: {
    transition: {
      staggerChildren: 0.15,
      ease: easeOutSlow
    }
  }
};


export const buttonTap = {
  scale: 0.98,
  transition: { duration: 0.1 }
};



export const slideDown: Variants = {
  hidden: { opacity: 0, y: -50 },
  visible: {
    opacity: 1,
    y: 0,
    transition: {
      duration: 0.6,
      ease: easeOut, 
    },
  },
  exit: {
    opacity: 0,
    y: -50,
    transition: {
      duration: 0.4,
      ease: easeIn, 
    },
  },
};

