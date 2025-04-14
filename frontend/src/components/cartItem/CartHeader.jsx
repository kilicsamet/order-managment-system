import React from 'react';

const CartHeader = () => {
  return (
    <div className="hidden lg:flex justify-between items-center space-x-4 border-b-2 border-solid border-gray-300">
    <h2 className="text-l font-bold mb-4">Product</h2>
    <h2 className="text-l font-bold mb-4">Total</h2>
  </div>
  
  );
};

export default CartHeader;
