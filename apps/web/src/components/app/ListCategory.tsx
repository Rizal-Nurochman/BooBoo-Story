import React, { useState, useEffect, useRef } from "react";
import { motion, AnimatePresence } from "motion/react";

const Category: React.FC = () => {
  const [open, setOpen] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);

  const categoriesLeft = [
    "Dongeng Hewan 🦊",
    "Petualangan 🧭",
    "Fantasi & Sihir ✨",
    "Persahabatan 💖",
    "Kerajaan & Pahlawan 👑",
    "Misteri Ringan 🔍",
    "Cerita Edukatif 📚",
    "Keluarga & Kasih Sayang 🏡",
    "Cerita Lucu 😂",
    "Alam & Lingkungan 🌳",
    "Nilai Moral 🌟",
  ];

  const categoriesRight = [
    "Favorit Minggu Ini 💫",
    "Cerita Populer 🔥",
    "Daftar Cerita 🎓",
  ];

  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      if (dropdownRef.current && !dropdownRef.current.contains(e.target as Node)) {
        setOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  return (
    <div ref={dropdownRef} className="relative">
      <motion.button
        onClick={() => setOpen(!open)}
        whileHover={{ scale: 1.02 }}
        whileTap={{ scale: 0.98 }}
        className="flex items-center gap-2 text-[#8C6A4E] font-semibold text-lg hover:text-[#B15A33] transition-colors duration-200"
      >
        <span className="flex items-center">
          Jelajahi
          <motion.span
            animate={{ rotate: open ? [0, -10, 10, -10, 0] : 0 }}
            transition={{ duration: 0.5 }}
            className="inline-block ml-1 align-middle text-xl"
          >
            🌞
          </motion.span>
        </span>
        <motion.svg
          animate={{ rotate: open ? 180 : 0 }}
          transition={{ duration: 0.3, ease: "easeInOut" }}
          className="w-4 h-4"
          fill="none"
          stroke="currentColor"
          strokeWidth="2"
          viewBox="0 0 24 24"
        >
          <path strokeLinecap="round" strokeLinejoin="round" d="M19 9l-7 7-7-7" />
        </motion.svg>
      </motion.button>

      <AnimatePresence>
        {open && (
          <motion.div
            initial={{ opacity: 0, y: -20, scale: 0.95 }}
            animate={{ opacity: 1, y: 0, scale: 1 }}
            exit={{ opacity: 0, y: -10, scale: 0.95 }}
            transition={{
              type: "spring",
              stiffness: 300,
              damping: 25,
              mass: 0.5
            }}
            className="absolute top-full left-0 mt-2 flex rounded-2xl shadow-2xl border border-[#8C6A4E] overflow-hidden z-50 
                       bg-gradient-to-b from-[#F6E9C8] to-[#F1E0C0]"
          >
            <motion.div
              initial={{ x: -20, opacity: 0 }}
              animate={{ x: 0, opacity: 1 }}
              transition={{ delay: 0.1, duration: 0.3 }}
              className="p-4 w-64 bg-gradient-to-b from-[#EED3B1] to-[#F6E9C8]"
            >
              <h3 className="font-bold text-[#5B4231] mb-3">KATEGORI CERITA</h3>
              <ul className="space-y-1 text-[#5B4231] text-sm">
                {categoriesLeft.map((item, index) => (
                  <motion.li
                    key={item}
                    initial={{ x: -20, opacity: 0 }}
                    animate={{ x: 0, opacity: 1 }}
                    transition={{ 
                      delay: 0.15 + index * 0.03,
                      duration: 0.3,
                      ease: "easeOut"
                    }}
                    whileHover={{ 
                      x: 4,
                      backgroundColor: "rgba(238, 211, 177, 0.4)",
                      transition: { duration: 0.2 }
                    }}
                    className="hover:text-[#B15A33] rounded-lg px-2 py-1 cursor-pointer transition-colors duration-200"
                  >
                    {item}
                  </motion.li>
                ))}
              </ul>
            </motion.div>

            <motion.div
              initial={{ x: 20, opacity: 0 }}
              animate={{ x: 0, opacity: 1 }}
              transition={{ delay: 0.2, duration: 0.3 }}
              className="p-4 w-56 bg-gradient-to-b from-[#E49A6B]/40 to-[#F6E9C8] border-l border-[#C09373]"
            >
              <h3 className="font-bold text-[#5B4231] mb-3">PILIHAN CERITA</h3>
              <ul className="space-y-1 text-[#5B4231] text-sm">
                {categoriesRight.map((item, index) => (
                  <motion.li
                    key={item}
                    initial={{ x: 20, opacity: 0 }}
                    animate={{ x: 0, opacity: 1 }}
                    transition={{ 
                      delay: 0.25 + index * 0.05,
                      duration: 0.3,
                      ease: "easeOut"
                    }}
                    whileHover={{ 
                      x: 4,
                      backgroundColor: "rgba(238, 211, 177, 0.4)",
                      transition: { duration: 0.2 }
                    }}
                    className="hover:text-[#B15A33] rounded-lg px-2 py-1 cursor-pointer transition-colors duration-200"
                  >
                    {item}
                  </motion.li>
                ))}
              </ul>
            </motion.div>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  );
};

export default Category;