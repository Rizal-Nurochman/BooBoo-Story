import React, { useState, useEffect, useRef } from "react";

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

  // Tutup dropdown jika klik di luar area
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
      {/* Tombol utama */}
      <button
  onClick={() => setOpen(!open)}
  className="flex items-center gap-2 text-[#8C6A4E] font-semibold text-lg hover:text-[#B15A33] transition"
>
  <span className="flex items-center">
    Jelajahi
    <span className="inline-block ml-1 align-middle text-xl">🌞</span>
  </span>
  <svg
    className={`w-4 h-4 transform transition-transform ${
      open ? "rotate-180" : "rotate-0"
    }`}
    fill="none"
    stroke="currentColor"
    strokeWidth="2"
    viewBox="0 0 24 24"
  >
    <path strokeLinecap="round" strokeLinejoin="round" d="M19 9l-7 7-7-7" />
  </svg>
</button>


      {/* Dropdown */}
      {open && (
        <div
          className="absolute top-full left-0 mt-2 flex rounded-2xl shadow-lg border border-[#8C6A4E] overflow-hidden z-50 
                     bg-gradient-to-b from-[#F6E9C8] to-[#F1E0C0]"
        >
          {/* Kolom kiri */}
          <div className="p-4 w-64 bg-gradient-to-b from-[#EED3B1] to-[#F6E9C8]">
            <h3 className="font-bold text-[#5B4231] mb-3">KATEGORI CERITA</h3>
            <ul className="space-y-1 text-[#5B4231] text-sm">
              {categoriesLeft.map((item) => (
                <li
                  key={item}
                  className="hover:text-[#B15A33] hover:bg-[#EED3B1]/40 rounded-lg px-2 py-1 cursor-pointer transition-all"
                >
                  {item}
                </li>
              ))}
            </ul>
          </div>

          {/* Kolom kanan */}
          <div className="p-4 w-56 bg-gradient-to-b from-[#E49A6B]/40 to-[#F6E9C8] border-l border-[#C09373]">
            <h3 className="font-bold text-[#5B4231] mb-3">PILIHAN CERITA</h3>
            <ul className="space-y-1 text-[#5B4231] text-sm">
              {categoriesRight.map((item) => (
                <li
                  key={item}
                  className="hover:text-[#B15A33] hover:bg-[#EED3B1]/40 rounded-lg px-2 py-1 cursor-pointer transition-all"
                >
                  {item}
                </li>
              ))}
            </ul>
          </div>
        </div>
      )}
    </div>
  );
};

export default Category;
