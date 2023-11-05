import React, { useCallback, useEffect, useRef, useState } from 'react';
import { Images } from '../../constants/images';


const MAX_COUNT = 4;
const videosArray = Object.values(Images.GlitchPunks);

const PunkCarousel = () => {
  const videoRef = useRef(null);
  const [currentVideo, setCurrentVideo] = useState(0);
  const [isFadeOut, setIsFadeOut] = useState(false);
  const isUp = useRef(false);

  const onVideoEnd = useCallback(() => {
    setIsFadeOut(true);
   setTimeout(() => {
     setCurrentVideo((prev) => {
       return !isUp.current ? prev + 1 : prev - 1;
     });
     setIsFadeOut(false);
   }, 400)
  }, [currentVideo, isUp]);

  useEffect(() => {
    if (currentVideo >= MAX_COUNT) {
      isUp.current = true;
    } else if (currentVideo <= 0) {
      isUp.current = false;
    }
  }, [currentVideo]);

  useEffect(() => {
      // const preLoadImage = new Image();
      //
      //
      // for(let i = 0; i <= videosArray.length; i++){
      //   preLoadImage.src= videosArray[i];
      // }
  }, [])

  useEffect(() => {
    videoRef.current = setInterval(() => {
      onVideoEnd();
    }, 5000);

    return () => videoRef.current && clearInterval(videoRef.current)
  }, [])

  return (
  <>
    <div className="w-[100vw] relative overflow-x-visible mt-[-20px]">
      <img style={{  }} className="absolute heading-img" src={Images.BackgroundHeading} alt=""/>
      <div className="relative h-[300px] w-full flex flex-row justify-center">
        <img
            key={currentVideo}
            style={{ top: 0, height: 320, width: 267, bottom: -50, zIndex: 2, objectFit: "contain" }}
            className={`absolute fade-in ${isFadeOut && 'fade-out'}`}
            src={videosArray[currentVideo]}
            alt=""
        />
      </div>
    </div>
</>

  );
};

//<video
//             key={currentVideo}
//             onEnded={onVideoEnd}
//             controls={false}
//             muted
//             autoPlay
//             style={{ top: 0, height: 320, width: 267, bottom: -50, zIndex: 2 }}
//             className={`absolute video-source-def fade-in ${isFadeOut ? "fade-out" : ""}`}>
//           <source
//               ref={videoRef}
//               style={{ background: "transparent" }}
//               src={videosArray[currentVideo]}
//               type="video/webm" />
//         </video>

export default PunkCarousel;
