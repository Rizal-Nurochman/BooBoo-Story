import { useState } from "react";
import BookList from "@/components/books/BookList";
import { createFileRoute, useNavigate } from "@tanstack/react-router";
import { bookSearchSchema } from "@/schemas/books.schema";

const allBooks = [
  { title: "Petualangan Si Kelinci", author: "Rani A.", cover: "/books/book1.jpg", category: "Dongeng Hewan" },
  { title: "Rahasia Hutan Ajaib", author: "Taqin R.", cover: "/books/book2.jpg", category: "Fantasi & Sihir" },
  { title: "Sahabat di Tepi Danau", author: "Rizal N.",cover: "/books/book3.jpg", category: "Persahabatan" },
  { title: "Istana Di Balik Awan", author: "Alya P.", cover: "/books/book4.jpg", category: "Kerajaan & Pahlawan" },
  { title: "Istana Di Balik Awan", author: "Alya P.", cover: "/books/book4.jpg", category: "Kerajaan & Pahlawan" },
  { title: "Istana Di Balik Awan", author: "Alya P.", cover: "/books/book4.jpg", category: "Kerajaan & Pahlawan" },
  { title: "Istana Di Balik Awan", author: "Alya P.", cover: "/books/book4.jpg", category: "Kerajaan & Pahlawan" },
  { title: "Istana Di Balik Awan", author: "Alya P.", cover: "/books/book4.jpg", category: "Kerajaan & Pahlawan" },
  { title: "Istana Di Balik Awan", author: "Alya P.", cover: "/books/book4.jpg", category: "Kerajaan & Pahlawan" },
];

export const Route = createFileRoute("/app/_appLayout/books/")({
  component: RouteComponent,
  validateSearch: (search) => {
    const result = bookSearchSchema.safeParse(search);
    if (!result.success) {
      throw new Error("Invalid search parameters");
    }
    return result.data;
  },
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
    <div className="px-6 sm:px-12 min-h-screen">

      {/* ================= DAFTAR BUKU ================= */}
      <div className="mt-12 space-y-10">
        <BookList title="📖 Hasil Pencarian" books={filteredBooks} />
      </div>
    </div>
  );
}

export default RouteComponent;
