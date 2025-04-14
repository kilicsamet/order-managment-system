import React from 'react'

const CartItemName = ({item}) => {
  return (
    <h2 className="text-lg font-medium">{item.name}</h2>
  )
}

export default CartItemName
