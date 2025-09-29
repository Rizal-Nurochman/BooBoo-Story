import { useEffect, useState } from "react";
import { ArrowUp } from "lucide-react";
import { Button } from "../ui/button";

const BackToTop = () => {
  const [visible, setVisible] = useState(false);

  useEffect(() => {
    const toggleVisibility = () => {
      setVisible(window.scrollY > 200);
    };

    window.addEventListener("scroll", toggleVisibility);
    return () => window.removeEventListener("scroll", toggleVisibility);
  }, []);

  const scrollToTop = () => {
    window.scrollTo({
      top: 0,
      behavior: "smooth",
    });
  };

  if (!visible) return null;

  return (
    <Button
      onClick={scrollToTop}
      aria-label="Scroll to top"
      size={'icon'}
      className="
        fixed bottom-6 right-6 
        md:bottom-10 md:right-10 
        lg:bottom-12 lg:right-12
        bg-primary opacity-75 text-white shadow-lg
        hover:opacity-100 active:scale-95
        transition-all duration-300
      "
    >
      <ArrowUp className="w-6 h-6 md:w-5 md:h-5" />
    </Button>
  );
};

export default BackToTop;
