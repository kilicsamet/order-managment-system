import React from 'react'

const CartSubTotal = ({totalPrice}) => {
  return (
    <div className="flex justify-between items-start border-b-2 border-solid border-gray-300 text-lg font-medium m-2 p-2">
    <span>Subtotal:</span>
    <span className="text-gray-500">${totalPrice.toFixed(2)}</span>
  </div>
  )
}

export default CartSubTotal
