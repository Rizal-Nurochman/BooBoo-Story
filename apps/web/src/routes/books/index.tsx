import React, { useState } from "react";
import Category from "@/components/books/Category";
import Find from "@/components/books/Find";
import BookList from "@/components/books/BookList";
import { createFileRoute, useNavigate } from "@tanstack/react-router";

const allBooks = [
  { title: "Petualangan Si Kelinci", author: "Rani A.", cover: "/books/book1.jpg", category: "Dongeng Hewan" },
  { title: "Rahasia Hutan Ajaib", author: "Taqin R.", cover: "/books/book2.jpg", category: "Fantasi & Sihir" },
  { title: "Sahabat di Tepi Danau", author: "Rizal N.", cover: "/books/book3.jpg", category: "Persahabatan" },
  { title: "Istana Di Balik Awan", author: "Alya P.", cover: "/books/book4.jpg", category: "Kerajaan & Pahlawan" },
];

export const Route = createFileRoute("/books/")({
  component: RouteComponent,
});

function RouteComponent() {
  const navigate = useNavigate();
  const [filteredBooks, setFilteredBooks] = useState(allBooks);

  const handleSearch = (query: string) => {
    if (query.trim() === "") {
      setFilteredBooks(allBooks);
      return;
    }
    const result = allBooks.filter((book) =>
      book.title.toLowerCase().includes(query.toLowerCase())
    );
    setFilteredBooks(result);
  };

  return (
    <div className="pt-28 px-6 sm:px-12 bg-[#F6E9C8] min-h-screen">
      <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4 sm:gap-6">
        <div className="flex flex-col sm:flex-row sm:items-center gap-3 sm:gap-6 w-full sm:w-auto">
          <Category />
          <Find onSearch={handleSearch} />
        </div>

        <div className="flex flex-row items-center justify-start sm:justify-end gap-3 w-full sm:w-auto">
          <button
            onClick={() => navigate({ to: "/books/create" })}
            className="bg-[#B6C497] hover:bg-[#8FA97B] text-[#4A3A25] font-semibold px-4 py-2 rounded-lg shadow-sm 
                       transition-all hover:scale-105 w-full sm:w-auto"
          >
            + Buat Cerita
          </button>

          <button
            onClick={() => navigate({ to: "/books/profile" })}
            className="bg-[#B15A33] hover:bg-[#8C4A2B] text-[#F6E9C8] font-semibold px-4 py-2 rounded-lg shadow-sm 
                       transition-all hover:scale-105 w-full sm:w-auto"
          >
            Profil Saya
          </button>
        </div>
      </div>

      {/* ================= DAFTAR BUKU ================= */}
      <div className="mt-12 space-y-10">
        <BookList title="📖 Hasil Pencarian" books={filteredBooks} />
      </div>
    </div>
  );
}

export default RouteComponent;
