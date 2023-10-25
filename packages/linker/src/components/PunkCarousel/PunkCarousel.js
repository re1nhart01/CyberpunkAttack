import React from 'react';
import { Images } from '../../constants/images';
import Ellipse from '../../images/ellips.png';

const PunkCarousel = () => {
  return (
    <>
      <img className="absolute w-[200vw] h-[600px] aspect-square" src={Images.Ellipse} alt="" />
      <div className="relative h-[300px] w-full flex flex-row justify-center">
        <img style={{ top: 0, height: 360, width: 277, bottom: -50 }} className="absolute" src={Images.GlitchPunk4} alt="" />
      </div>
    </>
  );
};

export default PunkCarousel;
