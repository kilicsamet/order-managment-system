import React from "react";
import CartHeader from "./CartHeader";
import CartItemImage from "./CartItemImage";
import CartItemName from "./CartItemName";
import CartItemPrice from "./CartItemPrice";
import CartItemDescription from "./CartItemDescription";
import CartItemQuantity from "./CartItemQuantity";
import CartItemActions from "./CartItemActions";
import CartItemPriceTotal from "./CartItemPriceTotal";

const CartList = ({
  items,
  increaseQuantity,
  decreaseQuantity,
  removeItem,
}) => {
  return (
    <div className="bg-white p-6 rounded-md w-full lg:w-2/3">
      <CartHeader />
      <div className="flex flex-col ">
        {items.map((item) => (
          <div
            key={item.id}
            className="flex items-start border-b border-solid border-gray-300 min-h-[13rem] "
          >
            <CartItemImage item={item} />

            <div className="flex-1 text-start ms-3 mt-5">
              <CartItemName item={item} />
              <CartItemPrice item={item} />
              <CartItemDescription item={item} />
              <CartItemQuantity
                item={item}
                increaseQuantity={increaseQuantity}
                decreaseQuantity={decreaseQuantity}
              />
              <CartItemActions item={item} removeItem={removeItem} />
            </div>

          <CartItemPriceTotal item={item}/>
          </div>
        ))}
      </div>
    </div>
  );
};

export default CartList;
