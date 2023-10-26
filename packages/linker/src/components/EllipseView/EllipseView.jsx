import React, { useMemo } from 'react';

const COUNT = 5;
const MIN_COUNT = 2;
const EllipseView = () => {
  const getEllipseStyles = (count) => {
    return {
      width: count * 150,
      height: count * 150,
    };
  };

  const cycle = (count) => {
    if (count <= MIN_COUNT) return null;
    return renderEllipse(cycle(count - 1), count - 1);
  };

  const renderEllipse = (children, count) => {
    return (
      <div style={getEllipseStyles(count)} className="ellipse">
        {children}
      </div>
    );
  };

  return (
    <div className="ellipse-container">
      {cycle(COUNT)}
    </div>
  );
};

export default EllipseView;
