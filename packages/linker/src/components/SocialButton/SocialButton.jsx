import React from 'react';

const SocialButton = ({ label, icon, url }) => {
  const goTo = () => {
    window.location.href = url;
  };

  return (
    <div className="w-full mb-3 flex flex-row justify-center items-center">
      <button className="social-button__container" onClick={goTo}>
        <div className="social-button__container_inner social-button-text">
          <img className="w-[24px] h-[24px]" src={icon} alt={`SOCIAL_BUTTON_${url.substring(8, 16)}`} />
          { label }
        </div>
      </button>
    </div>
  );
};

export default SocialButton;
