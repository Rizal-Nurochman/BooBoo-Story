import React from "react";
import { ChevronRight } from "lucide-react"; 
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

const BookList: React.FC<BookListProps> = ({ title, books, onSeeMore }) => {
  return (
    <section className="mb-10">
      {/* Header section */}
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-xl sm:text-2xl font-bold text-[#5B4231]">
          {title}
        </h2>
        <button
          onClick={onSeeMore}
          className="flex items-center gap-1 text-[#B15A33] text-sm sm:text-base font-medium hover:text-[#8C4A2B] transition-all"
        >
          Lihat Semua
          <ChevronRight className="w-4 h-4" />
        </button>
      </div>

      {/* Daftar buku */}
      <div className="flex overflow-x-auto gap-4 pb-2 scrollbar-hide">
        {books.map((book, idx) => (
          <div
            key={idx}
            className="min-w-[150px] sm:min-w-[180px] bg-[#FFFFFF] rounded-xl shadow-md border border-[#EFD9B5] 
                       hover:shadow-lg hover:-translate-y-1 transition-all"
          >
            <img
              src={book.cover}
              alt={book.title}
              className="w-full h-40 object-cover rounded-t-xl"
            />
            <div className="p-3">
              <h3 className="text-[#5B4231] font-semibold text-sm sm:text-base">
                {book.title}
              </h3>
              <p className="text-[#8C6A4E] text-xs">{book.author}</p>
              <span className="inline-block text-[#B15A33] text-xs font-medium mt-1">
                {book.category}
              </span>
            </div>
          </div>
        ))}
      </div>
    </section>
  );
};

export default BookList;
