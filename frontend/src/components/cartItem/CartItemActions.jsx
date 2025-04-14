import React from "react";

const CartItemActions = ({ item, removeItem }) => {
  return (
    <div className="mt-2">
      <button
        onClick={() => removeItem(item.id)}
        className="text-black underline hover:text-red-600 transition font-medium text-sm"
      >
        Remove Item
      </button>
    </div>
  );
};

export default CartItemActions;
