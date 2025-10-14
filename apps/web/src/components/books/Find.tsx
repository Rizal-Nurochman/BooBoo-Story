import { Search } from "lucide-react";
import React, { useState } from "react";

interface FindProps {
  onSearch: (query: string) => void;
}

const Find: React.FC<FindProps> = ({ onSearch }) => {
  const [query, setQuery] = useState("");

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setQuery(value);
    onSearch(value); // 🔍 langsung kirim ke parent tiap ketik
  };

  return (
    <div
      className="flex items-center bg-[#F6E9C8] border border-[#8C6A4E] rounded-full px-4 py-2 shadow-sm 
                 focus-within:ring-2 focus-within:ring-[#E49A6B] transition-all w-full max-w-sm"
    >
      <Search className="text-[#8C6A4E] text-lg mr-2" />
      <input
        type="text"
        value={query}
        onChange={handleChange}
        placeholder="Cari cerita berdasarkan judul..."
        className="bg-transparent outline-none text-[#5B4231] placeholder-[#A28767] w-full"
      />
    </div>
  );
};

export default Find;
