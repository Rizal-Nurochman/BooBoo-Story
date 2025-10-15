import React, { useState } from "react";
import { createFileRoute, useNavigate } from "@tanstack/react-router";

export const Route = createFileRoute("/app/_appLayout/books/create")({
  component: CreateStory,
});

function CreateStory() {
  const navigate = useNavigate();
  const [cover, setCover] = useState<string | null>(null);
  const [title, setTitle] = useState("");
  const [desc, setDesc] = useState("");
  const [mainChar, setMainChar] = useState("");

  // Handle upload cover
  const handleImageChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      setCover(URL.createObjectURL(file));
    }
  };

  // Handle submit form
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log({ title, desc, mainChar, cover });
    navigate({ to: "/books" }); // nanti bisa diarahkan ke daftar cerita
  };

  return (
    <div className="pt-28 px-6 py-20 sm:px-12 bg-[#F6E9C8] min-h-screen text-[#5B4231]">
      {/* Tombol Back */}
      <button
        onClick={() => navigate({ to: "/books" })}
        className="flex items-center gap-2 mb-6 text-[#5B4231] font-medium hover:text-[#B15A33] transition"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          strokeWidth={2}
          stroke="currentColor"
          className="w-5 h-5"
        >
          <path strokeLinecap="round" strokeLinejoin="round" d="M15 19l-7-7 7-7" />
        </svg>
        Kembali
      </button>

      {/* Judul halaman */}
      <h1 className="text-3xl font-bold mb-8">✨ Buat Cerita Baru</h1>

      {/* Form utama */}
      <form
        onSubmit={handleSubmit}
        className="flex flex-col sm:flex-row gap-6 bg-[#FFF8E6] p-6 rounded-2xl shadow-md border border-[#EFD9B5]"
      >
        {/* KIRI: Tambah cover */}
        <div className="flex flex-col items-center w-full sm:w-1/3">
          <label
            htmlFor="cover-upload"
            className="w-48 h-64 flex items-center justify-center bg-[#EFD9B5] hover:bg-[#EAC98F] rounded-xl border border-[#C5A574] cursor-pointer overflow-hidden transition"
          >
            {cover ? (
              <img
                src={cover}
                alt="Sampul Cerita"
                className="w-full h-full object-cover"
              />
            ) : (
              <span className="text-[#8C6A4E] font-medium text-sm text-center">
                + Tambahkan Sampul 📚
              </span>
            )}
          </label>
          <input
            id="cover-upload"
            type="file"
            accept="image/*"
            className="hidden"
            onChange={handleImageChange}
          />
        </div>

        {/* KANAN: Form detail */}
        <div className="flex-1 space-y-5">
          {/* Judul */}
          <div>
            <label className="block font-semibold mb-1">Judul Cerita</label>
            <input
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              placeholder="Masukkan judul ceritamu..."
              className="w-full p-3 rounded-lg border border-[#EFD9B5] bg-[#FFFDF8] focus:ring-2 focus:ring-[#B6C497] outline-none"
              required
            />
          </div>

          {/* Deskripsi */}
          <div>
            <label className="block font-semibold mb-1">Deskripsi</label>
            <textarea
              value={desc}
              onChange={(e) => setDesc(e.target.value)}
              placeholder="Tulis deskripsi singkat tentang ceritamu..."
              rows={4}
              className="w-full p-3 rounded-lg border border-[#EFD9B5] bg-[#FFFDF8] focus:ring-2 focus:ring-[#B6C497] outline-none"
              required
            />
          </div>

          {/* Tokoh utama */}
          <div>
            <label className="block font-semibold mb-1">Tokoh Utama</label>
            <input
              type="text"
              value={mainChar}
              onChange={(e) => setMainChar(e.target.value)}
              placeholder="Siapa tokoh utama di ceritamu?"
              className="w-full p-3 rounded-lg border border-[#EFD9B5] bg-[#FFFDF8] focus:ring-2 focus:ring-[#B6C497] outline-none"
            />
          </div>

          {/* Tombol aksi */}
          <div className="flex justify-end gap-3 pt-4">
            <button
              type="button"
              onClick={() => navigate({ to: "/books" })}
              className="bg-[#EAC98F] hover:bg-[#E7B974] text-[#5B4231] px-4 py-2 rounded-lg font-semibold shadow-sm transition"
            >
              Batal
            </button>
            <button
              type="submit"
              className="bg-[#B15A33] hover:bg-[#8C4A2B] text-[#FFF8E6] px-4 py-2 rounded-lg font-semibold shadow-sm transition"
            >
              Simpan Cerita
            </button>
          </div>
        </div>
      </form>
    </div>
  );
}

export default CreateStory;
