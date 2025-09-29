import { features, tags } from "@/constants";
import BadgeShape from "../ui/BadgeShape";
import { Card } from "../ui/card";
import { motion } from "motion/react";
import { 
  fadeIn,
  slideUp,
  slideRight,
  scaleIn,
  rotateIn,
  staggerContainer
} from "@/lib/animations";

const FeatureSection = () => {
  return (
    <div className="w-full bg-[#FBF9F2] py-8 sm:py-12 md:py-16 lg:py-20">
      <motion.div 
        className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8"
        variants={staggerContainer}
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, margin: "-100px" }}
      >
        <div className="flex items-center justify-between gap-8 lg:gap-4">
          <motion.div 
            className="font-inter flex flex-col gap-1 md:gap-2 italic text-xl sm:text-4xl lg:text-5xl font-semibold"
            variants={slideUp}
          >
            <h2 className="text-black">Berbagai Fitur</h2>
            <h3 className="text-[#9B7BFF]">Menarik</h3>
          </motion.div>

          <motion.div 
            className="relative h-24 sm:h-28 lg:h-32 w-full max-w-[40%] sm:w-48 lg:w-56 rounded-full border-4 lg:border-5 -rotate-6 lg:-rotate-12"
            variants={rotateIn}
          >
            {tags.map((tag, index) => (
              <motion.div
                key={index}
                className={`px-3 sm:px-4 lg:px-5 py-1.5 sm:py-2 block w-fit rounded-full text-sm sm:text-base font-medium shadow-sm ${
                  index % 2 === 1 ? "ml-[60%] sm:ml-[70%] lg:ml-[76%]" : "-ml-[10%] sm:-ml-[12%] lg:-ml-[15%]"
                }`}
                style={{ backgroundColor: tag.color, color: tag.textCol }}
                variants={slideRight}
                whileHover={{ scale: 1.05, transition: { duration: 0.2 } }}
              >
                {tag.text}
              </motion.div>
            ))}
          </motion.div>
        </div>

        <motion.div 
          className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 mt-8 md:mt-12 lg:mt-16 gap-6 lg:gap-8"
          variants={staggerContainer}
        >
          {features.map((ft, index) => (
            <motion.div
              key={index}
              variants={scaleIn}
              whileHover={{ 
                y: -8, 
                scale: 1.02,
                transition: { duration: 0.3, ease: [0.25, 0.1, 0.25, 1] }
              }}
            >
              <Card
                style={{ backgroundColor: ft.bgColor, color: ft.colorText }}
                className="relative border-none shadow-xl shadow-primary/10 rounded-2xl p-6 lg:p-8 overflow-hidden h-full"
              >
                <motion.div 
                  className="relative w-20 sm:w-24 lg:w-28 h-20 sm:h-24 lg:h-28 mb-6 lg:mb-8 flex items-center justify-center"
                  variants={scaleIn}
                >
                  <BadgeShape
                    color="rgba(255,255,255,0.3)"
                    className="absolute inset-0 w-full h-full"
                  />
                  <motion.span 
                    className="text-2xl sm:text-3xl lg:text-4xl relative z-10"
                    whileHover={{ 
                      rotate: 10, 
                      scale: 1.1,
                      transition: { duration: 0.2 }
                    }}
                  >
                    {ft.icon}
                  </motion.span>
                </motion.div>

                <div className="w-full mx-auto max-w-[90%]">
                  <motion.h3 
                    className="text-2xl sm:text-3xl lg:text-4xl tracking-wide font-bold"
                    variants={slideUp}
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    transition={{ delay: 0.1 }}
                  >
                    {ft.title}
                  </motion.h3>
                  <motion.h4 
                    className="text-lg sm:text-xl lg:text-2xl mb-3 lg:mb-4 mt-1 font-semibold tracking-wide"
                    variants={slideUp}
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    transition={{ delay: 0.2 }}
                  >
                    {ft.highlight}
                  </motion.h4>
                  <motion.p 
                    className="text-base sm:text-lg tracking-wider font-medium"
                    variants={fadeIn}
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    transition={{ delay: 0.3 }}
                  >
                    {ft.description}
                  </motion.p>
                </div>

                <motion.img
                  src={ft.image}
                  alt={ft.highlight}
                  className="absolute w-24 sm:w-28 lg:w-32 -top-6 sm:-top-7 lg:-top-8 right-0"
                  variants={slideRight}
                  whileHover={{ 
                    scale: 1.1, 
                    rotate: 5,
                    transition: { duration: 0.3 }
                  }}
                />
              </Card>
            </motion.div>
          ))}
        </motion.div>
      </motion.div>
    </div>
  );
};

export default FeatureSection;