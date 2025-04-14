import React from 'react';

const CartItemDescription = ({ item }) => {
  return (
<p className="text-gray-500 text-base mt-1 line-clamp-3 sm:line-clamp-3 md:line-clamp-2 lg:line-clamp-1 xl:line-clamp-1">
  {item.description}
</p>



  );
};

export default CartItemDescription;
