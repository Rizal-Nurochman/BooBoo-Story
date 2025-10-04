import { Minus, Plus } from "lucide-react";
import { useState } from "react";
import { motion, AnimatePresence } from "motion/react";
import {
  playfulCardEnter,
  playfulHover,
  playfulTap,
  rotateIcon,
  wiggleIcon,
  dropdownSlide,
  backgroundSlide,
  slideInLeft,
  fadeInUp,
  spinEnter,
  spinEnterReverse,
  iconHover,
  iconHoverReverse,
  floatingAnimation
} from "@/lib/animations";

const faqs = [
  {
    question: "Apa yang membuat WonderKids berbeda dari platform pendidikan lainnya?",
    answer: "WonderKids dirancang khusus untuk anak-anak dengan pendekatan pembelajaran interaktif dan menyenangkan. Platform ini memadukan teknologi, permainan edukatif, serta materi yang disesuaikan dengan usia sehingga anak-anak lebih mudah memahami dan menikmati proses belajar."
  },
  {
    question: "Bagaimana cara saya mengakses WonderKids?",
    answer: "Anda dapat mengakses WonderKids melalui website resmi atau mengunduh aplikasinya di perangkat mobile. Setelah mendaftar dan membuat akun, Anda bisa langsung menggunakan berbagai fitur pembelajaran yang tersedia."
  },
  {
    question: "Bagaimana dengan keamanan data anak-anak yang menggunakan platform ini?",
    answer: "Keamanan data anak-anak adalah prioritas utama di WonderKids. Semua data dilindungi dengan enkripsi dan sistem keamanan tingkat tinggi, serta mengikuti standar perlindungan data internasional agar privasi anak tetap terjaga."
  }
];

interface DropdownFaqProps {
  parent: string;
  child: string;
  index: number;
}

const DropdownFaq = ({ parent, child, index }: DropdownFaqProps) => {
  const [open, setOpen] = useState(false);
  const [isHovered, setIsHovered] = useState(false);
  
  return (
    <motion.div
      variants={playfulCardEnter(index)}
      initial="hidden"
      whileInView="visible"
      viewport={{ once: true, amount: 0.3 }}
      whileHover={playfulHover}
      whileTap={playfulTap}
      onHoverStart={() => setIsHovered(true)}
      onHoverEnd={() => setIsHovered(false)}
      className="w-full text-white rounded-3xl px-6 py-4 border-2 bg-foreground cursor-pointer transition-all hover:border-[#9B7BFF] hover:shadow-lg hover:shadow-[#9B7BFF]/20 relative overflow-hidden"
      onClick={() => setOpen(prev => !prev)}
    >
      <motion.div
        className="absolute inset-0 bg-gradient-to-r from-[#9B7BFF]/10 to-transparent"
        initial={{ x: "-100%" }}
        animate={backgroundSlide(isHovered)}
      />
      
      <div className="flex gap-4 justify-between items-start relative z-10">
        <p className="font-medium text-base sm:text-lg flex-1 pr-2">{parent}</p>
        <motion.div
          animate={rotateIcon(open)}
          className="flex-shrink-0 mt-1"
        >
          <motion.div animate={wiggleIcon(isHovered, open)}>
            {open ? <Minus size={20} /> : <Plus size={20} />}
          </motion.div>
        </motion.div>
      </div>
      
      <AnimatePresence mode="wait">
        {open && (
          <motion.div
            variants={dropdownSlide}
            initial="hidden"
            animate="visible"
            exit="exit"
            className="overflow-hidden relative z-10"
          >
            <motion.p 
              initial={{ opacity: 0 }}
              animate={{ opacity: 1 }}
              transition={{ delay: 0.2, duration: 0.3 }}
              className="text-sm sm:text-base text-gray-300 leading-relaxed"
            >
              {child}
            </motion.p>
          </motion.div>
        )}
      </AnimatePresence>
    </motion.div>
  );
};

const FaqSection = () => {
  return (
    <div className="w-full max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 pt-20 sm:pt-24 md:pt-32 pb-6 sm:pb-8 md:pb-12 lg:pb-16">
      <div className="flex flex-col lg:flex-row gap-8 lg:gap-12 items-start lg:items-center">
        <motion.div
          variants={slideInLeft}
          initial="hidden"
          whileInView="visible"
          viewport={{ once: true, amount: 0.3 }}
          className="flex-1 w-full flex flex-col items-center justify-center"
        >
          <motion.h3 
            variants={fadeInUp}
            initial="hidden"
            whileInView="visible"
            viewport={{ once: true }}
            transition={{ delay: 0.2 }}
            className="text-2xl sm:text-3xl lg:text-4xl xl:text-5xl font-semibold font-Inter leading-tight"
          >
            Frequently <span className="italic text-[#9B7BFF]">asked</span> questions
          </motion.h3>
          
          <div className="flex gap-4 sm:gap-6 items-center mt-6 sm:mt-8">
            <motion.img
              variants={spinEnter(0.4)}
              initial="hidden"
              whileInView="visible"
              viewport={{ once: true }}
              whileHover={iconHover}
              src="/images/home/ic-1.png"
              className="h-24 sm:h-32 lg:h-40 w-auto object-contain cursor-pointer"
              alt="img-ic-1"
            />
            <motion.img
              variants={spinEnterReverse(0.6)}
              initial="hidden"
              whileInView="visible"
              viewport={{ once: true }}
              whileHover={iconHoverReverse}
              animate={floatingAnimation}
              src="/images/home/ic-2.png"
              alt="img-ic-2"
              className="h-16 sm:h-20 lg:h-24 w-auto object-contain cursor-pointer"
            />
          </div>
        </motion.div>

        <div className="flex-1 w-full flex flex-col gap-4 sm:gap-5 lg:gap-6">
          {faqs.map((faq, index) => (
            <DropdownFaq
              key={index}
              parent={faq.question}
              child={faq.answer}
              index={index}
            />
          ))}
        </div>
      </div>
    </div>
  );
};

export default FaqSection;