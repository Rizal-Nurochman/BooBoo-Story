import React from "react";

interface BadgeShapeProps {
  color?: string;
  className?: string;
}

const BadgeShape: React.FC<BadgeShapeProps> = ({ color = "#FFFFFF", className }) => {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 100 100"
      className={className}
      fill={color}
    >
      <path d="M50 0 
        C60 10, 90 10, 100 20 
        C90 30, 100 60, 90 70 
        C80 80, 60 90, 50 100 
        C40 90, 20 80, 10 70 
        C0 60, 10 30, 0 20 
        C10 10, 40 10, 50 0Z" />
    </svg>
  );
};

export default BadgeShape;
