import React from 'react';

function Modal({ children, onClose }) {
  return (
    <div className="modal-overlay">
      <div className="modal-content">
        <button onClick={onClose}>&times;</button>
        {children}
      </div>
    </div>
  );
}

export default Modal;