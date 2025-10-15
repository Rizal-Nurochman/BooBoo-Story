import React, { useState } from "react";
import { createFileRoute, useNavigate } from "@tanstack/react-router";

export const Route = createFileRoute("/app/books/profile")({
  component: ProfilePage,
});

function ProfilePage() {
  const navigate = useNavigate();
  const [activeTab, setActiveTab] = useState<"perihal" | "percakapan" | "mengikuti">("perihal");

  return (
    <div className="min-h-screen bg-[#F6E9C8] text-[#4A3A25]">
      {/* ===== HEADER ===== */}
      <div className="bg-[#9C7C63] h-56 flex flex-col items-center justify-center text-center relative">
        <button
          onClick={() => navigate({ to: "/books" })}
          className="absolute top-5 left-5 text-[#F6E9C8] hover:text-white transition flex items-center gap-2"
        >
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={2} stroke="currentColor" className="w-5 h-5">
            <path strokeLinecap="round" strokeLinejoin="round" d="M15 19l-7-7 7-7" />
          </svg>
          Kembali
        </button>

        <div className="bg-[#B18E73] text-[#F6E9C8] w-16 h-16 rounded-full flex items-center justify-center text-xl font-bold mb-2 shadow-md">
          hy
        </div>
        <h1 className="text-xl font-semibold text-[#F6E9C8]">hy rain</h1>
        <p className="text-[#F6E9C8]/90 text-sm">@hyrain5</p>

        <div className="flex gap-8 text-[#F6E9C8] mt-3 text-sm">
          <div className="flex flex-col items-center">
            <span className="font-bold text-lg">0</span>
            <span>Karya</span>
          </div>
          <div className="flex flex-col items-center">
            <span className="font-bold text-lg">2</span>
            <span>Daftar Bacaan</span>
          </div>
          <div className="flex flex-col items-center">
            <span className="font-bold text-lg">0</span>
            <span>Pengikut</span>
          </div>
        </div>
      </div>

      {/* ===== TAB NAVIGATION ===== */}
      <div className="flex justify-center border-b border-[#D7BFA5] bg-[#F6E9C8]">
        {["perihal", "percakapan", "mengikuti"].map((tab) => (
          <button
            key={tab}
            onClick={() => setActiveTab(tab as any)}
            className={`px-6 py-3 font-semibold capitalize transition ${
              activeTab === tab
                ? "border-b-4 border-[#B15A33] text-[#B15A33]"
                : "text-[#5B4231] hover:text-[#B15A33]"
            }`}
          >
            {tab}
          </button>
        ))}
      </div>

      {/* ===== TAB CONTENT ===== */}
      <div className="p-6 sm:px-12">
        {activeTab === "perihal" && (
          <div className="grid sm:grid-cols-3 gap-6 mt-4">
            {/* Deskripsi */}
            <div className="bg-[#FFF8E6] p-6 rounded-2xl shadow-md border border-[#EFD9B5]">
              <h2 className="font-semibold mb-3">Bantu pengguna lain mengenal dirimu</h2>
              <button className="bg-[#B15A33] hover:bg-[#8C4A2B] text-[#F6E9C8] font-semibold px-4 py-2 rounded-lg transition">
                Tambahkan deskripsi
              </button>
              <p className="mt-4 text-sm text-[#7A5D4A]">
                Bergabung <strong>September 14, 2023</strong>
              </p>

              <div className="mt-4">
                <h3 className="font-semibold text-sm mb-2">Bagikan Profil</h3>
                <div className="flex gap-3">
                  {["facebook", "twitter", "pinterest", "tumblr", "email"].map((icon) => (
                    <div
                      key={icon}
                      className="w-8 h-8 bg-[#D7BFA5] hover:bg-[#B18E73] rounded-full flex items-center justify-center cursor-pointer transition"
                    >
                      <span className="text-[#FFF8E6] text-sm capitalize">{icon[0]}</span>
                    </div>
                  ))}
                </div>
              </div>
            </div>

            {/* Placeholder untuk karya / daftar */}
            <div className="sm:col-span-2 bg-[#FFF8E6] p-6 rounded-2xl shadow-md border border-[#EFD9B5] flex flex-col items-center justify-center text-center">
              <div className="w-full h-36 bg-[#EFD9B5] rounded-lg mb-3 animate-pulse"></div>
              <button className="bg-[#B15A33] hover:bg-[#8C4A2B] text-[#F6E9C8] font-semibold px-5 py-2 rounded-lg transition">
                Buat Daftar Bacaan
              </button>
            </div>
          </div>
        )}

        {activeTab === "percakapan" && (
          <div className="text-center text-[#7A5D4A] mt-6">
            <p>Belum ada percakapan.</p>
          </div>
        )}

        {activeTab === "mengikuti" && (
          <div className="text-center text-[#7A5D4A] mt-6">
            <p>Kamu belum mengikuti siapa pun.</p>
          </div>
        )}
      </div>
    </div>
  );
}

export default ProfilePage;
