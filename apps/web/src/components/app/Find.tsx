import { Search, X } from "lucide-react";
import React, { useEffect, useState, useRef } from "react";
import { motion, AnimatePresence } from "motion/react";
import { useNavigate, useSearch } from "@tanstack/react-router";
import { Button } from "../ui/button";

const Find: React.FC = () => {
  const navigate = useNavigate();
  const search = useSearch({ from: "/app/_appLayout/books/" });
  const [isOpen, setIsOpen] = useState(false);
  const [query, setQuery] = useState(() => search.q || "");
  const [isFocused, setIsFocused] = useState(false);
  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    if (query !== (search.q || "")) setQuery(search.q || "");
  }, [search.q]);

  useEffect(() => {
    const timeoutId = setTimeout(() => {
      const nextQuery = query.trim();
      const currentQuery = search.q || "";
      if (nextQuery === currentQuery) return;
      if (nextQuery) {
        navigate({
          to: "/app/books",
          search: { q: nextQuery, page: search.page, limit: search.limit },
        });
        localStorage.setItem("book_search", nextQuery);
      } else {
        navigate({
          to: "/app/books",
          search: { page: search.page, limit: search.limit },
        });
        localStorage.removeItem("book_search");
      }
    }, 300);
    return () => clearTimeout(timeoutId);
  }, [query]);

  useEffect(() => {
    const saved = localStorage.getItem("book_search");
    if (saved && !search.q) setQuery(saved);
  }, []);

  const handleOpen = () => {
    setIsOpen(true);
    setTimeout(() => inputRef.current?.focus(), 50);
  };

  const handleClose = () => {
    if (!query.trim()) {
      setIsOpen(false);
      setIsFocused(false);
    }
  };

  const handleClear = () => {
    setQuery("");
    inputRef.current?.focus();
  };

  return (
    <div className="relative flex items-center">
      <motion.div
        animate={{
          width: isOpen ? "20rem" : "2.75rem",
          paddingLeft: isOpen ? "2.75rem" : "0rem",
          paddingRight: isOpen
            ? query.trim()
              ? "2.75rem"
              : "0.75rem"
            : "0rem",
          borderRadius: isOpen ? "9999px" : "50%",
        }}
        transition={{
          type: "spring",
          stiffness: 260,
          damping: 22,
          mass: 0.6,
        }}
        className="relative flex items-center overflow-hidden border-2 bg-gradient-to-br from-[var(--secondary)]/80 to-[var(--secondary)]/60 backdrop-blur-md rounded-full shadow-lg hover:shadow-xl transition-shadow duration-300"
        style={{
          borderColor: isFocused ? "var(--primary)" : "var(--border)",
        }}
      >
        <AnimatePresence mode="wait">
          {isOpen ? (
            <motion.div
              key="search-icon"
              initial={{ scale: 0.8, opacity: 0 }}
              animate={{ scale: 1, opacity: 1 }}
              exit={{ scale: 0.8, opacity: 0 }}
              transition={{ duration: 0.2 }}
              className="absolute left-3 top-1/2 -translate-y-1/2"
            >
              <Search className="text-[var(--primary)] w-5 h-5" />
            </motion.div>
          ) : (
            <motion.div
              key="button"
              initial={{ scale: 0.9 }}
              animate={{ scale: 1 }}
              whileHover={{ scale: 1.05 }}
              whileTap={{ scale: 0.95 }}
            >
              <Button
                onClick={handleOpen}
                size="icon"
                variant="ghost"
                className="cursor-pointer p-0 w-11 h-11 bg-gradient-to-br from-[var(--primary)]/10 to-[var(--primary)]/5 hover:from-[var(--primary)]/20 hover:to-[var(--primary)]/10 rounded-full transition-all duration-300"
              >
                <Search className="text-[var(--primary)] w-5 h-5" />
              </Button>
            </motion.div>
          )}
        </AnimatePresence>

        <AnimatePresence>
          {isOpen && (
            <motion.div
              key="input-wrapper"
              initial={{ opacity: 0, x: -20 }}
              animate={{ opacity: 1, x: 0 }}
              exit={{ opacity: 0, x: -20 }}
              transition={{ duration: 0.25, ease: "easeOut" }}
              className="w-full flex items-center"
            >
              <input
                ref={inputRef}
                type="text"
                value={query}
                onChange={(e) => setQuery(e.target.value)}
                onFocus={() => setIsFocused(true)}
                onBlur={() => {
                  setIsFocused(false);
                  handleClose();
                }}
                placeholder="Cari cerita..."
                className="bg-transparent w-full outline-none placeholder:text-[var(--muted-foreground)] py-2 text-[var(--foreground)] font-medium text-sm tracking-wide pr-2"
              />
            </motion.div>
          )}
        </AnimatePresence>

        <AnimatePresence>
          {isOpen && query.trim() && (
            <motion.button
              key="clear-button"
              initial={{ scale: 0, opacity: 0, rotate: -90 }}
              animate={{ scale: 1, opacity: 1, rotate: 0 }}
              exit={{ scale: 0, opacity: 0, rotate: 90 }}
              transition={{
                type: "spring",
                stiffness: 300,
                damping: 20,
              }}
              whileHover={{ scale: 1.1, rotate: 90 }}
              whileTap={{ scale: 0.9 }}
              onMouseDown={(e) => {
                e.preventDefault();
                handleClear();
              }}
              className="absolute right-2 top-1/2 -translate-y-1/2 p-1.5 rounded-full bg-[var(--muted)]/50 hover:bg-[var(--muted)] transition-colors duration-200"
            >
              <X className="w-3.5 h-3.5 text-[var(--muted-foreground)]" />
            </motion.button>
          )}
        </AnimatePresence>
      </motion.div>

      <AnimatePresence>
        {isOpen && isFocused && (
          <motion.div
            key="ripple"
            initial={{ scale: 0.8, opacity: 0.6 }}
            animate={{ scale: 1.5, opacity: 0 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.6, ease: "easeOut" }}
            className="absolute inset-0 rounded-full border-2 border-[var(--primary)]/30 pointer-events-none"
          />
        )}
      </AnimatePresence>
    </div>
  );
};

export default Find;
