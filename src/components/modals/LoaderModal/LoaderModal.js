import React, {useRef, useEffect } from 'react';

const LoaderModal = () => {
    const modalRef = useRef(null);
    useEffect(() => {
        const body = document.querySelector("body");

        modalRef.current?.classList.toggle("loader-modal-body");

            if (!modalRef.current?.classList.contains("loader-modal-body")) {
                body.style.overflow = "loader-modal-body";
            } else {
                body.style.overflow = "auto";
            }
    }, [])

    return (
        <div ref={modalRef} className="loader-modal-body">
        </div>
    );
};

export default LoaderModal;