import React, { useEffect, useState, useRef, useLayoutEffect } from "react";
import { Blurhash } from "react-blurhash";
import clsx from "clsx";

interface ImageProps extends React.ImgHTMLAttributes<HTMLImageElement> {
  blurhash?: string;
  width: number;
  height: number;
}

const Image: React.FC<ImageProps> = ({
  src,
  alt = "image",
  blurhash,
  width,
  height,
  className,
  ...rest
}) => {
  const [loading, setLoading] = useState(true);
  const [aspectRatio, setAspectRatio] = useState(`${width}/${height}`);
  const ref = useRef<HTMLDivElement>(null);

  const [wi, setWi] = useState(0);
  const [hi, setHi] = useState(0);

  useEffect(() => {
    if (!src) return;

    const image = new window.Image();
    image.onload = () => {
      setLoading(false);
    };
    image.src = src;
  }, [src]);

  useEffect(() => {
    setAspectRatio(`${width}/${height}`);
  }, [height, width]);

  useLayoutEffect(() => {
    if (ref.current) {
      setWi(ref.current.offsetWidth);
      setHi(ref.current.offsetHeight);
    }
  }, []);

  return (
    <div
      ref={ref}
      className={clsx("relative overflow-hidden", className)}
      style={{ aspectRatio }}
    >
      {loading ? (
        blurhash ? (
          <Blurhash
            hash={blurhash}
            width={wi}
            height={hi}
            resolutionX={64}
            resolutionY={64}
            punch={1}
          />
        ) : (
          <div className="w-full h-full blur-md animate-pulse" />
        )
      ) : (
        <img
          src={src}
          alt={alt}
          className="w-full h-full object-cover"
          {...rest}
        />
      )}
    </div>
  );
};

export default Image;
