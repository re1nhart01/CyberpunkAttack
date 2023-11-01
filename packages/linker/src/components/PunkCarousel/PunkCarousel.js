import React, {useCallback, useEffect, useRef, useState} from 'react';
import {Images as images, Videos} from '../../constants/images';


const MAX_COUNT = 5;
const videosArray = Object.values(Videos);

const PunkCarousel = () => {
  const videoRef = useRef(null);
  const [currentVideo, setCurrentVideo] = useState(0);
  const [isFadeOut, setIsFadeOut] = useState(false);
  const isUp = useRef(false);

  const onVideoEnd = useCallback(() => {
    setIsFadeOut(true);
   setTimeout(() => {
     if (currentVideo >= MAX_COUNT) {
       isUp.current = true;
     } else if (currentVideo <= 0) {
       isUp.current = false;
     }
     setCurrentVideo((prev) => {
       return !isUp.current ? prev + 1 : prev - 1;
     });
     setIsFadeOut(false);
   }, 400)
  }, [currentVideo, isUp]);

  useEffect(() => {
    console.log(currentVideo, videosArray[currentVideo]);
  }, [currentVideo])

  return (
<>
    <img  style={{ top: 0, height: 600, width: 600, zIndex: 1, }} className="absolute" src={images.BackgroundHeading} alt=""/>
    <div className="w-[100vw] relative overflow-x-visible mt-[-20px]">
    {/*<EllipseView />*/}
      <div className="relative h-[300px] w-full flex flex-row justify-center">
        <video
            key={currentVideo}
            onEnded={onVideoEnd}
            controls={false}
            muted
            autoPlay
            style={{ top: 0, height: 320, width: 267, bottom: -50, zIndex: 2 }}
            className={`absolute video-source-def fade-in ${isFadeOut ? "fade-out" : ""}`}>
          <source
              ref={videoRef}
              style={{ background: "transparent" }}
              src={videosArray[currentVideo]}
              type="video/webm" />
        </video>
      </div>
    </div>
</>

  );
};

export default PunkCarousel;
