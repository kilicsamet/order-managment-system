import React from 'react'

const CartTotalAmount = ({totalPrice}) => {
  return (
    <div className="flex justify-between items-start text-lg font-medium m-2 p-2">
    <span>Total:</span>
    <span className="text-gray-500">${totalPrice.toFixed(2)}</span>
  </div>
  )
}

export default CartTotalAmount
