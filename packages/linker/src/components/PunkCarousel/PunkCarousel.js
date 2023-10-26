import React from 'react';
import { Images } from '../../constants/images';
import Ellipse from '../../images/ellips.png';
import EllipseView from "../EllipseView/EllipseView";

const PunkCarousel = () => {
  return (
    <>
    <EllipseView />
      <div className="relative h-[300px] w-full flex flex-row justify-center">
        <img style={{ top: 0, height: 360, width: 277, bottom: -50 }} className="absolute" src={Images.GlitchPunk4} alt="" />
      </div>
    </>
  );
};

export default PunkCarousel;
