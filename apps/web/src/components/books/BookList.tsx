import React, { useState } from "react";
import { ChevronRight } from "lucide-react";
import { motion } from "motion/react";

interface Book {
  title: string;
  author: string;
  cover: string;
  category: string;
}

interface BookListProps {
  title: string;
  books: Book[];
  onSeeMore?: () => void;
}

const BookCard: React.FC<{ book: Book; index: number }> = ({ book, index }) => {
  const [imageLoaded, setImageLoaded] = useState(false);
  const [imageError, setImageError] = useState(false);

  const placeholderImage = `https://images.unsplash.com/photo-1532012197267-da84d127e765?w=400&h=300&fit=crop`;

  return (
    <motion.div
      initial={{ opacity: 0, y: 20, scale: 0.95 }}
      animate={{ opacity: 1, y: 0, scale: 1 }}
      transition={{
        delay: index * 0.08,
        duration: 0.4,
        ease: "easeOut"
      }}
      whileHover={{
        y: -8,
        scale: 1.02,
        transition: { duration: 0.3, ease: "easeOut" }
      }}
      className="min-w-[150px] sm:min-w-[180px] bg-[#FFFFFF] rounded-xl shadow-md border border-[#EFD9B5] 
                 hover:shadow-2xl transition-shadow duration-300 cursor-pointer overflow-hidden"
    >
      <div className="relative w-full h-40 bg-gradient-to-br from-[#EFD9B5] to-[#F6E9C8] overflow-hidden">
        {!imageLoaded && !imageError && (
          <motion.div
            animate={{
              backgroundPosition: ["0% 0%", "100% 100%"],
            }}
            transition={{
              duration: 1.5,
              repeat: Infinity,
              ease: "linear"
            }}
            className="absolute inset-0 bg-gradient-to-r from-[#EFD9B5] via-[#F6E9C8] to-[#EFD9B5] bg-[length:200%_200%]"
          />
        )}
        
        <motion.img
          src={imageError ? placeholderImage : book.cover}
          alt={book.title}
          loading="lazy"
          onLoad={() => setImageLoaded(true)}
          onError={() => setImageError(true)}
          initial={{ opacity: 0, scale: 1.1 }}
          animate={{ 
            opacity: imageLoaded ? 1 : 0,
            scale: imageLoaded ? 1 : 1.1
          }}
          transition={{ duration: 0.5 }}
          className="w-full h-full object-cover"
        />
      </div>

      <motion.div
        initial={{ opacity: 0 }}
        animate={{ opacity: 1 }}
        transition={{ delay: index * 0.08 + 0.2, duration: 0.3 }}
        className="p-3"
      >
        <h3 className="text-[#5B4231] font-semibold text-sm sm:text-base line-clamp-2 mb-1">
          {book.title}
        </h3>
        <p className="text-[#8C6A4E] text-xs mb-2">{book.author}</p>
        <motion.span
          whileHover={{ scale: 1.05 }}
          className="inline-block text-[#B15A33] text-xs font-medium px-2 py-1 bg-[#FFF5E6] rounded-full"
        >
          {book.category}
        </motion.span>
      </motion.div>
    </motion.div>
  );
};

const BookList: React.FC<BookListProps> = ({ title, books, onSeeMore }) => {
  return (
    <section className="mb-10">
      <motion.div
        initial={{ opacity: 0, y: -10 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.4 }}
        className="flex items-center justify-between mb-4"
      >
        <h2 className="text-xl sm:text-2xl font-bold text-[#5B4231]">
          {title}
        </h2>
        <motion.button
          onClick={onSeeMore}
          whileHover={{ x: 4, scale: 1.05 }}
          whileTap={{ scale: 0.95 }}
          className="flex items-center gap-1 text-[#B15A33] text-sm sm:text-base font-medium hover:text-[#8C4A2B] transition-colors duration-200"
        >
          Lihat Semua
          <motion.div
            animate={{ x: [0, 4, 0] }}
            transition={{ duration: 1.5, repeat: Infinity, ease: "easeInOut" }}
          >
            <ChevronRight className="w-4 h-4" />
          </motion.div>
        </motion.button>
      </motion.div>

      <div className="grid grid-cols-4 gap-4">
        {books.map((book, idx) => (
          <BookCard key={idx} book={book} index={idx} />
        ))}
      </div>
    </section>
  );
};

export default BookList;