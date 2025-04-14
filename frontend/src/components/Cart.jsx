// Cart.js
import React from 'react';
import CartList from './cartItem/CartList';
import CartTotal from './cartTotal/CartTotal';


const Cart = ({items, increaseQuantity, decreaseQuantity, removeItem, totalPrice, setLoading, isLoading, createOrder }) => {
  return (
    <div className="flex flex-col sm:flex-row justify-between items-start space-x-4 m-4">
    <CartList items={items} increaseQuantity={increaseQuantity} decreaseQuantity={decreaseQuantity} removeItem={removeItem} />
    <CartTotal totalPrice={totalPrice} isLoading={isLoading} setLoading={setLoading} createOrder={createOrder}/>
</div>

  );
};

export default Cart;
