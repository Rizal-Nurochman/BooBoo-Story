import { Link } from "@tanstack/react-router";
import { MoveUpRight } from "lucide-react";
import { motion } from "motion/react";
import {
  fadeIn,
  slideRight,
  slideUp,
  staggerContainer,
} from "@/lib/animations";

const childrenData = [
  {
    id: 1,
    name: "Child 1",
    image: "/images/home/child-1.png",
    icon: "/images/home/ic-1.png",
    background: "#E5D4FF",
    width: "65%",
  },
  {
    id: 2,
    name: "Child 2",
    image: "/images/home/child-2.png",
    icon: "/images/home/ic-2.png",
    background: "#FFD966",
    width: "70%",
  },
  {
    id: 3,
    name: "Child 3",
    image: "/images/home/child-3.png",
    icon: "/images/home/ic-3.png",
    background: "#9B7BFF",
    width: "65%",
  },
];

const HighlightSection = () => {
  return (
    <div className="py-12 md:py-20 lg:py-32 w-full max-w-7xl mx-auto px-4 md:px-6">
      <motion.div
        variants={staggerContainer}
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, amount: 0.2 }}
        className="flex flex-col justify-center lg:flex-row items-center gap-8 lg:gap-12"
      >
        <motion.div
          variants={slideRight}
          className="flex-1 space-y-4 md:space-y-6 text-center lg:text-left"
        >
          <h3 className="text-2xl md:text-3xl lg:text-4xl xl:text-5xl font-medium font-Inter leading-tight">
            Materi pembelajaran yang disediakan{" "}
            <span className="text-[#9B7BFF] font-semibold">menyenangkan</span>{" "}
            untuk anak-anak.
          </h3>
          <motion.p
            variants={fadeIn}
            className="text-base md:text-lg lg:text-xl xl:text-2xl font-medium text-gray-500 leading-relaxed"
          >
            Don't worry! Your children will be having a fun time while learning
            with our materials that are easy to understand.
          </motion.p>

          <motion.div variants={slideUp}>
            <Link
              to="/login"
              className="inline-flex gap-3 md:gap-4 items-center border-2 border-[#9B7BFF]/80 py-2 md:py-3 px-4 md:px-6 text-[#9B7BFF] rounded-full cursor-pointer text-base md:text-lg font-semibold hover:bg-[#9B7BFF]/5 transition-colors duration-300"
            >
              Learn More
              <span className="size-8 md:size-10 text-white rounded-full flex items-center justify-center bg-[#9B7BFF] hover:scale-110 transition-transform duration-300">
                <MoveUpRight className="w-4 h-4 md:w-5 md:h-5" />
              </span>
            </Link>
          </motion.div>
        </motion.div>

        <motion.div
          variants={fadeIn}
          className="flex-1 w-full rootate- max-w-xl flex flex-col gap-4 md:gap-6 lg:gap-8 items-center lg:items-end"
        >
          {childrenData.map((child, index) => (
            <motion.div
              key={child.id}
              initial={{ opacity: 0, x: 50 }}
              whileInView={{ opacity: 1, x: 0 }}
              viewport={{ once: true }}
              transition={{ duration: 0.6, delay: index * 0.15 }}
              className="rounded-full flex items-center justify-between px-3 py-2 md:px-4 md:py-3 lg:px-6 lg:py-4 overflow-visible relative"
              style={{
                backgroundColor: child.background,
                width: child.width,
                minWidth: "280px",
              }}
            >
              <div className="flex relative gap-2 md:gap-4 items-center w-full overflow-visible">
                <motion.div
                  whileHover={{ scale: 1.05 }}
                  transition={{ duration: 0.3 }}
                  className="flex-shrink-0 -my-8 md:-my-10 lg:-my-12 overflow-visible"
                >
                  <img
                    src={child.image}
                    className="h-[200px] w-auto object-contain"
                    alt={child.name}
                  />
                </motion.div>

                <motion.div
                  whileHover={{ scale: 1.1, rotate: 10 }}
                  whileTap={{ scale: 0.95 }}
                  transition={{ duration: 0.3 }}
                  className={`flex-shrink-0 absolute ${
                    index % 2 === 1
                      ? "top-1/2 -translate-y-1/2 -left-12 md:-left-16"
                      : "-top-8 md:-top-12 -right-16 md:-right-20"
                  }`}
                >
                  <img
                    src={child.icon}
                    className="object-contain w-32 md:w-40 h-auto"
                    alt={`${child.name} icon`}
                  />
                </motion.div>
              </div>
            </motion.div>
          ))}
        </motion.div>
      </motion.div>
    </div>
  );
};

export default HighlightSection;
