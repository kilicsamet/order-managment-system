import React from 'react';

const CartItemQuantity = ({ item, increaseQuantity, decreaseQuantity }) => {
  return (
    <div className="flex items-center bg-gray-100 text-gray-500 border border-gray-300 rounded overflow-hidden max-w-[100px] w-full mt-1">
    <button
      onClick={() => decreaseQuantity(item.id)}
      className="px-3 py-1 text-gray-500 hover:bg-gray-200 w-1/3"
    >
      −
    </button>
    <span className="w-1/3 text-center ">
      {item.quantity}
    </span>
    <button
      onClick={() => increaseQuantity(item.id)}
      className="px-3 py-1 text-gray-500 hover:bg-gray-200 w-1/3"
    >
      +
    </button>
  </div>
  );
};

export default CartItemQuantity;
