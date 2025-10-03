import { Variants } from "motion/react";

const easeOutFast: [number, number, number, number] = [0.25, 0.1, 0.25, 1];
const easeInOut: [number, number, number, number] = [0.4, 0, 0.2, 1];
const easeOutSlow: [number, number, number, number] = [0.22, 1, 0.36, 1];
const easeOut: [number, number, number, number] = [0.25, 0.1, 0.25, 1];
const easeIn: [number, number, number, number] = [0.42, 0, 1, 1];
const bounceEase: [number, number, number, number] = [0.68, -0.55, 0.265, 1.55];

export const fadeInUp: Variants = {
  hidden: { opacity: 0, y: 20 },
  visible: {
    opacity: 1,
    y: 0,
    transition: {
      duration: 0.6,
      ease: easeOutFast
    }
  }
};

// Slide animations
export const slideInLeft: Variants = {
  hidden: { opacity: 0, x: -100 },
  visible: {
    opacity: 1,
    x: 0,
    transition: {
      duration: 0.8,
      ease: bounceEase
    }
  }
};

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


export const playfulCardEnter = (index: number): Variants => ({
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
});

// Gamification hover effects
export const playfulHover = {
  scale: 1.02,
  y: -4,
  rotate: 0.5,
  transition: { 
    duration: 0.3, 
    ease: easeOutFast 
  }
};

export const playfulTap = {
  scale: 0.98,
  transition: { duration: 0.1 }
};

// Icon animations
export const spinEnter = (delay: number = 0): Variants => ({
  hidden: { 
    opacity: 0, 
    scale: 0.5, 
    rotate: -45 
  },
  visible: {
    opacity: 1, 
    scale: 1, 
    rotate: 0,
    transition: { 
      duration: 0.7, 
      delay,
      ease: bounceEase,
      type: "spring",
      stiffness: 200
    }
  }
});

export const spinEnterReverse = (delay: number = 0): Variants => ({
  hidden: { 
    opacity: 0, 
    scale: 0.5, 
    rotate: 45 
  },
  visible: {
    opacity: 1, 
    scale: 1, 
    rotate: 0,
    transition: { 
      duration: 0.7, 
      delay,
      ease: bounceEase,
      type: "spring",
      stiffness: 200
    }
  }
});

export const iconHover = {
  scale: 1.1, 
  rotate: 5,
  y: -10,
  transition: { duration: 0.3 }
};

export const iconHoverReverse = {
  scale: 1.15, 
  rotate: -5,
  y: -8,
  transition: { duration: 0.3 }
};

// Floating animation
export const floatingAnimation = {
  y: [0, -8, 0],
  transition: {
    duration: 2,
    repeat: Infinity,
    ease: "easeInOut" as const
  }
};

// Icon rotation animations
export const rotateIcon = (open: boolean) => ({
  rotate: open ? 180 : 0,
  scale: open ? 1.1 : 1,
  transition: { 
    duration: 0.4, 
    ease: bounceEase,
    type: "spring" as const,
    stiffness: 300
  }
});

export const wiggleIcon = (isHovered: boolean, open: boolean) => ({
  rotate: isHovered && !open ? [0, -10, 10, -10, 0] : 0,
  transition: {
    duration: 0.5,
    ease: easeOut
  }
});

// Dropdown animations
export const dropdownSlide = {
  hidden: { 
    opacity: 0, 
    height: 0, 
    marginTop: 0,
    y: -10
  },
  visible: { 
    opacity: 1, 
    height: "auto", 
    marginTop: 16,
    y: 0,
    transition: {
      height: { duration: 0.4, ease: easeOut },
      opacity: { duration: 0.3, delay: 0.1, ease: easeOut },
      y: { duration: 0.4, ease: bounceEase }
    }
  },
  exit: { 
    opacity: 0, 
    height: 0, 
    marginTop: 0,
    y: -10,
    transition: {
      height: { duration: 0.3, ease: easeIn },
      opacity: { duration: 0.2, ease: easeIn },
      y: { duration: 0.3, ease: easeIn }
    }
  }
};

// Background slide effect
export const backgroundSlide = (isHovered: boolean) => ({
  x: isHovered ? "0%" : "-100%",
  transition: { duration: 0.3, ease: easeOut }
});
