import React from 'react';
import {Images as images, Images} from '../../constants/images';
import Ellipse from '../../images/ellips.png';
import EllipseView from "../EllipseView/EllipseView";

const PunkCarousel = () => {
  return (
<>
    <img  style={{ top: 0, height: 600, width: 600, zIndex: 1, }} className="absolute" src={images.BackgroundHeading} alt=""/>
    <div className="w-[100vw] relative overflow-x-visible mt-[-20px]">
    {/*<EllipseView />*/}
      <div className="relative h-[300px] w-full flex flex-row justify-center">
        <img style={{ top: 0, height: 320, width: 267, bottom: -50, zIndex: 2 }} className="absolute" src={Images.GlitchPunk4} alt="" />
      </div>
    </div>
</>

  );
};

export default PunkCarousel;
