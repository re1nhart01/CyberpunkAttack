import React from 'react';
import { Images, Images as images } from '../../constants/images';
import Ellipse from '../../images/ellips.png';
import EllipseView from '../EllipseView/EllipseView';

const PunkCarousel = () => {
  return (
    <>
      <img
        loading="lazy"
        decoding="async"
        style={{ top: 0, height: 600, width: 600, zIndex: 1 }}
        className="absolute"
        src={images.BackgroundHeading}
        alt=""
      />
      <div className="w-[100vw] relative overflow-x-visible mt-[-12px]">
        {/* <EllipseView /> */}
        <div className="relative h-[300px] w-full flex flex-row justify-center">
          <img
            loading="lazy"
            decoding="async"
            style={{ top: 0, height: 320, width: 267, bottom: -50, zIndex: 2 }}
            className="absolute"
            src={Images.GlitchPunk4}
            alt=""
          />
        </div>
      </div>
    </>

  );
};

export default PunkCarousel;
