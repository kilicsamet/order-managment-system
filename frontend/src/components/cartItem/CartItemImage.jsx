import React from 'react';

const CartItemImage = ({ item }) => {
  return (
    <div className="flex items-start mt-5">
    <img
      src={item.image_url}
      alt={item.name}
      className="w-20 h-20 sm:w-30 sm:h-30 object-cover object-center rounded-md border border-gray-200"
    />
  </div>
  
  );
};

export default CartItemImage;
