import React from "react";

const CartItemPrice = ({ item }) => {
  return (
    <p className="text-gray-500 mt-1 text-xl font-medium">
      ${item.price.toFixed(2)}
    </p>
  );
};

export default CartItemPrice;
