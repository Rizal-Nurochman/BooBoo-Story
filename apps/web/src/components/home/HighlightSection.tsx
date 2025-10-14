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
    name: "Anak 1",
    image: "/images/home/child-1.png",
    icon: "/images/home/ic-1.png",
    background: "#E5D4FF",
    width: "55%",
  },
  {
    id: 2,
    name: "Anak 2",
    image: "/images/home/child-2.png",
    icon: "/images/home/ic-2.png",
    background: "#FFD966",
    width: "60%",
  },
  {
    id: 3,
    name: "Anak 3",
    image: "/images/home/child-3.png",
    icon: "/images/home/ic-3.png",
    background: "#9B7BFF",
    width: "55%",
  },
];

const HighlightSection = () => {
  return (
    <div className="py-8 md:py-12 lg:py-16 w-full max-w-6xl mx-auto px-3 md:px-5">
      <motion.div
        variants={staggerContainer}
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, amount: 0.2 }}
        className="flex flex-col justify-center lg:flex-row items-center gap-4 g:gap-6 flex-1"
      >
        <motion.div
          variants={slideRight}
          className="flex-1 space-y-3 md:space-y-4 text-center lg:text-left"
        >
          <h3 className="text-xl md:text-2xl lg:text-3xl xl:text-4xl font-medium font-Inter leading-snug">
            Materi pembelajaran yang disediakan{" "}
            <motion.span
              initial={{ scale: 0.8, opacity: 0 }}
              whileInView={{ scale: 1, opacity: 1 }}
              viewport={{ once: true }}
              transition={{ duration: 0.6, ease: "easeOut" }}
              className="relative inline-flex items-center"
            >
              <span className="relative z-10 text-[#9B7BFF] font-semibold px-2 py-0.5">
                menyenangkan
              </span>
              <motion.span
                initial={{ rotate: -10, scale: 0 }}
                whileInView={{ rotate: -2, scale: 1 }}
                viewport={{ once: true }}
                transition={{ duration: 0.7, delay: 0.2, ease: "easeOut" }}
                className="absolute inset-0 rounded-full border-[3px] border-yellow-500"
              />
            </motion.span>{" "}
            untuk anak-anak.
          </h3>

          <motion.p
            variants={fadeIn}
            className="text-sm md:text-base lg:text-lg xl:text-xl font-medium text-gray-500 leading-relaxed"
          >
            Jangan khawatir! Anak Anda akan belajar sambil bersenang-senang
            dengan materi kami yang mudah dipahami.
          </motion.p>

          <motion.div variants={slideUp}>
            <Link
              to="/auth/login"
              className="inline-flex gap-2.5 md:gap-3 items-center border-2 border-[#9B7BFF]/80 py-1.5 md:py-2 px-3.5 md:px-5 text-[#9B7BFF] rounded-full cursor-pointer text-sm md:text-base font-semibold hover:bg-[#9B7BFF]/5 transition-colors duration-300"
            >
              Pelajari Lebih Lanjut
              <span className="size-7 md:size-8 text-white rounded-full flex items-center justify-center bg-[#9B7BFF] hover:scale-110 transition-transform duration-300">
                <MoveUpRight className="w-3.5 h-3.5 md:w-4 md:h-4" />
              </span>
            </Link>
          </motion.div>
        </motion.div>

        <motion.div
          variants={fadeIn}
          className="flex-1 w-full max-w-sm flex flex-col gap-3 md:gap-4 lg:gap-6 items-center lg:items-center"
        >
          {childrenData.map((child, index) => (
            <motion.div
              key={child.id}
              initial={{ opacity: 0, x: 50 }}
              whileInView={{ opacity: 1, x: 0 }}
              viewport={{ once: true }}
              transition={{ duration: 0.6, delay: index * 0.15 }}
              className="rounded-full flex items-center justify-between px-2.5 py-1.5 md:px-3.5 md:py-2.5 lg:px-5 lg:py-3 relative"
              style={{
                backgroundColor: child.background,
                width: child.width,
                minWidth: "240px",
              }}
            >
              <div className="flex relative gap-2 md:gap-3 items-center w-full">
                <motion.div
                  whileHover={{ scale: 1.05 }}
                  transition={{ duration: 0.3 }}
                  className="flex-shrink-0 -my-6 md:-my-8 lg:-my-10"
                >
                  <img
                    src={child.image}
                    className="h-[150px] w-auto object-contain"
                    alt={child.name}
                  />
                </motion.div>

                <motion.div
                  whileHover={{ scale: 1.1, rotate: 10 }}
                  whileTap={{ scale: 0.95 }}
                  transition={{ duration: 0.3 }}
                  className={`flex-shrink-0 absolute ${
                    index % 2 === 1
                      ? "top-1/2 -translate-y-1/2 -left-10 md:-left-14"
                      : "-top-6 md:-top-10 -right-14 md:-right-16"
                  }`}
                >
                  <img
                    src={child.icon}
                    className="object-contain w-20 md:w-24 h-auto"
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
