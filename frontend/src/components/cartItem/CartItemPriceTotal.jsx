import React from 'react'

const CartItemPriceTotal = ({item}) => {
  return (
    <div className="flex items-start mt-5">
    <p className="text-gray-500 text-xl font-medium">
      ${(item.price * item.quantity).toFixed(2)}
    </p>
  </div>
  )
}

export default CartItemPriceTotal
