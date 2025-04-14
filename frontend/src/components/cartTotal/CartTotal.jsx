import React from 'react'
import CartTotalHeader from './CartTotalHeader'
import CartTotalCoupon from './CartTotalCoupon'
import CartSubTotal from './CartSubTotal'
import CartTotalAmount from './CartTotalAmount'
import ProceedButton from './ProceedButton'

const CartTotal = ({totalPrice, isLoading, setLoading, createOrder={createOrder}}) => {
  return (
    <div className="bg-white p-6 rounded-md w-full lg:w-2/3">
      <CartTotalHeader/>
      <CartTotalCoupon/>
      <CartSubTotal totalPrice={totalPrice}/>
      <CartTotalAmount totalPrice={totalPrice}/>
      <ProceedButton isLoading={isLoading} setLoading={setLoading} createOrder={createOrder}/>
    </div>
  )
}

export default CartTotal
