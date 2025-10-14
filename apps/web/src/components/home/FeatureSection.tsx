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
    <div className="w-full bg-[#FBF9F2] py-6 sm:py-10 md:py-14 lg:py-18">
      <motion.div 
        className="max-w-6xl mx-auto px-3 sm:px-5 lg:px-6"
        variants={staggerContainer}
        initial="hidden"
        whileInView="visible"
        viewport={{ once: true, margin: "-80px" }}
      >
        <div className="flex items-center justify-between gap-6 lg:gap-3">
          <motion.div 
            className="font-inter flex flex-col gap-1 md:gap-1.5 italic text-lg sm:text-3xl lg:text-4xl font-semibold"
            variants={slideUp}
          >
            <h2 className="text-black">Berbagai Fitur</h2>
            <h3 className="text-[#9B7BFF]">Menarik</h3>
          </motion.div>

          <motion.div 
            className="relative h-20 sm:h-24 lg:h-28 w-full max-w-[45%] sm:w-40 lg:w-48 rounded-full border-4 lg:border-[5px] -rotate-6 lg:-rotate-10"
            variants={rotateIn}
          >
            {tags.map((tag, index) => (
              <motion.div
                key={index}
                className={`px-2.5 sm:px-3.5 lg:px-4 py-1 sm:py-1.5 block w-fit rounded-full text-xs sm:text-sm font-medium shadow-sm ${
                  index % 2 === 1 ? "ml-[55%] sm:ml-[65%] lg:ml-[72%]" : "-ml-[8%] sm:-ml-[10%] lg:-ml-[12%]"
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
          className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 mt-6 md:mt-10 lg:mt-12 gap-5 lg:gap-6"
          variants={staggerContainer}
        >
          {features.map((ft, index) => (
            <motion.div
              key={index}
              variants={scaleIn}
              whileHover={{ 
                y: -6, 
                scale: 1.02,
                transition: { duration: 0.3, ease: [0.25, 0.1, 0.25, 1] }
              }}
            >
              <Card
                style={{ backgroundColor: ft.bgColor, color: ft.colorText }}
                className="relative border-none shadow-lg shadow-primary/10 rounded-xl p-3 lg:p-5 overflow-hidden h-full"
              >
                <motion.div 
                  className="relative w-14 sm:w-16 lg:w-20 h-14 sm:h-16 lg:h-20 mb-5 lg:mb-4 flex items-center justify-center"
                  variants={scaleIn}
                >
                  <BadgeShape
                    color="rgba(255,255,255,0.3)"
                    className="absolute inset-0 w-full h-full"
                  />
                  <motion.span 
                    className="text-lg sm:text-xl lg:text-2xl relative z-10"
                    whileHover={{ 
                      rotate: 10, 
                      scale: 1.1,
                      transition: { duration: 0.2 }
                    }}
                  >
                    {ft.icon}
                  </motion.span>
                </motion.div>

                <div className="w-full mx-auto max-w-[92%]">
                  <motion.h3 
                    className="text-lg sm:text-xl lg:text-2xl tracking-wide font-bold"
                    variants={slideUp}
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    transition={{ delay: 0.1 }}
                  >
                    {ft.title}
                  </motion.h3>
                  <motion.h4 
                    className="text-sm sm:text-md lg:text-lg mb-2 lg:mb-3 mt-1 font-semibold tracking-wide"
                    variants={slideUp}
                    initial="hidden"
                    whileInView="visible"
                    viewport={{ once: true }}
                    transition={{ delay: 0.2 }}
                  >
                    {ft.highlight}
                  </motion.h4>
                  <motion.p 
                    className="text-xs sm:text-sm tracking-wider font-medium"
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
                  className="absolute w-20 sm:w-24 lg:w-28 -top-5 sm:-top-6 lg:-top-7 right-0"
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
