import React from 'react'

const CartTotalCoupon = () => {
  return (
    <div className="flex justify-between items-start border-b-2 border-solid border-gray-300 m-2 p-2">
    <p className="text-gray-500 underline transition font-medium">
      Add a coupon
    </p>
    <p className="text-gray-500">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        className="h-5 w-5 inline-block"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          strokeLinecap="round"
          strokeLinejoin="round"
          strokeWidth={2}
          d="M19 9l-7 7-7-7"
        />
      </svg>
    </p>
  </div>
  )
}

export default CartTotalCoupon
