import React from 'react';

const SocialButtonSmall = ({ label, icon, url }) => {
  const goTo = () => {
    window.location.href = url;
  };

  return (
    <button className="social-button__container-small" onClick={goTo}>
      <img className="w-[24px] h-[24px]" src={icon} alt={`SOCIAL_BUTTON_${url.substring(8, 16)}`} />
    </button>
  );
};

export default SocialButtonSmall;
